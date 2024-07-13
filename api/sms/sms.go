// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms

import (
	"context"

	"bamboo-service/api/sms/v2"
)

type ISmsV2 interface {
	SmsSend(ctx context.Context, req *v2.SmsSendReq) (res *v2.SmsSendRes, err error)
}
