package model

import (
	"apier/internal/global/variable"
	"apier/internal/utils/encryption"
	"go.uber.org/zap"
)

func CreateSuperAdminFactory() *SuperAdminModel {
	return &SuperAdminModel{BaseModel: BaseModel{DB: UseDbConn("mysql")}}
}

type SuperAdminModel struct {
	BaseModel
	Username string `gorm:"unique"`
	Password string
}

// TableName 表名
func (u *SuperAdminModel) TableName() string {
	return "super_admin"
}

// SuperAdminRegister 超级管理员注册
func (u *SuperAdminModel) SuperAdminRegister(userName string, password string) bool {
	sql := "INSERT INTO `super_admin` (`username`, `password`) " +
		"SELECT ? AS `username`, ? AS `password` FROM DUAL WHERE NOT EXISTS (" +
		"SELECT 1 FROM `super_admin` WHERE `username` = ?" +
		");"
	result := u.Exec(sql, userName, password, userName)
	if result.RowsAffected > 0 {
		return true
	} else {
		return false
	}
}

// SuperAdminLogin 超级管理员登录
func (sa *SuperAdminModel) SuperAdminLogin(userName string, pass string) *SuperAdminModel {
	sql := "select id, user_name,real_name,pass,phone  from tb_users where  user_name=?  limit 1"
	result := sa.Raw(sql, userName).First(sa)
	if result.Error == nil {
		// 账号密码验证成功
		if len(sa.Password) > 0 && (sa.Password == encryption.Base64Md5(pass)) {
			return sa
		}
	} else {
		variable.ZapLog.Error("根据账号查询单条记录出错:", zap.Error(result.Error))
	}
	return nil
}
