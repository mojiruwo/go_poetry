package routers

import (
	"github.com/gin-gonic/gin"
	"go_poetry/controller/web"
)

func WebRouters(e *gin.Engine) {
	e.GET("/", web.Poetry)
	e.GET("/web/index", web.Poetry)
}
