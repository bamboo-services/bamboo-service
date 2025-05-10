package v1

import (
	"bamboo-service/internal/model/dto/dsingle"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
)

type AuthLoginReq struct {
	g.Meta   `path:"/auth/login" method:"Post" tags:"授权控制器" sm:"用户登录" dc:"用户可以该接口进行授权登录"`
	Username string `json:"username" v:"required|length:4,32|regex:^[0-9A-Za-z-\\_]+$#用户名不能为空|用户名长度必须在4到32个字符之间|用户名格式不正确(只允许英文、数字、下划线、横线)"`
	Password string `json:"password" v:"required#请输入密码"`
}

type AuthLoginRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*dsingle.UserLoginDTO]
}
