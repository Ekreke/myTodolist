package t

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl *TestController) Lazy(ctx *gin.Context) {
	time.Sleep(10 * time.Second)
	ctx.JSON(200, gin.H{"msg": "ok"})
}
