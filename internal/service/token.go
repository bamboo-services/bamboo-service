// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/os/gtime"
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
		// GetToken 获取指定用户的授权令牌。
		//
		// 参数:
		//   - ctx: 上下文信息，用于请求追踪和控制。
		//   - userUUID: 用户唯一标识符。
		//   - tokenUUID: 授权令牌唯一标识符。
		//
		// 返回:
		//   - *dto.AuthorizeTokenDTO: 包含授权令牌信息的 DTO。
		//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
		//
		// 错误:
		//   - berror.ErrCacheError: Redis 操作错误。
		//   - berror.ErrInvalidToken: 提供的令牌无效或不存在。
		//   - berror.ErrInternalServer: 数据转换错误。
		GetToken(ctx context.Context, userUUID string, tokenUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode)
		// GenerateProxyToken 生成新的代理令牌。
		//
		// 参数:
		//   - ctx: 上下文信息，用于请求追踪和控制。
		//   - userUUID: 用户的UUID。
		//   - name: 代理令牌的名称。
		//   - desc: 代理令牌的描述信息。
		//   - expiredAt: 代理令牌的过期时间，不能超过当前时间的一年后。
		//
		// 返回:
		//   - *entity.ProxyToken: 包含生成的代理令牌信息。
		//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
		//
		// 错误:
		//   - berror.ErrInvalidParameters: 参数无效，例如有效期超过一年。
		GenerateProxyToken(ctx context.Context, userUUID string, name string, desc string, expiredAt *gtime.Time) (*dto.ProxyTokenDTO, *berror.ErrorCode)
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
