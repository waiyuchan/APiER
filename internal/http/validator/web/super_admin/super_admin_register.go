package super_admin

import (
	"apier/internal/global/consts"
	"apier/internal/http/controller/web"
	"apier/internal/http/validator/data_transfer"
	"apier/internal/utils/response"
	"github.com/gin-gonic/gin"
)

type SuperAdminRegister struct {
	UserName string `format:"username" json:"username" binding:"required,min=1"`        // 必填、对于文本,表示它的长度>=1
	Password string `format:"password" json:"password" binding:"required,min=6,max=20"` //  密码为 必填，长度>=6
}

func (sa SuperAdminRegister) CheckParams(context *gin.Context) {

	// 先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&sa); err != nil {
		response.ValidatorError(context, err)
		return
	}
	// 该函数主要是将本结构体的字段（成员）按照 consts.ValidatorPrefix+ json标签对应的 键 => 值 形式绑定在上下文，便于下一步（控制器）可以直接通过 context.Get(键) 获取相关值
	extraAddBindDataContext := data_transfer.DataAddContext(sa, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "UserRegister表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&web.SuperAdmin{}).SuperAdminRegister(extraAddBindDataContext)
	}
}
