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
	DeleteItem(itemid int64, collectionid int64, username string) (*v1.CommonResponseWizMsg, error)
	Update(collectionid int64, icon int64, name string, username string) (*v1.CommonResponseWizMsg, error)
	LoadItems(collectionid int64, username string) (*v1.CollectionLoadItemsResp, error)
}

type collection struct {
	db *gorm.DB
}

// AddItem implements CollectionStore.
func (c *collection) AddItem(itemid int64, collectionid int64, username string) (*v1.CommonResponseWizMsg, error) {
	// save record to collections-items
	tx := c.db.Begin()
	ci := &model.CollectionsItems{ItemsId: itemid, CollectionId: collectionid}
	if err := tx.Debug().Create(&ci).Error; err != nil {
		log.Errorw("save record to collections-items failed")
		tx.Rollback()
		return nil, err
	}
	// FIXME: 废弃的字段：items->dollection_id
	// add item's collection id
	// return response
	tx.Commit()
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
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
func (c *collection) DeleteItem(itemid int64, collectionid int64, username string) (*v1.CommonResponseWizMsg, error) {
	// delete item from collections_items
	ci := &model.CollectionsItems{}
	if err := c.db.Debug().Where("collection_id = ? and items_id = ?", collectionid, itemid).Delete(&ci).Error; err != nil {
		log.Errorw("delete ci failed")
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// LoadItems implements CollectionStore.
func (c *collection) LoadItems(collectionid int64, username string) (*v1.CollectionLoadItemsResp, error) {
	tx := c.db.Begin()
	citemids := &[]model.CollectionsItems{}
	if err := tx.Debug().Where("collection_id = ?", collectionid).Find(&citemids).Error; err != nil {
		log.Errorw("get items id failed")
		tx.Rollback()
		return nil, err
	}
	ids := []int{}
	for _, itemid := range *citemids {
		ids = append(ids, int(itemid.ItemsId))
	}

	items := &[]model.Items{}
	if err := tx.Debug().Where("id in ?", ids).Find(&items).Error; err != nil {
		log.Errorw("get item by item id failed")
		tx.Rollback()
		return nil, err
	}
	return &v1.CollectionLoadItemsResp{Items: *items}, nil
}

// Update implements CollectionStore.
func (c *collection) Update(collectionid int64, icon int64, name string, username string) (*v1.CommonResponseWizMsg, error) {
	// get user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	if err := c.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error; err != nil {
		log.Fatalw("get userid from username failed")
	}

	// select pre collection
	preco := &model.Collections{}
	tx := c.db.Begin()
	if err := tx.Debug().Table("collections").Where("id = ?", collectionid).First(&preco).Error; err != nil {
		log.Errorw("get collection by id failed")
		tx.Rollback()
		return nil, err
	}

	// update icon if icon changed
	if preco.Icon != icon {
		if err := tx.Debug().Model(&model.Collections{}).Where("id = ?", collectionid).Update("icon", icon).Error; err != nil {
			log.Errorw("update icon failed")
			tx.Rollback()
			return nil, err
		}
	}
	// update name if name changed
	if preco.Name != name {
		if err := tx.Debug().Model(&model.Collections{}).Where("id = ?", collectionid).Update("name", name).Error; err != nil {
			log.Errorw("update name failed")
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
	// commit
}

var _ CollectionStore = (*collection)(nil)

func newCollection(db *gorm.DB) *collection {
	return &collection{db: db}
}
