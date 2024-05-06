package bootstrap

import (
	"apier/internal/global/errors"
	"apier/internal/global/variable"
	"apier/internal/http/validator/common/register_validator"
	"apier/internal/utils/gorm_v2"
	"apier/internal/utils/logger/sys_log_hook"
	"apier/internal/utils/logger/zap_factory"
	"apier/internal/utils/yaml_config"
	"fmt"
	"log"
)

func init() {
	fmt.Println("【Bootstrap】开始项目启动前准备...")

	// ================================================================================

	// 1. 初始化 项目根路径，参见 variable 常量包，相关路径：app\global\variable\variable.go

	//2.检查配置文件以及日志目录等非编译性的必要条件
	//checkRequiredFolders()

	//3.初始化表单参数验证器，注册在容器（Web、Api共用容器）

	fmt.Println("【Bootstrap】开始初始化参数验证器，并注册到全局容器中......")

	register_validator.WebRegisterValidator()
	//register_validator.ApiRegisterValidator()

	fmt.Println("【Bootstrap】参数验证器初始化成功")

	// ================================================================================

	fmt.Println("【Bootstrap】开始初始化配置文件载入，并配置文件指针......")

	// 4.启动针对配置文件(`config.yml`、`gorm.yml`)变化的监听， 配置文件操作指针，初始化为全局变量
	variable.ConfigYaml = yaml_config.CreateYamlFactory()
	variable.ConfigYaml.ConfigFileChangeListen()

	// config > gorm.yml 启动文件变化监听事件
	variable.ConfigGormYaml = variable.ConfigYaml.Clone("gorm")
	variable.ConfigGormYaml.ConfigFileChangeListen()

	fmt.Println("【Bootstrap】配置文件载入初始化成功")

	// ================================================================================

	fmt.Println("【Bootstrap】开始初始化全局日志句柄，并载入日志钩子处理函数......")

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.ZapLog = zap_factory.CreateZapFactory(sys_log_hook.ZapLogHandler)

	fmt.Println("【Bootstrap】全局日志句柄初始化成功")

	// ================================================================================

	// 6.根据配置初始化 gorm mysql 全局 *gorm.Db
	if variable.ConfigGormYaml.GetInt("Gorm.Mysql.IsInitGlobalGormMysql") == 1 {
		if dbMysql, err := gorm_v2.GetMysqlClient(); err != nil {
			log.Fatal(errors.ErrorsGormInitFail + err.Error())
		} else {
			variable.GormDbMysql = dbMysql
		}
	}

	// ================================================================================

	//// 7.雪花算法全局变量
	//variable.SnowFlake = snow_flake.CreateSnowflakeFactory()

	fmt.Println("【Bootstrap】项目初始化完成")

}
