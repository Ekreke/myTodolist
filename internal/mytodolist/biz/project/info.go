package project

import v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"

// Info implements ProjectBiz.
func (pb *projectBiz) Info(projectid int64) (*v1.ProjectInfoResponse, error) {
	resp, err := pb.ds.Projects().GetProjectInfoById(projectid)
	if err != nil {
		return nil, err
	}
	userInfos, err := pb.ds.Projects().GetProjectJoinedUserinfoByProjectId(projectid)
	if err != nil {
		return nil, err
	}
	r := &v1.ProjectInfoResponse{Projects: resp, Users: userInfos}
	return r, nil
}
