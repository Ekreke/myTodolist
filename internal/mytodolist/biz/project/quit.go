package project

import v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"

// Quit implements ProjectBiz.
func (pb *projectBiz) Quit(userid int64, projectid int64) (resp *v1.CommonResponseWizMsg, err error) {
	_, err = pb.ds.Projects().DeleteRecordFromPU(userid, projectid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}
