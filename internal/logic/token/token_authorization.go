package token

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/dto/dredis"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"time"
)

// GenerateNewAuthorizationToken 生成新的授权令牌并缓存。
//
// 参数:
//   - ctx: 上下文信息，用于请求追踪和控制。
//   - userUUID: 用户唯一标识符。
//
// 返回:
//   - *dto.AuthorizeTokenDTO: 包含授权令牌信息的 DTO。
//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
func (s *sToken) GenerateNewAuthorizationToken(ctx context.Context, userUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "AuthorizationToken", "授权 %s 登录", userUUID)
	request := ghttp.RequestFromCtx(ctx)
	tokenUUID := uuid.New().String()
	newToken := &dredis.AuthorizeTokenRedis{
		UserUUID:  userUUID,
		Token:     tokenUUID,
		CreatedAt: gtime.Now(),
		ExpiredAt: gtime.Now().Add(6 * time.Hour),
		ClientIP:  request.GetClientIp(),
		UserAgent: request.UserAgent(),
	}
	// 检查该用户已登录信息
	keys, redisErr := g.Redis().Keys(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, "*"))
	if redisErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrCacheError, redisErr)
	}
	// 检查登录设备是否超过 5 份
	if len(keys) >= 5 {
		// 删除创建最早一份
		_, redisErr = g.Redis().Del(ctx, keys[0])
		if redisErr != nil {
			return nil, berror.ErrorAddData(&berror.ErrCacheError, redisErr)
		}
	}
	// 缓存设置
	_, redisErr = g.Redis().HSet(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID), butil.StructToMap(newToken))
	if redisErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrCacheError, redisErr)
	}
	_, redisErr = g.Redis().Expire(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID), int64(2*time.Hour))
	if redisErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrCacheError, redisErr)
	}
	// 数据转换
	var tokenDTO *dto.AuthorizeTokenDTO
	operateErr := gconv.Struct(newToken, &tokenDTO)
	if operateErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr)
	}
	return tokenDTO, nil
}

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
func (s *sToken) RemoveToken(ctx context.Context, userUUID, tokenUUID string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "RemoveToken", "删除授权令牌 %s", tokenUUID)
	_, errorCode := s.GetToken(ctx, userUUID, tokenUUID)
	if errorCode != nil {
		return errorCode
	}
	_, redisErr := g.Redis().Del(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID))
	if redisErr != nil {
		return berror.ErrorAddData(&berror.ErrCacheError, redisErr)
	}
	return nil
}

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
func (s *sToken) GetToken(ctx context.Context, userUUID, tokenUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetToken", "获取 %s 授权令牌 %s", userUUID, tokenUUID)
	getToken, redisErr := g.Redis().GetEX(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID))
	if redisErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrCacheError, redisErr)
	}
	if getToken.IsNil() || getToken.IsEmpty() {
		return nil, berror.ErrorAddData(&berror.ErrInvalidToken, "无效的令牌")
	}
	var tokenDTO *dto.AuthorizeTokenDTO
	operateErr := gconv.Struct(getToken.Map(), &tokenDTO)
	if operateErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr)
	}
	return tokenDTO, nil
}
