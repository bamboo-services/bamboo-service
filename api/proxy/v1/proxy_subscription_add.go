package v1

import (
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"go/types"
)

type ProxySubscriptionAddReq struct {
	g.Meta         `path:"/proxy/subscription" method:"Post" sm:"添加订阅" tags:"代理控制器" dc:"用于向某一个代理组添加订阅内容"`
	ProxyGroupUuid string `json:"proxy_group_uuid" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$#请输入代理组UUID|代理组UUID格式不正确" dc:"代理组UUID"`
	Name           string `json:"name" v:"required#请输入订阅名称" dc:"请输入订阅名称"`
	Merchant       string `json:"merchant" v:"required#请输入商户名称" dc:"请输入商户名称"`
	Url            string `json:"url" v:"required#请输入订阅URL" dc:"请输入订阅URL"`
	Description    string `json:"description" v:"max-length:255#订阅描述长度不能超过255个字符" dc:"订阅描述"`
}

type ProxySubscriptionAddRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*types.Nil]
}
