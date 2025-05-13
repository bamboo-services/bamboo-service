// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxySubscription is the golang structure for table proxy_subscription.
type ProxySubscription struct {
	SubscriptionUuid string      `json:"subscription_uuid" orm:"subscription_uuid" description:"代理订阅地址唯一标识符"`      // 代理订阅地址唯一标识符
	ProxyGroupUuid   string      `json:"proxy_group_uuid"  orm:"proxy_group_uuid"  description:"代理组唯一标识符"`         // 代理组唯一标识符
	UserUuid         string      `json:"user_uuid"         orm:"user_uuid"         description:"用户唯一标识符"`          // 用户唯一标识符
	Name             string      `json:"name"              orm:"name"              description:"代理订阅地址名称"`         // 代理订阅地址名称
	Merchant         string      `json:"merchant"          orm:"merchant"          description:"代理订阅地址商户"`         // 代理订阅地址商户
	Description      string      `json:"description"       orm:"description"       description:"代理订阅地址描述"`         // 代理订阅地址描述
	Url              string      `json:"url"               orm:"url"               description:"代理订阅地址URL"`        // 代理订阅地址URL
	OriginalContent  string      `json:"original_content"  orm:"original_content"  description:"代理订阅地址所订阅获取的原始内容"` // 代理订阅地址所订阅获取的原始内容
	CreatedAt        *gtime.Time `json:"created_at"        orm:"created_at"        description:"代理订阅地址创建时间"`       // 代理订阅地址创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"代理订阅地址更新时间"`       // 代理订阅地址更新时间
	SubscribeAt      *gtime.Time `json:"subscribe_at"      orm:"subscribe_at"      description:"代理订阅地址订阅时间"`       // 代理订阅地址订阅时间
}
