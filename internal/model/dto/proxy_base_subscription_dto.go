package dto

import "github.com/gogf/gf/v2/os/gtime"

// ProxyBaseSubscriptionDTO 表示代理订阅地址的基础信息。
//
// 包括唯一标识符、代理组、用户、订阅名称、商户、描述、URL，以及创建、更新和订阅时间等字段。
type ProxyBaseSubscriptionDTO struct {
	SubscriptionUuid string      `json:"subscription_uuid" orm:"subscription_uuid" description:"代理订阅地址唯一标识符"`
	ProxyGroupUuid   string      `json:"proxy_group_uuid"  orm:"proxy_group_uuid"  description:"代理组唯一标识符"`
	UserUuid         string      `json:"user_uuid"         orm:"user_uuid"         description:"用户唯一标识符"`
	Name             string      `json:"name"              orm:"name"              description:"代理订阅地址名称"`
	Merchant         string      `json:"merchant"          orm:"merchant"          description:"代理订阅地址商户"`
	Description      string      `json:"description"       orm:"description"       description:"代理订阅地址描述"`
	Url              string      `json:"url"               orm:"url"               description:"代理订阅地址URL"`
	CreatedAt        *gtime.Time `json:"created_at"        orm:"created_at"        description:"代理订阅地址创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at"        orm:"updated_at"        description:"代理订阅地址更新时间"`
	SubscribeAt      *gtime.Time `json:"subscribe_at"      orm:"subscribe_at"      description:"代理订阅地址订阅时间"`
}
