package main

import (
	"apier/internal/db"
	"apier/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init() // 初始化数据库连接

	r := gin.Default()
	// 设置路由...

	routes.RegisterRoutes(r)

	r.Run() // 在localhost:8080上启动服务
}
