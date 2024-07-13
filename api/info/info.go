// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package info

import (
	"context"

	"bamboo-service/api/info/v1"
)

type IInfoV1 interface {
	InfoWebEdit(ctx context.Context, req *v1.InfoWebEditReq) (res *v1.InfoWebEditRes, err error)
	InfoWebEditFiling(ctx context.Context, req *v1.InfoWebEditFilingReq) (res *v1.InfoWebEditFilingRes, err error)
	InfoWebShow(ctx context.Context, req *v1.InfoWebShowReq) (res *v1.InfoWebShowRes, err error)
}
