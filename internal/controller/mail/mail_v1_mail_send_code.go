package mail

import (
	"bamboo-service/api/mail/v1"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
)

func (c *ControllerV1) MailSendCode(ctx context.Context, req *v1.MailSendCodeReq) (res *v1.MailSendCodeRes, err error) {
	iMail := service.Mail()
	mailCodeDTO, errorCode := iMail.GetMailCode(ctx, req.Email, req.Purpose)
	if errorCode != nil {
		return nil, errorCode
	}
	return &v1.MailSendCodeRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "验证码发送成功", mailCodeDTO),
	}, nil
}
