// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxySubscription is the golang structure of table fy_proxy_subscription for DAO operations like Where/Data.
type ProxySubscription struct {
	g.Meta           `orm:"table:fy_proxy_subscription, do:true"`
	SubscriptionUuid interface{} // 代理订阅地址唯一标识符
	ProxyGroupUuid   interface{} // 代理组唯一标识符
	UserUuid         interface{} // 用户唯一标识符
	Name             interface{} // 代理订阅地址名称
	Merchant         interface{} // 代理订阅地址商户
	Description      interface{} // 代理订阅地址描述
	Url              interface{} // 代理订阅地址URL
	OriginalContent  interface{} // 代理订阅地址所订阅获取的原始内容
	CreatedAt        *gtime.Time // 代理订阅地址创建时间
	UpdatedAt        *gtime.Time // 代理订阅地址更新时间
	SubscribeAt      *gtime.Time // 代理订阅地址订阅时间
}
