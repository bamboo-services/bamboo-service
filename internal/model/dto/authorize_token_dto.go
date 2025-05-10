package dto

import "github.com/gogf/gf/v2/os/gtime"

// AuthorizeTokenDTO 表示授权令牌的数据传输对象。
type AuthorizeTokenDTO struct {
	UserUUID  string      `json:"user_uuid"`  // 用户唯一标识符
	Token     string      `json:"token"`      // 授权token
	CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	ExpiredAt *gtime.Time `json:"expired_at"` // 过期时间
}
