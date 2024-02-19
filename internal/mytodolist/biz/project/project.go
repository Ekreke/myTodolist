package project

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/ekreke/myTodolist/internal/pkg/model"
	v1 "github.com/ekreke/myTodolist/pkg/api/mytodolist"
)

type ProjectBiz interface {
	Join(userid int64, projectid int64, pwd string) (*v1.CommonResponseWizMsg, error)
	Myprojects(userid int64) (*v1.MyprojectsResponse, error)
	Quit(userid int64, projectid int64) (*v1.CommonResponseWizMsg, error)
	Create(description string, endtime int64, name string, userid int64) (*v1.CommonResponseWizMsg, error)
	Info(projectid int64) (*v1.ProjectInfoResponse, error)
	Delete(projectid string, userid int64) (*v1.CommonResponseWizMsg, error)
	DeleteNode(projectid string, nodeid string, userid int64) (*v1.CommonResponseWizMsg, error)
	AddNode(projectid string, nodeid string, userid int64) (*v1.CommonResponseWizMsg, error)
	Icreated(userid int64) (*v1.IcreatedResponse, error)
	Update(id int, description string, endtime int64, name string, userid int64) (*v1.CommonResponseWizMsg, error)
	// the user id here is int not int64
	UpdateNode(projectid string, nodeid string, userid int, item *model.Items) (*v1.CommonResponseWizMsg, error)
	NodeInfo(projectid string, nodeid string, userid int) (*v1.ProjectNodeInfoResponse, error)
	Nodes(projectid string, userid int) (*v1.ProjectNodes, error)
}

type projectBiz struct {
	ds store.Istore
}

// Nodes implements ProjectBiz.
func (pb *projectBiz) Nodes(projectid string, userid int) (*v1.ProjectNodes, error) {
	resp, err := pb.ds.Projects().Nodes(projectid, userid)
	if err != nil {
		return nil, err
	}
	r := &v1.ProjectNodes{
		Nodes: *resp,
	}
	return r, nil
}

var _ ProjectBiz = (*projectBiz)(nil)

func New(ds store.Istore) *projectBiz {
	return &projectBiz{ds: ds}
}

// DeleteNode implements ProjectBiz.
func (pb *projectBiz) DeleteNode(projectid string, nodeid string, userid int64) (*v1.CommonResponseWizMsg, error) {
	_, err := pb.ds.Projects().DeleteNode(projectid, nodeid, userid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// Delete implements ProjectBiz.
func (pb *projectBiz) Delete(projectid string, userid int64) (*v1.CommonResponseWizMsg, error) {
	_, err := pb.ds.Projects().DeleteProject(projectid, userid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil

}

// Create implements ProjectBiz.
func (pb *projectBiz) Create(description string, endtime int64, name string, userid int64) (*v1.CommonResponseWizMsg, error) {
	// 将int64的时间转换成time.Time

	pwd, err := GenerateRandomProjectPwd()
	if err != nil {
		return nil, err
	}
	p := &model.Projects{
		Name:        name,
		Description: description,
		CreatedTime: time.Now(),
		EndTime:     time.Unix(0, endtime),
		Password:    pwd,
		AdminId:     userid,
	}
	_, err = pb.ds.Projects().CreateProject(p, userid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// AddNode implements ProjectBiz.
func (pb *projectBiz) AddNode(projectid string, nodeid string, userid int64) (*v1.CommonResponseWizMsg, error) {
	_, err := pb.ds.Projects().AddNode(projectid, nodeid, userid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{"success"}, nil
}

// Icreated implements ProjectBiz.
func (pb *projectBiz) Icreated(userid int64) (*v1.IcreatedResponse, error) {
	ps, err := pb.ds.Projects().GetAllProjectsICreated(userid)
	if err != nil {
		return nil, err
	}
	resp := &v1.IcreatedResponse{Projects: ps}
	return resp, nil
}

// Update implements ProjectBiz.
func (pb *projectBiz) Update(id int, description string, endtime int64, name string, userid int64) (*v1.CommonResponseWizMsg, error) {
	_, err := pb.ds.Projects().UpdateProjectInfo(id, description, endtime, name, userid)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil
}

// UpdateNode implements ProjectBiz.
func (pb *projectBiz) UpdateNode(projectid string, nodeid string, userid int, item *model.Items) (*v1.CommonResponseWizMsg, error) {
	_, err := pb.ds.Projects().UpdateNode(projectid, nodeid, userid, item)
	if err != nil {
		return nil, err
	}
	return &v1.CommonResponseWizMsg{Msg: "success"}, nil

}

// NodeInfo implements ProjectBiz.
func (pb *projectBiz) NodeInfo(projectid string, nodeid string, userid int) (*v1.ProjectNodeInfoResponse, error) {
	info, err := pb.ds.Projects().NodeInfo(projectid, nodeid, userid)
	if err != nil {
		return nil, err
	}
	nodeinfo := &v1.ProjectNodeInfoResponse{NodeInfo: *info}
	return nodeinfo, nil
}

func GenerateRandomProjectPwd() (pwd string, err error) {
	const pwdLength = 15                                                                // 密码长度
	const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 可用字符集
	b := make([]byte, pwdLength)
	_, err = rand.Read(b)
	if err != nil {
		return "", err
	}
	// 将生成的随机字节转换为字符串
	pwd = strings.Join(strings.Fields(fmt.Sprintf("%x", b)), "")
	// 确保密码长度不超过15个字符
	if len(pwd) > pwdLength {
		pwd = pwd[:pwdLength]
	}
	return pwd, nil
}
