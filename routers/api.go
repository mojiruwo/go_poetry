package routers

import (
	"github.com/gin-gonic/gin"
	"go_poetry/controller"
)

func ApiRouters(e *gin.Engine){
	e.GET("/api/index", controller.index())
}
