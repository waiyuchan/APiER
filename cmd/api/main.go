package main

import (
	"apier/internal/db"
	"apier/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库连接
	db.Init()

	r := gin.Default()

	// 路由注册
	routes.RegisterRoutes(r)

	// 启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
