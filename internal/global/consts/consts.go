package consts

// 这里定义的常量，一般是具有错误代码+错误说明组成，一般用于接口返回
const (

	// 进程被结束
	ProcessKilled string = "收到信号，进程被结束"

	// 参数验证器前缀
	ValidatorPrefix              string = "Parameters_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	// 服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误, "

	// token相关
	JwtTokenOK            int    = 200100                      // token有效
	JwtTokenInvalid       int    = 400100                      // 无效的token
	JwtTokenExpired       int    = 400101                      // 过期的token
	JwtTokenFormatErrCode int    = 400102                      // 提交的 token 格式错误
	JwtTokenFormatErrMsg  string = "提交的 token 格式错误"            // 提交的 token 格式错误
	JwtTokenMustValid     string = "token为必填项，请在请求header部分提交!" // 提交的 token 格式错误

	// 接口请求常用业务状态码
	RequestStatusOkCode int    = 200
	RequestStatusOkMsg  string = "Success"

	RequestCreatFailCode int    = 400200
	RequestCreatFailMsg  string = "新增失败"

	RequestUpdateFailCode int    = 400201
	RequestUpdateFailMsg  string = "更新失败"

	RequestDeleteFailCode int    = 400202
	RequestDeleteFailMsg  string = "删除失败"

	RequestSelectFailCode int    = 400203
	RequestSelectFailMsg  string = "查询无数据"

	RequestRegisterFailCode int    = 400204
	RequestRegisterFailMsg  string = "注册失败"

	RequestLoginFailCode int    = 400205
	RequestLoginFailMsg  string = "登录失败"

	RequestRefreshTokenFailCode int    = -400206
	RequestRefreshTokenFailMsg  string = "刷新Token失败"
)
