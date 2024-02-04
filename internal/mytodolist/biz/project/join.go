package project

import (
	"errors"

	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
	"gorm.io/gorm"
)

// Join implements ProjectBiz.
func (pb *projectBiz) Join(userid int64, projectid int64, pwd string) (resp *v1.CommonResponseWizMsg, err error) {
	// check project if exist / if user in the project
	exist, err := pb.ds.Projects().CheckProjectIfExist(projectid)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errors.New("project not exists")
	}
	// join twice
	in, err := pb.ds.Projects().CheckUserIfInProject(projectid, userid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if in {
		return &v1.CommonResponseWizMsg{Msg: "already in the project"}, nil
	}
	// check pwd if match
	match, err := pb.ds.Projects().CheckPwdIfMatch(projectid, pwd)
	if err != nil {
		return nil, errors.New("check pwd match failed ")
	}
	if !match {
		return nil, errors.New("pwd not match")
	}
	// add record to project_user table
	_, err = pb.ds.Projects().AddRecordPU(projectid, userid)
	if err != nil {
		return nil, errors.New("add record to project users failed ")
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}
