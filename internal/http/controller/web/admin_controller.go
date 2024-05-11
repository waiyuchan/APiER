package web

import (
	"apier/internal/db"
	"apier/internal/global/consts"
	"apier/internal/global/variable"
	"apier/internal/model"
	"apier/internal/service/web/super_admin"
	"apier/internal/utils/response"
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

// 超级管理员注册
func (sa *SuperAdmin) SuperAdminRegister(context *gin.Context) {
	username := context.GetString(consts.ValidatorPrefix + "username")
	password := context.GetString(consts.ValidatorPrefix + "password")

	variable.ZapLog.Info("用户注册账号：" + username + "，用户注册密码：" + password)

	if super_admin.CreateSuperAdminFactory().Register(username, password) {
		response.Success(context, consts.RequestStatusOkMsg, "")
	} else {
		response.Fail(context, consts.RequestRegisterFailCode, consts.RequestRegisterFailMsg, "")
	}

}

// 超级管理员登录
func (sa *SuperAdmin) SuperAdminLogin(context *gin.Context) {
	username := context.GetString(consts.ValidatorPrefix + "username")
	password := context.GetString(consts.ValidatorPrefix + "password")

	variable.ZapLog.Info("用户登录账号：" + username + "，用户登录密码：" + password)

	if super_admin.CreateSuperAdminFactory().Login(username, password) != nil {
		response.Success(context, "Login successful", "")
	} else {
		response.Fail(context, consts.RequestLoginFailCode, "Incorrect username or password", gin.H{"token": ""})
	}
}

func CreateAdmin(c *gin.Context) {
	// 绑定输入数据
	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	fmt.Println(input.Username)
	fmt.Println(input.Password)
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
