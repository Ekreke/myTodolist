package store

import (
	"github.com/ekreke/myTodolist/internal/pkg/model"
	"gorm.io/gorm"
)

type ProjectStore interface {
	DeleteRecordFromPU(userid int64, projectid int64) (affectedRows int, err error)
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
