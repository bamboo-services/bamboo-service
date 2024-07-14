// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package avatar

import (
	"context"

	"bamboo-service/api/avatar/v2"
	"bamboo-service/api/avatar/v3"
)

type IAvatarV2 interface {
	AvatarGet(ctx context.Context, req *v2.AvatarGetReq) (res *v2.AvatarGetRes, err error)
	AvatarGetNoHash(ctx context.Context, req *v2.AvatarGetNoHashReq) (res *v2.AvatarGetNoHashRes, err error)
}

type IAvatarV3 interface {
	AvatarGetImage(ctx context.Context, req *v3.AvatarGetImageReq) (res *v3.AvatarGetImageRes, err error)
	AvatarGetImageNoHash(ctx context.Context, req *v3.AvatarGetImageNoHashReq) (res *v3.AvatarGetImageNoHashRes, err error)
}
