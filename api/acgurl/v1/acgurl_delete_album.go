/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package v1

import "github.com/gogf/gf/v2/frame/g"

// AcgurlDeleteAlbumReq
//
// # 删除图库
//
// 删除一个图库，用于删除一个图库操作；
//
// # 参数
//   - Referer			来源地址(string)
//   - Authorization	授权码(string)
//   - AlbumID			图库ID(string)
type AcgurlDeleteAlbumReq struct {
	g.Meta        `path:"/acgurl/album" method:"Delete" tags:"图库控制器" summary:"删除图库" dc:"删除一个图库"`
	Referer       string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Authorization string `json:"Authorization" v:"required#请输入授权码" in:"header"`
	AlbumID       string `json:"album_id" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$#请输入图库ID|图库ID格式不正确"` //nolint:lll
}

// AcgurlDeleteAlbumRes
//
// # 删除图库
//
// 返回指定的信息
type AcgurlDeleteAlbumRes struct {
	g.Meta `mime:"application/json"`
}
