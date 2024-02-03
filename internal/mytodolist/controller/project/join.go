package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (pc *ProjectController) Join(ctx *gin.Context) {
	log.C(ctx).Infow("project join function called")

}
