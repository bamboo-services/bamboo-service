package duser

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// GetUserByUUID 根据用户 UUID 获取用户信息。
//
// 查询如果出现错误会返回报错，如果查询成功会返回用户信息（不存在 *entity.User 返回 nil）
//
// 参数:
//   - ctx: 上下文对象，用于控制请求的生命周期。
//   - UserUUID: 用户唯一标识符。
//
// 返回:
//   - *entity.User: 用户信息，如果存在。
//   - *berror.ErrorCode: 错误码对象，如果出错。
//
// 错误:
//   - berror.ErrCacheError: 缓存获取失败或保存失败。
//   - berror.ErrDatabaseError: 数据库查询错误。
//   - berror.ErrInternalServer: 数据构造或解析错误。
func GetUserByUUID(ctx context.Context, userUUID string) (*entity.User, *berror.ErrorCode) {
	blog.DaoInfo(ctx, "GetUserByUUID", "通过 UUID 获取用户")
	redisRecord, redisErr := g.Redis().HGetAll(ctx, fmt.Sprintf(consts.RedisUserUUID, userUUID))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
	}
	var user *entity.User
	if redisRecord.IsNil() || redisRecord.IsEmpty() {
		// 数据库获取
		sqlErr := dao.User.Ctx(ctx).Where(do.User{UserUuid: userUUID}).Scan(&user)
		if sqlErr != nil {
			return nil, berror.ErrorAddData(berror.ErrDatabaseError, sqlErr.Error())
		}
		if user != nil {
			_, redisErr := g.Redis().HSet(ctx, fmt.Sprintf(consts.RedisUserUUID, userUUID), butil.StructToMap(user))
			if redisErr != nil {
				return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
			}
			_, redisErr = g.Redis().Expire(ctx, fmt.Sprintf(consts.RedisUserUUID, userUUID), int64(time.Hour))
			if redisErr != nil {
				return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
			}
		}
		return user, nil
	} else {
		user, operateErr := butil.MapToStruct(redisRecord.Map(), user)
		if operateErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, operateErr.Error())
		}
		return user, nil
	}
}

// GetUserByUsername 通过用户名获取用户信息。
//
// 查询如果出现错误会返回报错，如果查询成功会返回用户信息（不存在 *entity.User 返回 nil）
//
// 参数:
//   - ctx: 上下文对象，用于控制请求生命周期。
//   - username: 用户名，用于查询用户数据。
//
// 返回:
//   - *entity.User: 用户结构体，包含用户的详细信息。
//   - *berror.ErrorCode: 错误码对象，表示操作失败的原因。
//
// 错误:
//   - berror.ErrCacheError: 缓存获取失败或保存失败。
//   - berror.ErrDatabaseError: 数据库查询错误。
//   - berror.ErrInternalServer: 数据构造或解析错误。
func GetUserByUsername(ctx context.Context, username string) (*entity.User, *berror.ErrorCode) {
	blog.DaoInfo(ctx, "GetUserByUsername", "通过 Username 获取用户")
	redisRecord, redisErr := g.Redis().GetEX(ctx, fmt.Sprintf(consts.RedisUserUsername, username))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
	}
	var user *entity.User
	if redisRecord.IsNil() || redisRecord.IsEmpty() {
		redisErr := dao.User.Ctx(ctx).Where(do.User{Username: username}).Scan(&user)
		if redisErr != nil {
			return nil, berror.ErrorAddData(berror.ErrDatabaseError, redisErr)
		}
		redisErr = g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisUserUsername, username), user.UserUuid, int64(time.Hour))
		if redisErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, redisErr)
		}
		_, redisErr = g.Redis().HSet(ctx, fmt.Sprintf(consts.RedisUserUUID, user.UserUuid), butil.StructToMap(user))
		if redisErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, redisErr)
		}
		return user, nil
	} else {
		getUser, errorCode := GetUserByUUID(ctx, redisRecord.String())
		if errorCode != nil {
			return nil, errorCode
		}
		return getUser, nil
	}
}
