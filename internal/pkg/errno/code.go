package errno


var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "ok"}
	ParamError          = &Errno{Code: 400, Message: "参数错误"}
	UnAuthorizedError   = &Errno{Code: 401, Message: "未经授权"}
	NoPermissionError	= &Errno{Code: 403, Message: "无权限访问"}
	NotFoundError 		= &Errno{Code: 404, Message: "请求API不存在"}
	InternalServerError = &Errno{Code: 500, Message: "内部服务器错误"}

	ValidationError         = &Errno{Code: 20001, Message: "验证失败"}
	DatabaseError           = &Errno{Code: 20002, Message: "数据库错误"}

	// user errors
	UserNotFoundError          = &Errno{Code: 20102, Message: "找不到用户"}
	GenTokenError          	   = &Errno{Code: 20101, Message: "生成token出错"}
	TokenInvalidError          = &Errno{Code: 20103, Message: "token无效"}
	PasswordIncorrectError     = &Errno{Code: 20104, Message: "密码不正确"}
	AreaCodeEmptyError         = &Errno{Code: 20105, Message: "手机区号不能为空"}
	PhoneEmptyError            = &Errno{Code: 20106, Message: "手机号不能为空"}
	GenVCodeError              = &Errno{Code: 20107, Message: "生成验证码错误"}
	SendSMSError               = &Errno{Code: 20108, Message: "发送短信错误"}
	SendSMSTooManyError        = &Errno{Code: 20109, Message: "已超出当日限制，请明天再试"}
	VerifyCodeError            = &Errno{Code: 20110, Message: "验证码错误"}
	EmailOrPasswordError       = &Errno{Code: 20111, Message: "邮箱或密码错误"}
	TwicePasswordNotMatchError = &Errno{Code: 20112, Message: "两次密码输入不一致"}
	RegisterFailedError        = &Errno{Code: 20113, Message: "注册失败"}
)
