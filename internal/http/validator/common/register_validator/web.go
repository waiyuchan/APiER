package register_validator

import (
	"apier/internal/container"
	"apier/internal/global/consts"
	"apier/internal/http/validator/web/super_admin"
)

// 各个业务模块验证器必须进行注册（初始化），程序启动时会自动加载到容器
func WebRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	//  key 按照前缀+模块+验证动作 格式，将各个模块验证注册在容器
	var key string

	// SuperAdmin 模块参数验证器注册
	key = consts.ValidatorPrefix + "SuperAdminRegister"
	containers.Set(key, super_admin.SuperAdminRegister{})

	key = consts.ValidatorPrefix + "SuperAdminLogin"
	containers.Set(key, super_admin.SuperAdminLogin{})

	//// Users 模块表单验证器按照 key => value 形式注册在容器，方便路由模块中调用
	//key = consts.ValidatorPrefix + "UsersRegister"
	//containers.Set(key, users.Register{})
	//key = consts.ValidatorPrefix + "UsersLogin"
	//containers.Set(key, users.Login{})
	//key = consts.ValidatorPrefix + "RefreshToken"
	//containers.Set(key, users.RefreshToken{})

	// Users基本操作（CURD）
	//key = consts.ValidatorPrefix + "UsersShow"
	//containers.Set(key, users.Show{})
	//key = consts.ValidatorPrefix + "UsersStore"
	//containers.Set(key, users.Store{})
	//key = consts.ValidatorPrefix + "UsersUpdate"
	//containers.Set(key, users.Update{})
	//key = consts.ValidatorPrefix + "UsersDestroy"
	//containers.Set(key, users.Destroy{})

}
