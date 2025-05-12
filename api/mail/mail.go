// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package mail

import (
	"context"

	"bamboo-service/api/mail/v1"
)

type IMailV1 interface {
	MailSendCode(ctx context.Context, req *v1.MailSendCodeReq) (res *v1.MailSendCodeRes, err error)
}
