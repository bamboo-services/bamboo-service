package dto

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyBaseSubscriptionDTO 表示代理订阅地址的基础信息。
//
// 包括唯一标识符、代理组、用户、订阅名称、商户、描述、URL，以及创建、更新和订阅时间等字段。
type ProxyBaseSubscriptionDTO struct {
	SubscriptionUuid string      `json:"subscription_uuid" description:"代理订阅地址唯一标识符"`
	ProxyGroupUuid   string      `json:"proxy_group_uuid"  description:"代理组唯一标识符"`
	UserUuid         string      `json:"user_uuid"         description:"用户唯一标识符"`
	Name             string      `json:"name"              description:"代理订阅地址名称"`
	Merchant         string      `json:"merchant"          description:"代理订阅地址商户"`
	Description      string      `json:"description"       description:"代理订阅地址描述"`
	Url              string      `json:"url"               description:"代理订阅地址URL"`
	CreatedAt        *gtime.Time `json:"created_at"        description:"代理订阅地址创建时间"`
	UpdatedAt        *gtime.Time `json:"updated_at"        description:"代理订阅地址更新时间"`
	SubscribeAt      *gtime.Time `json:"subscribe_at"      description:"代理订阅地址订阅时间"`
}

// ProxyBaseSubscriptionLiteDTO 表示代理订阅地址的简要信息传输对象。
//
// 提供与代理订阅地址相关的基本属性，如唯一标识符、名称和商户等信息。
type ProxyBaseSubscriptionLiteDTO struct {
	SubscriptionUuid string `json:"subscription_uuid" description:"代理订阅地址唯一标识符"`
	ProxyGroupUuid   string `json:"proxy_group_uuid"  description:"代理组唯一标识符"`
	UserUuid         string `json:"user_uuid"         description:"用户唯一标识符"`
	Name             string `json:"name"              description:"代理订阅地址名称"`
	Merchant         string `json:"merchant"          description:"代理订阅地址商户"`
	Description      string `json:"description"       description:"代理订阅地址描述"`
}
