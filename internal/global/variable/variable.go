package variable

import (
	"apier/internal/global/errors"
	"apier/internal/utils/yaml_config/yaml_config_interface"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var (
	BasePath        string
	DateFormat      = "2006-01-02 15:04:05" //  设置全局日期时间格式
	ConfigKeyPrefix = "Config_"             //  配置文件键值缓存时，键的前缀

	// 全局日志指针
	ZapLog *zap.Logger

	// 全局配置文件
	ConfigYaml     yaml_config_interface.YamlConfigInterface // 全局配置文件指针
	ConfigGormYaml yaml_config_interface.YamlConfigInterface // 全局配置文件指针

	// gorm 数据库客户端，如果您操作数据库使用的是gorm，请取消以下注释，在 bootstrap>init 文件，进行初始化即可使用
	GormDbMysql *gorm.DB // 全局gorm的客户端连接

)

func init() {

	// 1.初始化程序根目录
	if curPath, err := os.Getwd(); err == nil {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(curPath, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = curPath
		}
	} else {
		log.Fatal(errors.ErrorsBasePath)
	}

}
