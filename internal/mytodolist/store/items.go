package store

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"gorm.io/gorm"
)

type ItemStore interface {
	Create(it *model.Items, username string) (resp *v1.CommonResponseWizMsg, err error)
	Delete(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	Info(itemid int, username string) (resp *v1.ItemInfoResponse, err error)
	SetUnDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	SetDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error)
	Update(it *model.Items, username string) (resp *v1.CommonResponseWizMsg, err error)
}

type items struct {
	db *gorm.DB
}

func newItems(db *gorm.DB) *items {
	return &items{db: db}
}

// create items
func (i *items) Create(it *model.Items, username string) (resp *v1.CommonResponseWizMsg, err error) {
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = i.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
	}
	// it.Deadline = time.Now()
	// insert into projects nodes
	tx := i.db.Begin()
	tx.SavePoint("begin")
	// FIXME: item index
	// insert into items
	err = tx.Save(&it).Error
	if err != nil {
		log.Errorw("save record to items failed")
		tx.RollbackTo("begin")
		return nil, err
	}

	if it.ProjectId != 0 {
		i_p := &model.ProjectsNodes{
			ProjectId: it.ProjectId,
			ItemId:    it.ID,
			UserId:    tmpu.ID,
		}
		err := tx.Debug().Create(&i_p).Error
		if err != nil {
			log.Errorw("save record to projects_nodes failed")
			tx.RollbackTo("begin")
			return nil, err
		}
	}

	// insert into myday && users
	if it.Myday != 0 {
		i_m := &model.Myday{
			Item_id: it.ID,
			User_id: tmpu.ID,
		}
		err := tx.Create(&i_m).Error
		if err != nil {
			log.Errorw("save record to myday failed")
			return nil, err
		}
	}

	// insert into collections && items
	if it.CollectionId != 0 {
		i_c := &model.CollectionsItems{
			ItemsId:      it.ID,
			CollectionId: it.CollectionId,
		}

		err := tx.Create(&i_c).Error
		if err != nil {
			log.Errorw("save record to collections_items failed ")
			return nil, err
		}
	}
	// insert into items && users
	i_u := &model.ItemsUsers{
		ItemId: it.ID,
		UserId: tmpu.ID,
	}
	err = tx.Debug().Create(&i_u).Error
	if err != nil {
		tx.RollbackTo("begin")
		log.Errorw("save record to items_users failed")
		return nil, err
	}
	tx.Commit()
	resp = &v1.CommonResponseWizMsg{
		Msg: "success",
	}

	return resp, nil
}

// delete a item by item id
func (i *items) Delete(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {

	// get item info
	item := &model.Items{}
	err = i.db.Where("id = ?", itemid).Find(&item).Error
	if err != nil {
		log.Errorw("get item info error")
		return nil, err
	}

	tx := i.db.Begin()
	// delete from projects
	if item.ProjectId != 0 {
		i_p := &model.ProjectsNodes{}
		err = tx.Where("item_id = ?", itemid).Delete(&i_p).Error
		if err != nil {
			log.Errorw("delete project_item failed")
			tx.Rollback()
		}

	}
	// delete from mydays
	if item.Myday != 0 {
		i_m := &model.Myday{}
		err = tx.Where("item_id = ?", itemid).Delete(&i_m).Error
		if err != nil {
			log.Errorw("delete myday_item failed")
			tx.Rollback()
		}
	}
	// delete from collections_items

	if item.CollectionId != 0 {
		i_c := &model.CollectionsItems{}
		// FIXME: column should be item
		err = tx.Where("items_id = ?", itemid).Delete(&i_c).Error
		if err != nil {
			log.Errorw("delete collections_item failed")
			tx.Rollback()
		}
	}
	// delete from items_users
	i_u := &model.ItemsUsers{}
	err = tx.Where("item_id = ?", itemid).Delete(&i_u).Error
	if err != nil {
		log.Errorw("delete user_items failed")
		tx.Rollback()
	}

	// delete from items
	err = tx.Where("id = ?", itemid).Delete(&item).Error
	if err != nil {
		log.Errorw("delete items failed")
		tx.Rollback()
	}
	tx.Commit()
	resp = &v1.CommonResponseWizMsg{
		Msg: "success",
	}
	return resp, nil
}

// get info by item id
func (i *items) Info(itemid int, username string) (resp *v1.ItemInfoResponse, err error) {
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = i.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
		return nil, err
	}

	item := &model.Items{}
	err = i.db.Where("id = ?", itemid).Find(&item).Error
	if err != nil {
		log.Errorw("find item by item id failed")
	}

	r := &v1.ItemInfoResponse{
		ItemName:     item.ItemName,
		Description:  item.Description,
		ProjectId:    item.ProjectId,
		Deadline:     item.Deadline,
		Important:    item.Important,
		Done:         item.Done,
		Myday:        item.Myday,
		CreatedTime:  item.CreatedTime,
		Node:         item.Node,
		Checkpoint:   item.Checkpoint,
		CollectionId: item.CollectionId,
	}
	return r, nil
}

// set a item undone by item id
func (i *items) SetUnDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = i.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
	}
	return resp, nil
}

// set a item done by item id
func (i *items) SetDone(itemid int, username string) (resp *v1.CommonResponseWizMsg, err error) {

	return resp, nil
}

// update a item's info
func (i *items) Update(it *model.Items, username string) (resp *v1.CommonResponseWizMsg, err error) {
	return resp, nil
}
