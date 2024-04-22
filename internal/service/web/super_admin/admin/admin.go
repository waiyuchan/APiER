package admin

import (
	"apier/internal/model"
	"apier/internal/utils/encryption"
	"fmt"
)

func CreateSuperAdminFactory() *SuperAdminDao {
	fmt.Println("Create Super Admin Factory")
	return &SuperAdminDao{model.CreateSuperAdminFactory()}
}

type SuperAdminDao struct {
	superAdminModel *model.SuperAdminModel
}

func (sa *SuperAdminDao) Register(username string, password string) bool {
	fmt.Println(password)
	password = encryption.Base64Md5(password)
	return sa.superAdminModel.SuperAdminRegister(username, password)
}
