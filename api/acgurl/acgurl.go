/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package acgurl

import (
	"context"

	"bamboo-service/api/acgurl/v1"
)

type IAcgurlV1 interface {
	AcgurlAddPhotoLink(ctx context.Context, req *v1.AcgurlAddPhotoLinkReq) (res *v1.AcgurlAddPhotoLinkRes, err error)
	AcgurlCreateAlbum(ctx context.Context, req *v1.AcgurlCreateAlbumReq) (res *v1.AcgurlCreateAlbumRes, err error)
	AcgurlDeleteAlbum(ctx context.Context, req *v1.AcgurlDeleteAlbumReq) (res *v1.AcgurlDeleteAlbumRes, err error)
	AcgurlEditAlbum(ctx context.Context, req *v1.AcgurlEditAlbumReq) (res *v1.AcgurlEditAlbumRes, err error)
	AcgurlGetAlbum(ctx context.Context, req *v1.AcgurlGetAlbumReq) (res *v1.AcgurlGetAlbumRes, err error)
}
