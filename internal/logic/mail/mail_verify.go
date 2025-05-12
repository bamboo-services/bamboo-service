package mail

import (
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
)

// VerifyMailCode 验证邮件验证码是否有效。
//
// 参数:
//   - ctx: 上下文，用于控制生命周期和日志记录。
//   - purpose: 验证码的用途描述。
//   - email: 邮箱地址，待验证的用户邮箱。
//   - code: 用户输入的验证码。
//
// 返回:
//   - *berror.ErrorCode: 错误信息，若验证失败则返回对应错误。
func (s *sMail) VerifyMailCode(ctx context.Context, purpose string, email string, code string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "VerifyMailCode", "验证 %s 的验证码", email)

	// 获取验证码
	errorCode := s.CheckPurpose(ctx, purpose)
	if errorCode != nil {
		return errorCode
	}
	mailCodeEntity, errorCode := dao.EmailCode.GetLastMailCode(ctx, purpose, email)
	if errorCode != nil {
		return errorCode
	}
	if mailCodeEntity == nil {
		return custom.ErrorMailCodeNotExist
	}

	// 检查验证码信息
	if gtime.Now().After(mailCodeEntity.ExpiredAt) {
		// 删除旧验证码(过期)
		errorCode := dao.EmailCode.RemoveMailCode(ctx, mailCodeEntity.CodeUuid)
		if errorCode != nil {
			return errorCode
		}
		return custom.ErrMailCodeHasExpired
	}

	// 检查用途是否正确
	if mailCodeEntity.Purpose != purpose {
		glog.Noticef(ctx, "%s 验证码用途不正确，请检查 %s 的验证码用途", "VerifyMailCode", email)
		return custom.ErrorMailCodeIncorrect
	}

	// 检查验证码是否正确
	if mailCodeEntity.Code == code {
		// 删除旧验证码(已验证)
		errorCode := dao.EmailCode.RemoveMailCode(ctx, mailCodeEntity.CodeUuid)
		if errorCode != nil {
			return errorCode
		}
		return nil
	} else {
		glog.Noticef(ctx, "%s 验证码不正确，请检查 %s 的验证码", "VerifyMailCode", email)
		return custom.ErrorMailCodeIncorrect
	}
}
