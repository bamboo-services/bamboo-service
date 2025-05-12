package auth

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
)

// ResetPassword 重置用户密码。
//
// 参数:
//   - ctx: 请求上下文。
//   - email: 用户邮箱，用于标识将被重置密码的用户。
//   - password: 新密码。
//
// 返回:
//   - *berror.ErrorCode: 错误码对象，若发生错误。
//
// 错误:
//   - berror.ErrDatabaseError: 数据库操作失败。
//   - berror.ErrCacheError: 缓存错误。
func (s *sAuth) ResetPassword(ctx context.Context, email, password string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "ResetPassword", "重置密码 %s", email)

	// 获取用户
	iUser := service.User()
	getUser, errorCode := iUser.GetUserByEmail(ctx, email)
	if errorCode != nil {
		return errorCode
	}

	getUser.PasswordHash = butil.PasswordEncode(password)

	errorCode = dao.User.UpdateUser(ctx, getUser)
	if errorCode != nil {
		return errorCode
	}
	g.Log().Noticef(ctx, "%s 重置密码成功 %s", "ResetPassword", email)
	return nil
}
