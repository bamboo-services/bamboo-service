package dto

import "github.com/gogf/gf/v2/os/gtime"

// ProxyTokenDTO 表示代理令牌的数据传输对象，用于封装代理令牌相关信息。
type ProxyTokenDTO struct {
	ProxyTokenUuid string      `json:"proxy_token_uuid" description:"代理令牌UUID"`
	UserUuid       string      `json:"user_uuid"        description:"用户UUID"`
	Name           string      `json:"name"             description:"代理令牌名称"`
	Description    string      `json:"description"      description:"代理令牌描述"`
	CreatedAt      *gtime.Time `json:"created_at"       description:"代理令牌创建时间"`
	ExpiredAt      *gtime.Time `json:"expired_at"        description:"代理令牌过期时间"`
}
