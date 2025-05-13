package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

type AuthLogoutReq struct {
	g.Meta `path:"/auth/logout" method:"Get" sm:"用户登出" tags:"授权控制器" dc:"用于登出账号使用"`
}

type AuthLogoutRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil]
}
