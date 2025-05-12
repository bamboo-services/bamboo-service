package user

import (
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
)

// GetUserByUsername 根据用户名检查用户是否存在并返回用户信息。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - username: 需要检查的用户名。
//
// 返回:
//   - *entity.User: 用户信息，若用户存在则返回详细信息，否则为nil。
//   - *berror.ErrorCode: 错误代码，表示用户不存在或其他错误情况。
func (s *sUser) GetUserByUsername(ctx context.Context, username string) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetUserByUsername", "根据用户名 %s 获取用户信息", username)
	getUser, errorCode := dao.User.GetUserByUsername(ctx, username)
	if errorCode != nil {
		return nil, errorCode
	}
	if getUser == nil {
		return nil, custom.ErrorUserNotExist
	}
	return getUser, nil
}

// GetUserByEmail 根据邮箱检查用户是否存在并返回用户信息。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - email: 需要检查的邮箱地址。
//
// 返回:
//   - *entity.User: 用户信息，若用户存在则返回详细信息，否则为nil。
//   - *berror.ErrorCode: 错误代码，表示用户不存在或其他错误情况。
func (s *sUser) GetUserByEmail(ctx context.Context, email string) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetUserByEmail", "检查邮箱 %s 是否存在", email)
	getUser, errorCode := dao.User.GetUserByEmail(ctx, email)
	if errorCode != nil {
		return nil, errorCode
	}
	if getUser != nil {
		return nil, custom.ErrorUserNotExist
	}
	return getUser, nil
}

// GetUserByPhone 根据手机号检查用户是否存在并返回用户信息。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - phone: 需要检查的手机号。
//
// 返回:
//   - *entity.User: 用户信息，若用户存在则返回详细信息，否则为 nil。
//   - *berror.ErrorCode: 错误代码，表示用户不存在或其他错误情况。
func (s *sUser) GetUserByPhone(ctx context.Context, phone string) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetUserByPhone", "根据手机号 %s 获取用户信息", phone)
	getUser, errorCode := dao.User.GetUserByPhone(ctx, phone)
	if errorCode != nil {
		return nil, errorCode
	}
	if getUser != nil {
		return nil, custom.ErrorUserNotExist
	}
	return getUser, nil
}
