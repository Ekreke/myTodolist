package t

import (
	"github.com/ekreke/myTodolist/internal/mytodolist/store"
	"github.com/gin-gonic/gin"
)

// interface
/*
	tg.Use(middleware.Authn())
	tg.Use(middleware.Authz(authz))
	// test graceful shutdown , send a request and get response delay 10 seconds
	tg.GET("/lazy", tc.Lazy)
	// get logs from remote use grpc server
	tg.GET("/logs", tc.Logs)
	// get user list remote use grpc server
	tg.GET("/users", tc.GetUsers)
	// get casbin: user/rules remote use grpc server
	tg.GET("/rules", tc.rules)
*/
// controller interface
type ITestController interface {
	// lazy
	Lazy(ctx *gin.Context)
	// logs
	Logs(ctx *gin.Context)
	// get users
	GetUsers(ctx *gin.Context)
	// get casbin: user/rules
	Rules(ctx *gin.Context)
}

type TestController struct {
}

// GetUsers implements ITestController.
func (ctrl *TestController) GetUsers(ctx *gin.Context) {
	panic("unimplemented")
}

// Logs implements ITestController.
func (ctrl *TestController) Logs(ctx *gin.Context) {
	panic("unimplemented")
}

// Rules implements ITestController.
func (*TestController) Rules(ctx *gin.Context) {
	panic("unimplemented")
}

func New(ds store.Istore) *TestController {
	return &TestController{}
}

var _ ITestController = (*TestController)(nil)
