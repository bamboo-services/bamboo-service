package auth

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

// AuthorizationToken 生成用户的授权令牌并存储到 Redis 缓存中。
//
// 参数:
//   - ctx: 请求上下文信息。
//   - userUUID: 用户唯一标识符。
//
// 返回:
//   - *dto.AuthorizeTokenDTO: 包含生成的授权令牌及其相关信息。
//   - *berror.ErrorCode: 错误代码，表示可能的存储或其他错误情况。
func (s *sAuth) AuthorizationToken(ctx context.Context, userUUID string) (*dto.AuthorizeTokenDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "AuthorizationToken", "授权 %s 登录", userUUID)
	request := ghttp.RequestFromCtx(ctx)
	tokenUUID := uuid.New().String()
	newToken := &dredis.AuthorizeTokenRedis{
		UserUUID:  userUUID,
		Token:     tokenUUID,
		CreatedAt: gtime.Now(),
		ExpiredAt: gtime.Now().Add(2 * time.Hour),
		ClientIP:  request.GetClientIp(),
		UserAgent: request.UserAgent(),
	}
	// 检查该用户已登录信息
	keys, redisErr := g.Redis().Keys(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, "*"))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
	}
	// 检查登录设备是否超过 5 份
	if len(keys) >= 5 {
		// 删除创建最早一份
		_, redisErr = g.Redis().Del(ctx, keys[0])
		if redisErr != nil {
			return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
		}
	}
	// 缓存设置
	_, redisErr = g.Redis().HSet(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID), butil.StructToMap(newToken))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
	}
	_, redisErr = g.Redis().Expire(ctx, fmt.Sprintf(consts.RedisUserToken, userUUID, tokenUUID), int64(2*time.Hour))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
	}
	// 数据转换
	var tokenDTO *dto.AuthorizeTokenDTO
	operateErr := gconv.Struct(newToken, &tokenDTO)
	if operateErr != nil {
		return nil, berror.ErrorAddData(berror.ErrInternalServer, operateErr)
	}
	return tokenDTO, nil
}
