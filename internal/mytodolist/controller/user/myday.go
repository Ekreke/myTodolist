package user

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Myday(c *gin.Context) {
	log.C(c).Infow("Myday function called")
}
