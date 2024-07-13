// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"bamboo-service/api/auth/v2"
)

type IAuthV2 interface {
	AuthRegister(ctx context.Context, req *v2.AuthRegisterReq) (res *v2.AuthRegisterRes, err error)
}
