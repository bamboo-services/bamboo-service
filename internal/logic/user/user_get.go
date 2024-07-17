/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package user

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

// GetUserByUUID
//
// # 获取用户信息
//
// 获取用户信息，用于获取用户信息，需要传递用户的 UUID；
// 该接口将会对用户的 UUID 进行查询，查询成功后将会返回 *entity.User 信息；
// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
// 若有其他错误信息返回错误信息；
// 优先从缓存读取数据信息，若缓存中不存在则从数据库中读取，并且读取的数据将会存入缓存中，缓存有效期 6 小时；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - userUUID		用户 UUID(uuid.UUID)
//
// # 返回
//   - user			用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (user *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.GetUserByUUID | 获取用户信息接口")
	hGetAll, err := g.Redis().HGetAll(ctx, "user:uuid:"+userUUID.String())
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取用户信息失败")
	}
	if hGetAll.IsNil() {
		// 通过 UUID 对用户数据进行获取
		err = dao.User.Ctx(ctx).Where(do.User{Uuid: userUUID}).Scan(&user)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
		}
	} else {
		err = hGetAll.Scan(&user)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "格式转化失败")
		}
	}
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	// 将数据存入至缓存，并设置过期时间为 6 小时
	err = s.GetUserIntoRedis(ctx, nil, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail
//
// # 获取用户信息
//
// 获取用户信息，用于获取用户信息，需要传递用户的 Email；
// 该接口将会对用户的 Email 进行查询，查询成功后将会返回 *entity.User 信息；
// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
// 若有其他错误信息返回错误信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - email		用户 Email(string)
//
// # 返回
//   - user			用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.GetUserByEmail | 获取用户信息接口")
	// 通过 UUID 对用户数据进行获取
	err = dao.User.Ctx(ctx).Where(do.User{Email: email}).Scan(&user)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
	}
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	return user, nil
}

// GetUserByPhone
//
// # 获取用户信息
//
// 获取用户信息，用于获取用户信息，需要传递用户的 Phone；
// 该接口将会对用户的 Phone 进行查询，查询成功后将会返回 *entity.User 信息；
// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
// 若有其他错误信息返回错误信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - phone		用户 Phone(string)
//
// # 返回
//   - user			用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) GetUserByPhone(ctx context.Context, phone string) (user *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.GetUserByPhone | 获取用户信息接口")
	// 通过 UUID 对用户数据进行获取
	err = dao.User.Ctx(ctx).Where(do.User{Phone: phone}).Scan(&user)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
	}
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	return user, nil
}

// GetUserByUsername
//
// # 获取用户信息
//
// 获取用户信息，用于获取用户信息，需要传递用户的 Username；
// 该接口将会对用户的 Username 进行查询，查询成功后将会返回 *entity.User 信息；
// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
// 若有其他错误信息返回错误信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - username		用户 Username(string)
//
// # 返回
//   - user			用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) GetUserByUsername(ctx context.Context, username string) (user *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.GetUserByUsername | 获取用户信息接口")
	// 通过 UUID 对用户数据进行获取
	err = dao.User.Ctx(ctx).Where(do.User{Username: username}).Scan(&user)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
	}
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	return user, nil
}

// GetUserIntoRedis
//
// # 获取用户信息存入缓存
//
// 获取用户信息存入缓存，用于获取用户信息，需要传递用户 uuid.UUID 或 entity.User；
// uuid.UUID 和 entity.User 务必选择其一写入，否则将产生错误；
// 执行成功后将不返回结果，若产生错误将返回错误信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - getUUID		用户 UUID(*uuid.UUID)
//   - user			用户信息(*entity.User)
//
// # 返回
//   - err			错误信息(error)
func (s *sUser) GetUserIntoRedis(ctx context.Context, getUUID *uuid.UUID, user *entity.User) (err error) {
	if user == nil {
		if getUUID == nil {
			return berror.NewError(bcode.OperationNotAllow, "用户信息 uuid 为空")
		}
		err := dao.User.Ctx(ctx).Where(do.User{Uuid: getUUID}).Scan(&user)
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户失败")
		}
		if user == nil {
			return berror.NewError(bcode.NotExist, "用户不存在")
		}
	}
	// 将数据存入至缓存，并设置过期时间为 6 小时
	err = g.Redis().HMSet(ctx, "user:uuid:"+user.Uuid, gconv.Map(user))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "存储用户信息失败")
	}
	_, err = g.Redis().Expire(ctx, "user:uuid:"+user.Uuid, 21600)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "设置用户信息过期时间失败")
	}
	return nil
}
