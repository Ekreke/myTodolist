package store

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"gorm.io/gorm"
)

type CollectionStore interface {
	AddItem(itemid int64, collectionid int64, username string) (*v1.CommonResponseWizMsg, error)
	Create(icon int64, name string, username string) (*v1.CommonResponseWizMsg, error)
	Delete(collectionid int64, username string) (*v1.CommonResponseWizMsg, error)
	DeleteItem(itemid int64, username string) (*v1.CommonResponseWizMsg, error)
	Update(collectionid int64, icon int64, name string, username string) (*v1.CommonResponseWizMsg, error)
	LoadItems(collectionid int64, username string) (*v1.CollectionLoadItemsResp, error)
}

type collection struct {
	db *gorm.DB
}

// AddItem implements CollectionStore.
func (c *collection) AddItem(itemid int64, collectionid int64, username string) (*v1.CommonResponseWizMsg, error) {
	panic("unimplemented")
}

// Create implements CollectionStore.
func (c *collection) Create(icon int64, name string, username string) (*v1.CommonResponseWizMsg, error) {
	collection := &model.Collections{Icon: icon, Name: name}
	tx := c.db.Begin()
	// save record to collections
	if err := tx.Debug().Create(&collection).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	// save record to collections-users
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	if err := c.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error; err != nil {
		log.Errorw("get userid from users failed")
		tx.Rollback()
		return nil, err
	}
	c_u := &model.CollectionsUsers{CollectionId: collection.ID, UserId: tmpu.ID}
	if err := tx.Debug().Create(&c_u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// Delete implements CollectionStore.
func (c *collection) Delete(collectionid int64, username string) (*v1.CommonResponseWizMsg, error) {
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	if err := c.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error; err != nil {
		log.Fatalw("get userid from username failed")
	}

	tx := c.db.Begin()
	// delete record from collections
	co := &model.Collections{}
	if err := tx.Debug().Where("id = ?", collectionid).Delete(&co).Error; err != nil {
		tx.Rollback()
		log.Errorw("delete record from collections failed")
		return nil, err
	}
	// delete record from collection_users
	cu := &model.CollectionsUsers{}
	if err := tx.Debug().Where("collection_id = ? and user_id = ?", collectionid, tmpu.ID).Delete(&cu).Error; err != nil {
		tx.Rollback()
		log.Errorw("delete record from collection_users failed")
		return nil, err
	}
	// delete record from collections_items
	ci := &model.CollectionsItems{}
	if err := tx.Debug().Where("collection_id = ?", collectionid).Delete(&ci).Error; err != nil {
		tx.Rollback()
		log.Errorw("delete record from collections_items failed")
		return nil, err
	}
	tx.Commit()
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// DeleteItem implements CollectionStore.
func (c *collection) DeleteItem(itemid int64, username string) (*v1.CommonResponseWizMsg, error) {
	panic("unimplemented")
}

// LoadItems implements CollectionStore.
func (c *collection) LoadItems(collectionid int64, username string) (*v1.CollectionLoadItemsResp, error) {
	panic("unimplemented")
}

// Update implements CollectionStore.
func (c *collection) Update(collectionid int64, icon int64, name string, username string) (*v1.CommonResponseWizMsg, error) {
	panic("unimplemented")
}

var _ CollectionStore = (*collection)(nil)

func newCollection(db *gorm.DB) *collection {
	return &collection{db: db}
}
