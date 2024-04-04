package main

import (
	_ "apier/bootstrap"
	"apier/internal/global/variable"
	"apier/routes"
)

// 后台管理系统的后台服务入口
func main() {
	router := routes.InitWebRouter()
	_ = router.Run(variable.ConfigYaml.GetString("HttpServer.Web.Port"))
}
