// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"bamboo-service/api/auth/v1"
)

type IAuthV1 interface {
	AuthInitial(ctx context.Context, req *v1.AuthInitialReq) (res *v1.AuthInitialRes, err error)
	AuthRegister(ctx context.Context, req *v1.AuthRegisterReq) (res *v1.AuthRegisterRes, err error)
}
