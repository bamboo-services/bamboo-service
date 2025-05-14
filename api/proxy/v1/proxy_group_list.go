package v1

import (
	"bamboo-service/internal/model/dto"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
)

type ProxyGroupListReq struct {
	g.Meta   `path:"/proxy/group/list" method:"Get" sm:"获取代理组列表" tags:"代理控制器" dc:"用于获取简单版本的代理组列表「主要用于选择框使用」"`
	UserUUID string `json:"user_uuid" v:"regex:^(|[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})$#请输入正确的用户UUID" dc:"用户UUID"`
	Search   string `json:"search" dc:"搜索内容"`
}

type ProxyGroupListRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*[]*dto.ProxyBaseGroupLiteDTO]
}
