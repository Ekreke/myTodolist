package store

import (
	"sync"

	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	"gorm.io/gorm"
)

var (
	once sync.Once
	// 全局变量，方便其他包直接调用已经初始化好的S实例
	S *datastore
)

type Istore interface {
	Users() UserStore
	Items() ItemStore
	Collection() CollectionStore
	Projects() ProjectStore
}

type datastore struct {
	db *gorm.DB
}

var _ Istore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// Users 返回一个实现了 UserStore 接口的实例.
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}

func (ds *datastore) Items() ItemStore {
	return newItems(ds.db)
}

func (ds *datastore) Collection() CollectionStore {
	return newCollection(ds.db)
}

func (ds *datastore) Projects() ProjectStore {
	return newProjects(ds.db)
}

func GetUserIdByUserName(username string) (int, error) {
	// 根据 username 查询 user id
	tmpu := &model.Users{}
	// select id from users where username = ?
	err := S.db.Debug().Table("users").Select("id").Where("username = ?", username).First(&tmpu).Error
	if err != nil {
		log.Infow("get userid from username failed")
		return 0, err
	}
	return int(tmpu.ID), nil
}

func (ds *datastore) DB() *gorm.DB {
	return ds.db
}
