package routers

import (
	"github.com/gin-gonic/gin"
	"go_poetry/controller/api"
)

func ApiRouters(e *gin.Engine){
	e.GET("/api/index", api.Buildpoetry)
}
