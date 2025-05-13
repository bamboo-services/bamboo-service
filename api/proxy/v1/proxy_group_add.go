package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

type ProxyGroupAddReq struct {
	g.Meta      `path:"/proxy/group" method:"Post" sm:"添加代理组" tags:"代理控制器" dc:"用于添加代理组"`
	Name        string `json:"name" v:"required|max-length:64#请输入代理组名称|代理组名称长度不能超过64个字符" dc:"请输入代理组名称"`
	Description string `json:"description" v:"length:0,255#代理组描述长度不能超过255个字符" dc:"代理组描述"`
}

type ProxyGroupAddRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil]
}
