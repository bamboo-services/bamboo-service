// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IMail interface {
		// GenerateMailCode 生成邮件验证码并限制频繁发送。
		//
		// 参数:
		//   - ctx: 上下文，用于控制生命周期和日志记录。
		//   - email: 邮箱地址，用于接收验证码的用户邮箱。
		//   - purpose: 验证码的用途描述。
		//
		// 返回:
		//   - *dto.MailCodeDTO: 包含生成的验证码及其相关信息。
		//   - *berror.ErrorCode: 错误信息，如发送频率限制或缓存操作失败。
		GenerateMailCode(ctx context.Context, email string, purpose string) (*dto.MailCodeDTO, *berror.ErrorCode)
		// CheckPurpose 检查邮件用途是否有效。
		//
		// 参数:
		//   - ctx: 上下文，用于控制生命周期和日志记录。
		//   - purpose: 邮件用途，需要验证的用途标识。
		//
		// 返回:
		//   - *berror.ErrorCode: 错误信息，如果用途无效或为空则返回相应错误。
		CheckPurpose(ctx context.Context, purpose string) *berror.ErrorCode
		// CheckMailTemplate 检查提供的邮件模板名称是否有效。
		//
		// 参数:
		//   - ctx: 上下文，用于控制生命周期和日志记录。
		//   - template: 模板名称，待检查的邮件模板标识。
		//
		// 返回:
		//   - *berror.ErrorCode: 错误信息，如果模板无效或为空则返回相应错误。
		CheckMailTemplate(ctx context.Context, template string) *berror.ErrorCode
		// SendMail 发送邮件至指定邮箱并填充模板数据。
		//
		// 参数:
		//   - ctx: 用于控制请求生命周期和传递上下文信息。
		//   - template: 模板名称，用于指定邮件内容格式。
		//   - mailTemplate: 包含邮件模板所需的补充数据。
		//
		// 返回:
		//   - *berror.ErrorCode: 错误信息，表示模板检查失败、邮件发送失败等可能原因。
		SendMail(ctx context.Context, template string, mailTemplate *dto.MailSendTemplateDTO) *berror.ErrorCode
	}
)

var (
	localMail IMail
)

func Mail() IMail {
	if localMail == nil {
		panic("implement not found for interface IMail, forgot register?")
	}
	return localMail
}

func RegisterMail(i IMail) {
	localMail = i
}
