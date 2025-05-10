package custom

import "github.com/XiaoLFeng/bamboo-utils/berror"

var (
	ErrorUserNotExist          = berror.NewErrorCode(40404, "用户不存在", nil)
	ErrorUserExist             = berror.NewErrorCode(40405, "用户已存在", nil)
	ErrorUserPasswordIncorrect = berror.NewErrorCode(40009, "用户密码错误", nil)
)
