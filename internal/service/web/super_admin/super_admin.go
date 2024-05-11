package super_admin

import (
	"apier/internal/global/variable"
	"apier/internal/model"
	"apier/internal/utils/encryption"
)

func CreateSuperAdminFactory() *SuperAdminDao {
	return &SuperAdminDao{model.CreateSuperAdminFactory()}
}

type SuperAdminDao struct {
	superAdminModel *model.SuperAdminModel
}

func (sa *SuperAdminDao) Register(username string, password string) bool {
	encryptedPassword := encryption.Base64Md5(password)
	variable.ZapLog.Info("加密后的密码：" + encryptedPassword)
	return sa.superAdminModel.SuperAdminRegister(username, encryptedPassword)
}

func (sa *SuperAdminDao) Login(username string, password string) *model.SuperAdminModel {
	return sa.superAdminModel.SuperAdminLogin(username, password)
}
