package main

import (
	"fmt"
	"go_poetry/routers"
)

func main() {
	// 加载多个APP的路由配置
	routers.Include(routers.ApiRouters, routers.WebRouters)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
