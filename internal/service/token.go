// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IToken interface {
		// GenerateNewAuthorizationToken 生成新的授权令牌并缓存。
		//
		// 参数:
		//   - ctx: 上下文信息，用于请求追踪和控制。
		//   - userUUID: 用户唯一标识符。
		//
		// 返回:
		//   - *dto.AuthorizeTokenDTO: 包含授权令牌信息的 DTO。
		//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
		GenerateNewAuthorizationToken(ctx context.Context, userUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode)
		// RemoveToken 删除指定用户的授权令牌。
		//
		// 参数:
		//   - ctx: 上下文信息，用于控制生命周期和请求追踪。
		//   - userUUID: 用户唯一标识符。
		//   - tokenUUID: 授权令牌唯一标识符。
		//
		// 返回:
		//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
		//
		// 错误:
		//   - berror.ErrCacheError: Redis 操作错误。
		//   - berror.ErrInvalidToken: 提供的令牌无效或不存在。
		RemoveToken(ctx context.Context, userUUID string, tokenUUID string) *berror.ErrorCode
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
