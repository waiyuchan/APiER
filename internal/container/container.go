package container

import (
	"apier/internal/global/custom_errors"
	"apier/internal/global/variable"
	"log"
	"strings"
	"sync"
)

// 定义一个全局键值对存储容器
var syncMap sync.Map

// CreateContainersFactory 创建一个容器工厂
func CreateContainersFactory() *containers {
	return &containers{}
}

// 定义一个容器结构体
type containers struct {
}

// Set  以键值对的形式将代码注册到容器
func (c *containers) Set(key string, value interface{}) (res bool) {
	if _, exists := c.KeyIsExists(key); exists == false {
		syncMap.Store(key, value)
		res = true
	} else {
		// 程序启动阶段，zaplog 未初始化，使用系统log打印启动时候发生的异常日志
		if variable.ZapLog == nil {
			log.Fatal(custom_errors.ErrorsContainerKeyAlreadyExists + "，请解决键名重复问题，相关键：" + key)
		} else {
			// 程序启动初始化完成
			variable.ZapLog.Warn(custom_errors.ErrorsContainerKeyAlreadyExists + "，相关键：" + key)
		}
	}
	return
}

// Delete  删除
func (c *containers) Delete(key string) {
	syncMap.Delete(key)
}

// Get 传递键，从容器获取值
func (c *containers) Get(key string) interface{} {
	if value, exists := c.KeyIsExists(key); exists {
		return value
	}
	return nil
}

// KeyIsExists 判断键是否被注册
func (c *containers) KeyIsExists(key string) (interface{}, bool) {
	return syncMap.Load(key)
}

// FuzzyDelete 按照键的前缀模糊删除容器中注册的内容
func (c *containers) FuzzyDelete(keyPre string) {
	syncMap.Range(func(key, value interface{}) bool {
		if keyName, ok := key.(string); ok {
			if strings.HasPrefix(keyName, keyPre) {
				syncMap.Delete(keyName)
			}
		}
		return true
	})
}
