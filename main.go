package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "go_poetry/config"
	"go_poetry/routers"
)
func main() {
	// 加载多个APP的路由配置
	routers.Include(routers.ApiRouters, routers.WebRouters)
	// 初始化路由
	r := routers.Init()
	port := viper.GetString("shell.port")
	if err := r.Run(port); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
