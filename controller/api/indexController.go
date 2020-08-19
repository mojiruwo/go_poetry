package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_poetry/utils"
	"net/http"
	"os/exec"
)

func Buildpoetry(c *gin.Context){
	name := c.DefaultQuery("name", "")
	command := "D:\\study\\ai_poetry\\build.py"
	cmd := exec.Command("python", command, name)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	txt, _ := utils.GbkToUtf8(output)
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(txt))
	c.JSON(http.StatusOK, gin.H{"txt": string(txt)})
}
