package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Info(ctx *gin.Context) {
	log.C(ctx).Infow("Item Info function called")

}
