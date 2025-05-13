package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

type AuthResetPasswordReq struct {
	g.Meta          `path:"/auth/password/reset" method:"Post" sm:"重置密码" tags:"授权控制器" dc:"用于用户重置密码「需要邮件验证码」"`
	Email           string `json:"email" v:"required|email#请输入邮箱|邮箱格式错误" dc:"请输入邮箱"`
	Code            string `json:"code" v:"required#请输入验证码" dc:"请输入验证码"`
	Password        string `json:"password" v:"required#请输入密码" dc:"请输入密码"`
	ConfirmPassword string `json:"confirm_password" v:"required|same:password#请输入确认密码|两次密码不一致" dc:"请输入确认密码"`
}

type AuthResetPasswordRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil]
}
