package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Update(ctx *gin.Context) {
	log.C(ctx).Infow("Item Update function called")

}
