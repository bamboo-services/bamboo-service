package mail

import (
	"bamboo-service/api/mail/v1"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

// MailSendCode 发送邮箱验证码。
//
// 参数:
//   - ctx: 上下文对象
//   - req: 包含邮箱、用途及模板信息的请求数据
//
// 返回:
//   - 包含邮件验证码及相关信息的响应数据
//   - 错误信息（如果有）
//
// 错误:
//   - 验证用途无效
//   - 验证码生成失败
//   - 验证码发送失败
func (c *ControllerV1) MailSendCode(ctx context.Context, req *v1.MailSendCodeReq) (res *v1.MailSendCodeRes, err error) {
	iMail := service.Mail()
	// 检查是否是允许的 Purpose
	errorCode := iMail.CheckPurpose(ctx, req.Purpose)
	if errorCode != nil {
		return nil, errorCode
	}

	// 生成验证码
	mailCodeDTO, errorCode := iMail.GenerateMailCode(ctx, req.Email, req.Purpose)
	if errorCode != nil {
		return nil, errorCode
	}
	// 发送验证码
	mailSendTemplate := &dto.MailSendTemplateDTO{
		Email: mailCodeDTO.Email,
		Code:  mailCodeDTO.Code,
	}
	errorCode = iMail.SendMail(ctx, req.Template, mailSendTemplate)
	if errorCode != nil {
		return nil, errorCode
	}

	return &v1.MailSendCodeRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "验证码发送成功", &dto.MailCodeDTO{
			Email:     mailCodeDTO.Email,
			CreatedAt: mailCodeDTO.CreatedAt,
		}),
	}, nil
}
