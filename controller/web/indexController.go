package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Poetry(c *gin.Context){
	c.HTML(http.StatusOK, "web/index.html", gin.H{"title": "自动生成古诗"})
}
