// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms

import (
	"context"

	"bamboo-service/api/sms/v1"
)

type ISmsV1 interface {
	SmsSend(ctx context.Context, req *v1.SmsSendReq) (res *v1.SmsSendRes, err error)
}
