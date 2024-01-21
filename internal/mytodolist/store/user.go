package store

import (
	"context"

	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"github.com/ekreke/myTodolist/pkg/token"
	"gorm.io/gorm"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.Users) error
	Get(ctx context.Context, username string) (*model.Users, error)
	GetInfo(username string) (*v1.InfoResponse, error)
	CheckUserIfExist(username string) (bool, error)
	UpdateInfo(req *v1.UpdateInfoRequest, username string) error
	GetImportantItems(nextid int, pagesize int, username string) ([]v1.ItemInfo, token.Page, error)
	UpdatePwd(username string, newpwd string) error
	GetCollections(username string) (*v1.CollectionsResponse, error)
	GetMydayItems(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error)
	LoadItems(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error)
	LoadNodes(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error)
	LoadMyItems(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error)
}

// UserStore 接口的实现.
type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

// Create 插入一条 user 记录.
func (u *users) Create(ctx context.Context, user *model.Users) error {
	return u.db.Create(&user).Error
}

// Get 根据用户名查询指定user的数据库记录
func (u *users) Get(ctx context.Context, username string) (*model.Users, error) {
	var user model.Users
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *users) GetInfo(username string) (*v1.InfoResponse, error) {
	var user model.Users
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	resp := &v1.InfoResponse{
		Username:   user.Username,
		Bio:        user.Bio,
		Link:       user.Link,
		Avatar:     user.Avatar,
		Root:       int(user.Root),
		Created_At: user.CreatedAt,
	}
	return resp, nil
}

// if exist , return true else return false
func (u *users) CheckUserIfExist(username string) (bool, error) {
	user := &model.Users{}
	if err := u.db.Debug().Where("username = ?", username).First(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (u *users) UpdateInfo(req *v1.UpdateInfoRequest, username string) error {
	// user := &model.Users{}
	user := &model.Users{}
	err := u.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Link != "" {
		user.Link = req.Link
	}
	return u.db.Save(&user).Error
}

func (u *users) GetImportantItems(next_id int, page_size int, username string) ([]v1.ItemInfo, token.Page, error) {
	// 根据 username 查询所有的item id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err := u.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Infow("根据 username 查询所有的item id err :", err)
	}
	users_items := &[]model.ItemsUsers{}
	// FIXME: pagesize
	err = u.db.Debug().Select("item_id").Where("user_id = ? and item_id > ?", tmpu.ID, next_id).Limit(page_size).Find(&users_items).Error
	if err != nil {
		return nil, token.Page{}, err
	}
	// 获取 全部的id
	var ids []int
	for _, v := range *users_items {
		ids = append(ids, int(v.ItemId))
	}

	var items []model.Items
	// 根据item id 查询所有的item
	err = u.db.Debug().Where("id in ?", ids).Find(&items).Order("created_time ASC").Error
	if err != nil {
		return nil, token.Page{}, err
	}
	lsItemId := items[len(items)-1].ID

	// 更新page
	page := token.Page{
		NextID:   int(lsItemId),
		PageSize: int64(page_size),
	}
	// 返回数据(page / info / error)
	var resp []v1.ItemInfo
	for _, v := range items {
		i := &v1.ItemInfo{
			ID:          v.ID,
			ItemName:    v.ItemName,
			Description: v.Description,
			ProjectId:   v.ProjectId,
			Deadline:    v.Deadline,
			Done:        v.Done,
			CreatedTime: v.CreatedTime,
		}
		resp = append(resp, *i)

	}
	return resp, page, nil

}

func (u *users) UpdatePwd(username string, newpwd string) error {
	user := &model.Users{}
	err := u.db.Debug().Where("username = ?", username).First(&user).Error
	if err != nil {
		return err
	}
	user.Password = newpwd
	return u.db.Save(&user).Error
}

func (u *users) GetCollections(username string) (*v1.CollectionsResponse, error) {
	// get user id by username
	tmpu := &model.Users{}
	// select  id from users where username = ?
	err := u.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Errorw("find user id failed")
		return nil, err
	}
	collectionsids := &[]model.CollectionsUsers{}
	var ids []int

	// select collection_id from collections_users where user_id = ?
	// get collection id by user_id from the middleware table
	err = u.db.Debug().Select("collection_id").Where("user_id = ?", tmpu.ID).Find(&collectionsids).Error
	if err != nil {
		log.Errorw("find collections_ids failed")
		return nil, err
	}
	for _, v := range *collectionsids {
		ids = append(ids, int(v.CollectionId))
	}
	// get collections by collection_id from collections table
	collections := &[]model.Collections{}
	// select * from collections where id in ?
	err = u.db.Debug().Where("id in ?", ids).Find(&collections).Error
	if err != nil {
		log.Errorw("find collections failed")
		return nil, err
	}

	// inject collections into response
	resp := &v1.CollectionsResponse{}
	resp.Collections = append(resp.Collections, *collections...)

	return resp, nil
}

func (u *users) GetMydayItems(next_id int, page_size int, username string) (items []model.Items, npage token.Page, err error) {
	// 根据 username 查询 user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = u.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
	}

	// get myday_items id from myday_items_users table
	myday_ids := &[]model.Myday{}
	err = u.db.Debug().Select("item_id").Where("user_id = ? and item_id > ? ", tmpu.ID, next_id).Limit(page_size).Find(&myday_ids).Error
	if err != nil {
		log.Fatalw("get myday from myday table failed")
	}
	// get ids
	var ids []int
	for _, v := range *myday_ids {
		ids = append(ids, int(v.Item_id))
	}
	its := []model.Items{}
	err = u.db.Debug().Where("id in ?", ids).Find(&its).Error
	if err != nil {
		log.Fatalw("get items from items table failed")
	}
	if len(its) != 0 {
		npage.NextID = int(its[len(its)-1].ID)
	}
	npage.PageSize = int64(page_size)
	return its, npage, nil
}

func (u *users) LoadItems(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error) {
	// 根据 username 查询 user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = u.db.Debug().Table("users").Select("id").Where("username = ?", username).Limit(int(npage.PageSize)).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
	}
	item_users := []model.ItemsUsers{}
	err = u.db.Debug().Where("user_id =  ? and item_id > ?", tmpu.ID, nextid).Find(&item_users).Error
	if err != nil {
		log.Fatalw("get items - users failed")
	}

	// get item ids
	ids := []int{}
	for _, v := range item_users {
		ids = append(ids, int(v.ItemId))
	}

	// select items by ids
	err = u.db.Debug().Where("id in ?", ids).Find(&items).Error
	if err != nil {
		log.Fatalw("get items by ids failed", "err:", err)
	}
	// if record is not nil , update npage
	if len(items) != 0 {
		npage.NextID = int(items[len(items)-1].ID)
		npage.PageSize = int64(pagesize)
	}

	return items, npage, nil

}
func (u *users) LoadNodes(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error) {
	// 根据 username 查询 user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = u.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
		return nil, token.Page{}, err
	}

	// get user's all nodes from projects_nodes
	projects_nodes := &[]model.ProjectsNodes{}
	// select  item_id from projects_nodes where user_id = ? and item_id > ?
	err = u.db.Debug().Where("user_id = ? and item_id > ?", tmpu.ID, nextid).Select("item_id").Limit(pagesize).Find(&projects_nodes).Error
	if err != nil {
		log.Fatalw("get nodes from projects_nodes failed")
		return nil, token.Page{}, err
	}
	// get item ids
	ids := []int{}
	for _, v := range *projects_nodes {
		ids = append(ids, int(v.ItemId))
	}

	// select items by ids
	err = u.db.Debug().Where("id in ?", ids).Find(&items).Error
	if err != nil {
		log.Fatalw("get items by ids failed", "err:", err)
	}
	// if record is not nil , update npage
	if len(items) != 0 {
		npage.NextID = int(items[len(items)-1].ID)
		npage.PageSize = int64(pagesize)
	}

	return items, npage, nil

}
func (u *users) LoadMyItems(nextid int, pagesize int, username string) (items []model.Items, npage token.Page, err error) {
	// 根据 username 查询 user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err = u.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Fatalw("get userid from username failed")
	}

	items_users := []model.ItemsUsers{}
	// select  item_id from items_users where user_id = ?
	err = u.db.Debug().Where("user_id = ? and item_id > ?", tmpu.ID, nextid).Limit(pagesize).Find(&items_users).Error
	if err != nil {
		log.Fatalw("get items - users failed")
		return nil, token.Page{}, err
	}
	// item ids
	ids := []int{}
	for _, v := range items_users {
		ids = append(ids, int(v.ItemId))
	}

	// select items by ids
	err = u.db.Debug().Where("id in ?", ids).Find(&items).Error
	if err != nil {
		log.Fatalw("get items by ids failed", "err:", err)
	}
	// if record is not nil , update npage
	if len(items) != 0 {
		npage.NextID = int(items[len(items)-1].ID)
		npage.PageSize = int64(pagesize)
	}
	return items, npage, nil
}
