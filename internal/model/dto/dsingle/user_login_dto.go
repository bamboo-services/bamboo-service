package dsingle

import "bamboo-service/internal/model/dto"

// UserLoginDTO 用户登录后返回的数据传输对象。
//
// 包含用户登录后的关键信息，如用户基础信息和授权令牌信息，用于客户端和服务端的数据交互。
type UserLoginDTO struct {
	UserInfo *dto.UserInfoDTO       `json:"user"`  // 用户信息
	Token    *dto.AuthorizeTokenDTO `json:"token"` // 授权信息
}
