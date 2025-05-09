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
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
func GetUserByUUID(ctx context.Context, UserUUID string) (*entity.User, *berror.ErrorCode) {
	blog.DaoInfo(ctx, "GetUserByUUID", "通过 UUID 获取用户")
	redisRecord, redisErr := g.Redis().HGetAll(ctx, fmt.Sprintf(consts.RedisUserUUID, UserUUID))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
	}
	var user *entity.User
	if redisRecord.IsNil() || redisRecord.IsEmpty() {
		// 数据库获取
		sqlErr := dao.User.Ctx(ctx).Where(do.User{UserUuid: UserUUID}).Scan(&user)
		if sqlErr != nil {
			return nil, berror.ErrorAddData(berror.ErrDatabaseError, sqlErr.Error())
		}
		if user != nil {
			_, redisErr := g.Redis().HSet(ctx, fmt.Sprintf(consts.RedisUserUUID, UserUUID), gconv.Map(user))
			if redisErr != nil {
				return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
			}
			_, redisErr = g.Redis().Expire(ctx, fmt.Sprintf(consts.RedisUserUUID, UserUUID), int64(time.Hour))
			if redisErr != nil {
				return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr.Error())
			}
		}
	} else {
		// 缓存获取
		operateErr := redisRecord.Struct(&user)
		if operateErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, operateErr.Error())
		}
	}
	return user, nil
}
