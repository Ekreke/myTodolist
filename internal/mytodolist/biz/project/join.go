package project

import v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"

// Join implements ProjectBiz.
func (pb *projectBiz) Join(userid int64, projectid int64, pwd string) (resp *v1.CommonResponseWizMsg, err error) {
	// check project if exist
	// check pwd if match
	// add record to project_user table
	panic("unimplemented")
}
