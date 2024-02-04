package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

// Myprojects implements ProjectBiz.
func (pb *projectBiz) Myprojects(userid int64) (resp *v1.MyprojectsResponse, err error) {
	// load projects ids
	ids, err := pb.ds.Projects().GetProjectsIdsByUserId(userid)
	if err != nil {
		log.Errorw("get projectsids by userid failed")
		return nil, err
	}
	// get projects info
	r := &v1.MyprojectsResponse{}
	for _, id := range ids {
		project, err := pb.ds.Projects().GetProjectInfoById(int64(id))
		if err != nil {
			log.Errorw("get project info by id failed")
			return nil, err
		}
		r.Projects = append(r.Projects, project)
	}

	// load in response
	return r, nil
}
