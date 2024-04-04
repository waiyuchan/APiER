package factory

import (
	"apier/internal/container"
	"apier/internal/global/errors"
	"apier/internal/global/variable"
	"apier/internal/http/validator/validator_interface"
	"github.com/gin-gonic/gin"
)

// Create 表单参数验证器工厂
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(validator_interface.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
