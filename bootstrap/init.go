package bootstrap

import (
	"apier/internal/global/errors"
	"apier/internal/global/variable"
	"apier/internal/utils/gorm_v2"
	"apier/internal/utils/logger/sys_log_hook"
	"apier/internal/utils/logger/zap_factory"
	"apier/internal/utils/yaml_config"
	"log"
)

func init() {
	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go

	//2.检查配置文件以及日志目录等非编译性的必要条件
	//checkRequiredFolders()

	//3.初始化表单参数验证器，注册在容器（Web、Api共用容器）
	//register_validator.WebRegisterValidator()
	//register_validator.ApiRegisterValidator()

	// 4.启动针对配置文件(confgi.yml、gorm_v2.yml)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYaml = yaml_config.CreateYamlFactory()
	variable.ConfigYaml.ConfigFileChangeListen()
	// config>gorm_v2.yml 启动文件变化监听事件
	variable.ConfigGormYaml = variable.ConfigYaml.Clone("gorm_v2")
	variable.ConfigGormYaml.ConfigFileChangeListen()

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	if variable.ConfigGormYaml.GetInt("Gormv2.Mysql.IsInitGlobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetMysqlClient(); err != nil {
			log.Fatal(errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = dbMysql
		}
	}

	//// 7.雪花算法全局变量
	//variable.SnowFlake = snow_flake.CreateSnowflakeFactory()
	//
	//// 8.websocket Hub中心启动
	//if variable.ConfigYaml.GetInt("Websocket.Start") == 1 {
	//	// websocket 管理中心hub全局初始化一份
	//	variable.WebsocketHub = core.CreateHubFactory()
	//	if Wh, ok := variable.WebsocketHub.(*core.Hub); ok {
	//		go Wh.Run()
	//	}
	//}
	//
	//// 9.casbin 依据配置文件设置参数(IsInit=1)初始化
	//if variable.ConfigYaml.GetInt("Casbin.IsInit") == 1 {
	//	var err error
	//	if variable.Enforcer, err = casbin_v2.InitCasbinEnforcer(); err != nil {
	//		log.Fatal(err.Error())
	//	}
	//}
	////10.全局注册 validator 错误翻译器,zh 代表中文，en 代表英语
	//if err := validator_translation.InitTrans("zh"); err != nil {
	//	log.Fatal(my_errors.ErrorsValidatorTransInitFail + err.Error())
	//}
}
