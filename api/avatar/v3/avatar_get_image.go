/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package v3

import "github.com/gogf/gf/v2/frame/g"

// AvatarGetImageReq
//
// # 直接获取用户头像
//
// 可以通过该接口获取用户的头像信息；
// 直接获取头像，而非 JSON 数据信息；
//
// # 参数
//   - Hash			头像Hash(string)
type AvatarGetImageReq struct {
	g.Meta `path:"/avatar/{hash}" method:"Get" summary:"直接获取用户头像" tags:"头像控制器"`
	Hash   string `json:"hash" v:"required|length:32,32#请输入头像Hash|头像Hash长度为 32 位" in:"path"`
}

// AvatarGetImageRes
//
// # 获取用户头像
//
// 返回相应的数据
type AvatarGetImageRes struct {
	g.Meta `mime:"image/webp"`
}
