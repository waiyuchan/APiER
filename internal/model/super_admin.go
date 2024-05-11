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
func (sam *SuperAdminModel) SuperAdminRegister(username string, encryptedPassword string) bool {
	variable.ZapLog.Info("准备写入数据库...")

	// 检查用户名是否已经存在
	var existingUser SuperAdminModel
	if result := sam.DB.Where("username = ?", username).First(&existingUser); errors.Is(result.Error, gorm.ErrRecordNotFound) {
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
		Username: username,
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
func (sam *SuperAdminModel) SuperAdminLogin(username string, password string) *SuperAdminModel {
	// 修改SQL查询，包括密码字段
	sql := "SELECT id, username, password FROM super_admin WHERE username=? LIMIT 1"
	result := sam.DB.Raw(sql, username).Scan(sam)
	if result.Error != nil {
		variable.ZapLog.Error("根据用户名查询出错:", zap.Error(result.Error))
		return nil
	}

	// 账号密码验证成功
	if sam.Password != "" && encryption.Base64Md5(password) == sam.Password {
		return sam
	}

	// 如果密码不匹配，也可以在这里记录日志或返回错误
	variable.ZapLog.Error("密码验证失败")
	return nil
}
