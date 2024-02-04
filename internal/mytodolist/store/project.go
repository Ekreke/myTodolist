package store

import (
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
}

type projectStore struct {
	db *gorm.DB
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
	p := &model.Projects{}
	if err := ps.db.Debug().Where("id = ?", projectid).First(&p).Error; err != nil {
		return *p, err
	}

	return *p, nil

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
