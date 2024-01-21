package item

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *ItemController) Delete(ctx *gin.Context) {
	log.C(ctx).Infow("Item delete function called")

}
