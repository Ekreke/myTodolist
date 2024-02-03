package project

import (
	"github.com/ekreke/myTodolist/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// Myprojects implements IProjectController.
func (*ProjectController) Myprojects(ctx *gin.Context) {
	log.C(ctx).Infow("project myprojects function called")
}
