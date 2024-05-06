package model

import (
	"apier/internal/global/variable"
	"apier/internal/utils/encryption"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func CreateSuperAdminFactory() *SuperAdminModel {
	fmt.Println("model create super admin factory")
	return &SuperAdminModel{BaseModel: BaseModel{DB: UseDbConn("mysql")}}
}

type SuperAdminModel struct {
	BaseModel
	Username string `gorm:"unique"`
	Password string
}

// TableName 表名
func (sam *SuperAdminModel) TableName() string {
	return "super_admin"
}

// SuperAdminRegister 超级管理员注册
func (sam *SuperAdminModel) SuperAdminRegister(userName string, encryptedPassword string) bool {
	variable.ZapLog.Info("准备写入数据库...")

	// 检查用户名是否已经存在
	var existingUser SuperAdminModel
	if result := sam.DB.Where("username = ?", userName).First(&existingUser); errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 如果用户不存在，则继续创建新用户
	} else if result.Error != nil {
		// 如果发生了其他错误，记录错误并返回
		variable.ZapLog.Error("查询用户失败", zap.Error(result.Error))
		return false
	} else {
		// 如果用户已经存在，则返回false表示注册失败
		variable.ZapLog.Info("用户名已经存在")
		return false
	}

	// 创建新用户
	newUser := SuperAdminModel{
		Username: userName,
		Password: encryptedPassword,
	}
	if err := sam.DB.Create(&newUser).Error; err != nil {
		variable.ZapLog.Error("插入数据失败", zap.Error(err))
		return false
	}

	// 返回true表示注册成功
	return true
}

// SuperAdminLogin 超级管理员登录
func (sam *SuperAdminModel) SuperAdminLogin(userName string, pass string) *SuperAdminModel {
	sql := "select id, user_name,real_name,pass,phone  from tb_users where  user_name=?  limit 1"
	result := sam.Raw(sql, userName).First(sam)
	if result.Error == nil {
		// 账号密码验证成功
		if len(sam.Password) > 0 && (sam.Password == encryption.Base64Md5(pass)) {
			return sam
		}
	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}
