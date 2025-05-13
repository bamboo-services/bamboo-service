package v1

import (
	"bamboo-service/internal/model/dto"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
)

type ProxyGroupPageReq struct {
	g.Meta `path:"/proxy/group/page" method:"Get" sm:"订阅组列表（分页）" tags:"代理控制器" dc:"用户可以通过此接口查询到自己创建的代理组有什么内容（分页）"`
	Page   int    `json:"page" default:"1" v:"min:1#请输入正确的页数" dc:"页数"`
	Size   int    `json:"size" default:"20" v:"size:1,100#请输入正确的条数(1-100)" dc:"条数"`
	Search string `json:"search" dc:"搜索内容"`
}

type ProxyGroupPageRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*[]*dto.ProxyBaseGroupDTO]
}
