// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyToken is the golang structure for table proxy_token.
type ProxyToken struct {
	ProxyTokenUuid string      `json:"proxy_token_uuid" orm:"proxy_token_uuid" description:"代理令牌UUID"` // 代理令牌UUID
	UserUuid       string      `json:"user_uuid"        orm:"user_uuid"        description:"用户UUID"`   // 用户UUID
	Name           string      `json:"name"             orm:"name"             description:"代理令牌名称"`   // 代理令牌名称
	Description    string      `json:"description"      orm:"description"      description:"代理令牌描述"`   // 代理令牌描述
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:"代理令牌创建时间"` // 代理令牌创建时间
	ExpiredAt      *gtime.Time `json:"expired_at"       orm:"expired_at"       description:"代理令牌过期时间"` // 代理令牌过期时间
}
