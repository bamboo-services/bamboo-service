package auth

import (
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao/command/duser"
	"bamboo-service/internal/model/dto"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/util/gconv"
)

// CheckUserExistByUsername 检查指定用户名的用户是否存在。
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - Username: 需要检查的用户名。
//
// 返回:
//   - err: 错误代码，表示用户不存在或其他错误情况。
func (s *sAuth) CheckUserExistByUsername(ctx context.Context, Username string) (err *berror.ErrorCode) {
	_, errorCode := duser.GetUserByUsername(ctx, Username)
	if errorCode != nil {
		return errorCode
	}
	return nil
}

// UserLogin 验证用户登录，并返回用户信息。
//
// 参数:
//   - ctx: 请求上下文。
//   - Username: 用户名。
//   - Password: 用户密码。
//
// 返回:
//   - userInfo: 用户信息数据传输对象。
//   - err: 错误代码，表示登录失败的原因。
func (s *sAuth) UserLogin(ctx context.Context, Username, Password string) (*dto.UserInfoDTO, *berror.ErrorCode) {
	getUser, errorCode := duser.GetUserByUsername(ctx, Username)
	if errorCode != nil {
		return nil, errorCode
	}
	// 检查用户密码
	if butil.PasswordVerify(Password, (getUser).PasswordHash) {
		userInfo := &dto.UserInfoDTO{}
		operateErr := gconv.Struct(getUser, userInfo)
		if operateErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, operateErr)
		}
		return userInfo, nil
	} else {
		return nil, custom.ErrorUserPasswordIncorrect
	}
}
