package collection

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (cc *CollectionController) Update(ctx *gin.Context) {
	log.C(ctx).Infow("collection create function called")
	username := ctx.GetString("X-Username")
	log.Debugw("the username is:", "username:", username)

}
