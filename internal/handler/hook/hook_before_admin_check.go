package hook

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/model/dto/dredis"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"go/types"
	"regexp"
)

// BeforeAdminCheckHook 在管理员权限检查前触发的钩子，用于预处理请求或执行额外验证逻辑。
func BeforeAdminCheckHook(r *ghttp.Request) {
	// 获取用户的请求头
	authorizationToken := r.GetHeader("Authorization")
	getUserUUID := r.GetHeader("X-User-UUID")
	// 检查是否符合格式
	if authorizationToken == "" || getUserUUID == "" {
		errorCode := berror.ErrInvalidToken
		r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
		})
		r.Response.WriteStatus(int(errorCode.Code / 100))
		return
	}

	// 验证 Bearer token 格式
	bearerTokenPattern := `^Bearer [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	bearerUserPattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`
	matchedToken, err := regexp.MatchString(bearerTokenPattern, authorizationToken)
	matchedUser, err := regexp.MatchString(bearerUserPattern, getUserUUID)
	if err != nil || !matchedToken || !matchedUser {
		errorCode := berror.ErrInvalidToken
		r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
		})
		r.Response.WriteStatus(int(errorCode.Code / 100))
		return
	}

	// 验证 token
	token := butil.TokenRemoveBearer(authorizationToken)
	getToken, redisErr := g.Redis().HGetAll(r.GetCtx(), fmt.Sprintf(consts.RedisUserToken, getUserUUID, token))
	if redisErr != nil {
		errorCode := berror.ErrCacheError
		r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
		})
		r.Response.WriteStatus(int(errorCode.Code / 100))
		return
	}

	if !getToken.IsNil() || !getToken.IsEmpty() {
		// 用户存在，检查是否过期
		var tokenEntity *dredis.AuthorizeTokenRedis
		operateErr := getToken.Scan(&tokenEntity)
		if operateErr != nil {
			errorCode := berror.ErrInternalServer
			r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
				Code:     errorCode.Code,
				Message:  errorCode.Message,
				Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
				Time:     gtime.Now().TimestampMilli(),
			})
			r.Response.WriteStatus(int(errorCode.Code / 100))
			return
		}

		// 检查 Token 是否已经过期
		if gtime.Now().After(tokenEntity.ExpiredAt) {
			errorCode := berror.ErrTokenExpired
			r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
				Code:     errorCode.Code,
				Message:  errorCode.Message,
				Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
				Time:     gtime.Now().TimestampMilli(),
			})
			r.Response.WriteStatus(int(errorCode.Code / 100))
			// 删除该 Redis Token
			_, redisErr := g.Redis().Del(r.GetCtx(), fmt.Sprintf(consts.RedisUserToken, getUserUUID, token))
			if redisErr != nil {
				errorCode := berror.ErrorAddData(berror.ErrCacheError, redisErr)
				r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
					Code:     errorCode.Code,
					Message:  errorCode.Message,
					Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
					Time:     gtime.Now().TimestampMilli(),
				})
				r.Response.WriteStatus(int(errorCode.Code / 100))
			}
			return
		}

		// 验证登录 IP 和 UserAgent 是否一致
		if tokenEntity.ClientIP != r.GetClientIp() || tokenEntity.UserAgent != r.UserAgent() {
			errorCode := berror.ErrInvalidToken
			r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
				Code:     errorCode.Code,
				Message:  errorCode.Message,
				Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
				Time:     gtime.Now().TimestampMilli(),
			})
			r.Response.WriteStatus(int(errorCode.Code / 100))
			// 删除该 Redis Token
			_, redisErr := g.Redis().Del(r.GetCtx(), fmt.Sprintf(consts.RedisUserToken, getUserUUID, token))
			if redisErr != nil {
				errorCode := berror.ErrorAddData(berror.ErrCacheError, redisErr)
				r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
					Code:     errorCode.Code,
					Message:  errorCode.Message,
					Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
					Time:     gtime.Now().TimestampMilli(),
				})
				r.Response.WriteStatus(int(errorCode.Code / 100))
			}
			return
		}

	} else {
		errorCode := berror.ErrInvalidToken
		r.Response.WriteJson(&bmodels.ResponseDTO[types.Nil]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
		})
		r.Response.WriteStatus(int(errorCode.Code / 100))
		return
	}
}
