package mail

import (
	"bamboo-service/internal/consts"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
)

// CheckPurpose 检查邮件用途是否有效。
//
// 参数:
//   - ctx: 上下文，用于控制生命周期和日志记录。
//   - purpose: 邮件用途，需要验证的用途标识。
//
// 返回:
//   - *berror.ErrorCode: 错误信息，如果用途无效或为空则返回相应错误。
func (s *sMail) CheckPurpose(ctx context.Context, purpose string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckPurpose", "检查邮件用途 %s", purpose)
	if purpose == "" {
		return berror.ErrorAddData(berror.ErrInvalidParameters, "用途不能为空")
	}
	for _, v := range consts.MailPurposeList {
		if v.Name == purpose {
			return nil
		}
	}
	return berror.ErrorAddData(berror.ErrInvalidParameters, "用途不存在")
}
