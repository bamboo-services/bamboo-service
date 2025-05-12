package mail

import (
	"bamboo-service/api/mail/v1"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

func (c *ControllerV1) MailSendCode(ctx context.Context, req *v1.MailSendCodeReq) (res *v1.MailSendCodeRes, err error) {
	iMail := service.Mail()
	// 检查是否是允许的 Purpose
	// TODO-[25051201] 检查是否是允许的 Purpose

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
