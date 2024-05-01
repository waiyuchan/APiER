package web

import (
	"apier/internal/db"
	"apier/internal/global/consts"
	"apier/internal/global/variable"
	"apier/internal/model"
	"apier/internal/service/web/super_admin/admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAdminInput struct {
	Username string `json:"username"`
	Password string `json:"password"` // 注意：实际应用中应先加密再存储
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type SuperAdmin struct {
}

func (sa *SuperAdmin) SuperAdminRegister(context *gin.Context) {
	username := context.GetString(consts.ValidatorPrefix + "username")
	password := context.GetString(consts.ValidatorPrefix + "password")

	fmt.Println(username)
	fmt.Println(password)

	admin.CreateSuperAdminFactory().Register(username, password)

	//if admin.CreateSuperAdminFactory().Register(username, password) {
	//	response.Success(context, consts.RequestStatusOkMsg, "")
	//} else {
	//	response.Fail(context, consts.RequestRegisterFailCode, consts.RequestRegisterFailMsg, "")
	//}

}

func (sa *SuperAdmin) SuperAdminLogin(context *gin.Context) {

	variable.ZapLog.Info("基本的运行提示类信息")

	userName := context.GetString(consts.ValidatorPrefix + "username")
	password := context.GetString(consts.ValidatorPrefix + "password")

	superAdminModelFact := model.CreateSuperAdminFactory()
	superAdminModel := superAdminModelFact.SuperAdminLogin(userName, password)

	fmt.Println(superAdminModel)
	variable.ZapLog.Info("基本的运行提示类信息")

	//var input LoginInput
	//if err := c.ShouldBindJSON(&input); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	//	return
	//}
	//
	//// 这里添加登录逻辑，比如检查用户名和密码是否匹配
	//// 假设超级管理员用户名和密码分别为admin和password（实际开发中需要更安全的验证机制）
	//
	//if input.Username == "admin" && input.Password == "password" {
	//	c.JSON(http.StatusOK, gin.H{"notification": "Login successful"})
	//} else {
	//	c.JSON(http.StatusUnauthorized, gin.H{"errors": "Incorrect username or password"})
	//}
}

func CreateAdmin(c *gin.Context) {
	// 绑定输入数据
	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	fmt.Println(input.Username)
	// 创建管理员记录
	admin := model.Admin{Username: input.Username, Password: input.Password, Email: input.Email, Role: input.Role}
	result := db.DB.Create(&admin) // `db` 是*gorm.DB类型的全局变量
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admin})
}

func ListAdmins(c *gin.Context) {
	var admins []model.Admin
	result := db.DB.Find(&admins) // `db` 是*gorm.DB类型的全局变量
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": admins})
}
