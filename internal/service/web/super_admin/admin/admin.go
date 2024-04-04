package admin

import (
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
	password = encryption.Base64Md5(password)
	return sa.superAdminModel.SuperAdminRegister(username, password)
}
