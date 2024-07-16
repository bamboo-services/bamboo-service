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
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"regexp"
)

// UserExistByUsername
//
// # 用户是否存在
//
// 用户是否存在，用于检查用户是否存在；检查用户名是否重复，如果重复则返回错误；
// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - username		用户名(string)
//
// # 返回
//   - getUser		用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) UserExistByUsername(ctx context.Context, username string) (getUser *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.UserExistByUsername | 用户(用户名)是否存在接口")
	// 对 username 进行正则表达式判断确认数据没有问题
	matched, err := regexp.Match("^[0-9A-Za-z-_]{6,30}$", []byte(username))
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.UnknownError, err)
	}
	if matched {
		getUser, err = s.GetUserByUsername(ctx, username)
		if err != nil {
			return getUser, err
		}
		return nil, nil
	} else {
		return nil, berror.NewError(bcode.VerifyFailed, "用户名格式不正确")
	}
}

// UserExistByPhone
//
// # 用户是否存在
//
// 用户是否存在，用于检查用户是否存在；检查手机号是否重复，如果重复则返回错误；
// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - phone		手机号(string)
//
// # 返回
//   - getUser		用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) UserExistByPhone(ctx context.Context, phone string) (getUser *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.UserExistByPhone | 用户(手机号)是否存在接口")
	// 对 phone 进行正则表达式判断确认数据没有问题
	matched, err := regexp.Match(
		"^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$",
		[]byte(phone),
	)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.UnknownError, err)
	}
	if matched {
		getUser, err = s.GetUserByPhone(ctx, phone)
		if err != nil {
			return getUser, err
		}
		return nil, nil
	} else {
		return nil, berror.NewError(bcode.VerifyFailed, "用户名格式不正确")
	}
}

// UserExistByEmail
//
// # 用户是否存在
//
// 用户是否存在，用于检查用户是否存在；检查邮箱是否重复，如果重复则返回错误；
// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - email		邮箱(string)
//
// # 返回
//   - getUser		用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) UserExistByEmail(ctx context.Context, email string) (getUser *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.UserExistByEmail | 用户(邮箱)是否存在接口")
	// 对 email 进行正则表达式判断确认数据没有问题
	matched, err := regexp.Match("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", []byte(email))
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.UnknownError, err)
	}
	if matched {
		getUser, err = s.GetUserByEmail(ctx, email)
		if err != nil {
			return getUser, err
		}
		return nil, nil
	} else {
		return nil, berror.NewError(bcode.VerifyFailed, "用户名格式不正确")
	}
}

// UserExistByUUID
//
// # 用户是否存在
//
// 用户是否存在，用于检查用户是否存在；检查 UUID 是否存在，如果存在则返回错误；
// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - getUUID		用户 UUID(string)
//
// # 返回
//   - getUser		用户信息(*entity.User)
//   - err			错误信息(error)
func (s *sUser) UserExistByUUID(ctx context.Context, getUUID uuid.UUID) (getUser *entity.User, err error) {
	g.Log().Notice(ctx, "[SERV] user.UserExistByUUID | 用户(UUID)是否存在接口")
	getUser, err = s.GetUserByUUID(ctx, getUUID)
	if err != nil {
		return getUser, err
	}
	return nil, nil
}
