package dredis

import "github.com/gogf/gf/v2/os/gtime"

type AuthorizeTokenRedis struct {
	UserUUID  string      `json:"user_uuid"`  // 用户唯一标识符
	Token     string      `json:"token"`      // 授权token
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	ExpiredAt *gtime.Time `json:"expired_at"` // 过期时间
	ClientIP  string      `json:"client_ip"`  // 客户端IP
	UserAgent string      `json:"user_agent"`
}
