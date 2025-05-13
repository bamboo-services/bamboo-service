package v1

import (
	"bamboo-service/internal/model/dto"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
)

type AuthRegisterReq struct {
	g.Meta          `path:"/auth/register" method:"Post" sm:"用户注册" tags:"授权控制器" dc:"用于普通用户进行系统注册使用"`
	Username        string `json:"username" v:"required|length:4,32|regex:^[0-9A-Za-z-\\_]+$#用户名不能为空|用户名长度必须在4到32个字符之间|用户名格式不正确(只允许英文、数字、下划线、横线)" dc:"用户名"`
	Email           string `json:"email" v:"required|email#请输入邮箱|邮箱格式错误" dc:"邮箱"`
	Phone           string `json:"phone" v:"phone#手机号格式不正确" dc:"手机号"`
	Password        string `json:"password" v:"required#请输入密码" dc:"密码"`
	ConfirmPassword string `json:"confirm_password" v:"required|same:password#请输入确认密码|两次密码不一致" dc:"确认密码"`
}

type AuthRegisterRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*dto.UserInfoDTO]
}
