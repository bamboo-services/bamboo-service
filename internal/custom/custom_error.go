package custom

import "github.com/XiaoLFeng/bamboo-utils/berror"

var (
	//
	// 401 Unauthorized errors (40101-40199)
	//

	ErrorUserPasswordIncorrect        = berror.NewErrorCode(40106, "用户密码错误", nil)
	ErrMailCodeHasExpired             = berror.NewErrorCode(40107, "邮件验证码已过期", nil)
	ErrorUserConfirmPasswordIncorrect = berror.NewErrorCode(40108, "两次输入的密码不一致", nil)

	//
	// 403 Forbidden errors (40301-40399)
	//

	ErrorSystemNotAbleRegister   = berror.NewErrorCode(40304, "系统不允许注册", nil)
	ErrMailCodeSentTooFrequently = berror.NewErrorCode(40305, "邮件验证码发送过于频繁", nil)
	ErrorMailSendFailed          = berror.NewErrorCode(40306, "邮件发送失败", nil)
	ErrorMailCodeIncorrect       = berror.NewErrorCode(40307, "邮件验证码错误", nil)
	ErrorEmailNotVerify          = berror.NewErrorCode(40308, "邮箱未验证，请联系管理员", nil)

	//
	// 404 Not Found errors (40401-40499)
	//

	ErrorUserNotExist     = berror.NewErrorCode(40404, "用户不存在", nil)
	ErrorUserExist        = berror.NewErrorCode(40405, "用户已存在", nil)
	ErrorMailCodeNotExist = berror.NewErrorCode(40406, "邮件验证码不存在", nil)

	//
	// 500 Internal Server errors (50001-50099)
	//

	ErrorMailConfigFailed = berror.NewErrorCode(50011, "邮件配置失败", nil)
)
