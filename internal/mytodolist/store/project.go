package store

import (
	"strconv"
	"time"

	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	"gorm.io/gorm"
)

type ProjectStore interface {
	DeleteRecordFromPU(userid int64, projectid int64) (affectedRows int, err error)
	CheckProjectIfExist(projectid int64) (exist bool, err error)
	CheckUserIfInProject(projectid int64, userid int64) (in bool, err error)
	CheckPwdIfMatch(projectid int64, pwd string) (match bool, err error)
	AddRecordPU(projectid int64, userid int64) (affectedRows int, err error)
	GetProjectInfoById(projectid int64) (project model.Projects, err error)
	GetProjectsIdsByUserId(userid int64) (ids []int, err error)
	DeleteNode(projectid string, nodeid string, userid int64) (affectedRows int, err error)
	DeleteProject(projectid string, userid int64) (affectedRows int, err error)
	CreateProject(project *model.Projects, userid int64) (affectedRows int, err error)
	AddNode(projectid string, nodeid string, userid int64) (affectedRows int, err error)
	GetAllProjectsICreated(userid int64) (projects []model.Projects, err error)
	UpdateProjectInfo(id int, description string, endtime int64, name string, userid int64) (affectedRows int, err error)
	UpdateNode(projectid string, nodeid string, userid int, item *model.Items) (affectedRows int, err error)
	NodeInfo(projectid string, nodeid string, userid int) (item *model.Items, err error)
	Nodes(projectid string, userid int) (items *[]model.Items, err error)
}

type projectStore struct {
	db *gorm.DB
}

// Nodes implements ProjectStore.
func (ps *projectStore) Nodes(projectid string, userid int) (items *[]model.Items, err error) {
	is := &[]model.ProjectsNodes{}
	// select ids of items
	tx := ps.db.Begin()
	err = tx.Debug().Where("project_id = ? and user_id = ? ", projectid, userid).Find(is).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// get ids
	ids := make([]int, 0)
	for _, v := range *is {
		ids = append(ids, int(v.ItemId))
	}
	// get items
	err = tx.Where("id in ? ", ids).Find(&items).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return items, nil
}

// AddNode implements ProjectStore.
func (ps *projectStore) AddNode(projectid string, nodeid string, userid int64) (affectedRows int, err error) {
	// set item's column node = true
	tx := ps.db.Begin()
	err = tx.Debug().Model(&model.Items{}).Where("id = ?", nodeid).Update("node", 1).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	err = tx.Debug().Model(&model.Items{}).Where("id = ?", nodeid).Update("project_id", projectid).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	// add item to projects nodes
	// FIXME: err checking
	pid, _ := strconv.Atoi(projectid)
	nid, _ := strconv.Atoi(nodeid)
	err = tx.Debug().Model(&model.ProjectsNodes{}).Create(&model.ProjectsNodes{ProjectId: int64(pid), ItemId: int64(nid), UserId: userid}).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	tx.Commit()
	return int(ps.db.RowsAffected), nil
}

// CreateProject implements ProjectStore.
// FIXME: if one person created a projectname duplicate it will cover the info of previous one
func (ps *projectStore) CreateProject(project *model.Projects, userid int64) (affectedRows int, err error) {
	tx := ps.db.Begin()
	// create project
	if err = tx.Create(project).Error; err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}

	// create projects_users
	pu := &model.ProjectsUsers{ProjectsId: project.ID, UsersId: userid}
	err = tx.Create(pu).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	return int(ps.db.RowsAffected), nil
}

// DeleteProject implements ProjectStore.
func (ps *projectStore) DeleteProject(projectid string, userid int64) (affectedRows int, err error) {
	tx := ps.db.Begin()
	//delete projects
	pid, err := strconv.Atoi(projectid)
	if err != nil {
		return 0, err
	}
	err = tx.Delete(&model.Projects{ID: int64(pid)}).Error
	if err != nil {
		return int(ps.db.RowsAffected), err
	}
	//delete projects nodes
	pn := &[]model.ProjectsNodes{}
	err = tx.Where("project_id = ?", projectid).Delete(&pn).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	//delete projects users
	err = tx.Where("projects_id = ?", projectid).Delete(&model.ProjectsUsers{}).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	//get item ids
	items := &[]model.Items{}
	err = tx.Where("project_id = ?", projectid).Find(&items).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	ids := []int{}
	for _, item := range *items {
		ids = append(ids, int(item.ID))
	}

	//delete items
	err = tx.Where("id in ? ", ids).Delete(&model.Items{}).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	//delete items users
	err = tx.Where("item_id in ?", ids).Delete(&model.ItemsUsers{}).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	//delete mydays
	err = tx.Where("item_id in ?", ids).Delete(&model.Myday{}).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	tx.Commit()
	return int(ps.db.RowsAffected), nil
	//commit
}

// GetAllProjectsICreated implements ProjectStore.
func (ps *projectStore) GetAllProjectsICreated(userid int64) (projects []model.Projects, err error) {
	// get projects id
	pus := &[]model.ProjectsUsers{}
	tx := ps.db.Begin()
	err = tx.Where("users_id = ?", userid).Find(&pus).Error
	if err != nil {
		tx.Rollback()
		return nil, err

	}
	// get projects info
	ids := []int{}
	for _, pu := range *pus {
		ids = append(ids, int(pu.ProjectsId))
	}

	prs := &[]model.Projects{}
	err = tx.Where("id in ?", ids).Find(&prs).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	// return
	return *prs, nil
}

// NodeInfo implements ProjectStore.
func (ps *projectStore) NodeInfo(projectid string, nodeid string, userid int) (item *model.Items, err error) {
	// get project nodes id
	// ids := &[]model.ProjectsNodes{}
	// err = ps.db.Where("project_id = ? and user_id = ?", projectid, userid).Error
	// if err != nil {
	// 	return nil, err
	// }

	// get item id and than get the item info
	err = ps.db.Where("id = ?", nodeid).Find(&item).Error
	if err != nil {
		return nil, err
	}
	// set item info into resp
	// return
	return item, nil

}

// UpdateNode implements ProjectStore.
func (ps *projectStore) UpdateNode(projectid string, nodeid string, userid int, item *model.Items) (affectedRows int, err error) {
	// get pre info
	iteminfo := &model.Items{}
	err = ps.db.Where("id = ?", nodeid).Find(&iteminfo).Error
	if err != nil {
		return int(ps.db.RowsAffected), err
	}
	// update info by rules
	err = ps.db.Save(item).Error
	if err != nil {
		return int(ps.db.RowsAffected), err
	}
	return int(ps.db.RowsAffected), nil
}

// UpdateProjectInfo implements ProjectStore.
func (ps *projectStore) UpdateProjectInfo(id int, description string, endtime int64, name string, userid int64) (affectedRows int, err error) {
	p := &model.Projects{
		Description: description,
		EndTime:     time.Unix(0, endtime),
		Name:        name,
		CreatedTime: time.Now(),
	}
	err = ps.db.Debug().Where("id = ?", id).First(&p).Error
	if err != nil {
		return int(ps.db.RowsAffected), err
	}
	p.Description = description
	p.EndTime = time.Unix(0, endtime)
	p.Name = name
	err = ps.db.Debug().Where("id = ?", id).Save(&p).Error
	if err != nil {
		return int(ps.db.RowsAffected), err
	}
	return int(ps.db.RowsAffected), nil
}

// DeleteNode implements ProjectStore.  only delete node
func (ps *projectStore) DeleteNode(projectid string, nodeid string, userid int64) (affectedRows int, err error) {
	// delete from projects_nodes
	tx := ps.db.Begin()
	pn := &model.ProjectsNodes{}
	err = tx.Where("item_id = ? and project_id = ?", nodeid, projectid).Delete(&pn).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	// unset item's node column
	err = tx.Model(&model.Items{}).Where("id = ?", nodeid).Update("node", 0).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	// update project column
	err = tx.Model(&model.Items{}).Where("id = ?", nodeid).Update("project_id", 0).Error
	if err != nil {
		tx.Rollback()
		return int(ps.db.RowsAffected), err
	}
	tx.Commit()
	return int(ps.db.RowsAffected), nil

}

// DeleteRecordFromPU implements ProjectStore.
func (ps *projectStore) DeleteRecordFromPU(userid int64, projectid int64) (affectedRows int, err error) {
	pu := &model.ProjectsUsers{}
	if err := ps.db.Debug().Where("users_id = ? and projects_id = ?", userid, projectid).Delete(&pu).Error; err != nil {
		return int(ps.db.RowsAffected), err
	}
	return int(ps.db.RowsAffected), nil
}

var _ ProjectStore = (*projectStore)(nil)

func newProjects(db *gorm.DB) *projectStore {
	return &projectStore{db: db}
}

// CheckProjectIfExist implements ProjectStore.
func (ps *projectStore) CheckProjectIfExist(projectid int64) (exist bool, err error) {
	p := &model.Projects{}
	if err = ps.db.Debug().Where("id = ?", projectid).Select("id").First(&p).Error; err != nil {
		log.Errorw("check project if exist error")
		return false, err

	}
	if p.ID == projectid {
		exist = true
	}
	return exist, nil

}

// CheckPwdIfMatch implements ProjectStore.
func (ps *projectStore) CheckPwdIfMatch(projectid int64, pwd string) (match bool, err error) {
	p := &model.Projects{}
	if err := ps.db.Debug().Where("id = ? and password = ?", projectid, pwd).First(&p).Error; err != nil {
		log.Errorw("check project pwd if match failed")
		return false, err
	}
	return true, nil
}

// CheckUserIfInProject implements ProjectStore.
func (ps *projectStore) CheckUserIfInProject(projectid int64, userid int64) (in bool, err error) {
	// select record from projects_users
	pu := &model.ProjectsUsers{}
	if err := ps.db.Debug().Where("projects_id = ? and users_id = ?", projectid, userid).First(&pu).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Errorw("check user if in project failed")
			return false, err
		} else {
			log.Errorw("user not in this project")
			return false, nil
		}
	}
	return true, nil
}

// addRecordPU implements ProjectStore.
func (ps *projectStore) AddRecordPU(projectid int64, userid int64) (affectedRows int, err error) {
	pu := &model.ProjectsUsers{ProjectsId: projectid, UsersId: userid}
	if err := ps.db.Debug().Create(&pu).Error; err != nil {
		return int(ps.db.RowsAffected), err
	}
	return int(ps.db.RowsAffected), nil
}

// GetProjectInfoById implements ProjectStore.
func (ps *projectStore) GetProjectInfoById(projectid int64) (project model.Projects, err error) {
	// p := &model.Projects{}
	// if err := ps.db.Debug().Where("id = ?", projectid).First(&p).Error; err != nil {
	// 	return *p, err
	// }
	p, err := getProjectInfoById(ps.db, projectid)
	if err != nil {
		return p, err
	}
	return p, nil

}

// GetProjectsIdsByUserId implements ProjectStore.
func (ps *projectStore) GetProjectsIdsByUserId(userid int64) (ids []int, err error) {
	pus := &[]model.ProjectsUsers{}
	if err := ps.db.Debug().Where("users_id = ?", userid).Find(&pus).Error; err != nil {
		return ids, err
	}
	for _, pu := range *pus {
		ids = append(ids, int(pu.ProjectsId))
	}
	return ids, nil
}

func getProjectInfoById(tx *gorm.DB, projectid int64) (project model.Projects, err error) {
	p := &model.Projects{}
	if err := tx.Debug().Where("id = ?", projectid).First(&p).Error; err != nil {
		return *p, err
	}
	return *p, nil
}
