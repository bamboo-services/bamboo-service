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

// AvatarGetImageNoHashReq
//
// # 直接获取用户头像(非Hash)
//
// 可以通过该接口获取用户的头像信息；
// 直接获取头像，而非 JSON 数据信息；
//
// # 参数
//   - UserId		用户ID(int64)
//   - Email		邮箱地址(string)
type AvatarGetImageNoHashReq struct {
	g.Meta `path:"/avatar" method:"Get" summary:"直接获取用户头像(非Hash)" tags:"头像控制器"`
	UserId int64  `json:"user_id" v:"regex:^[0-9]+$#请输入用户ID|用户ID格式不正确"`
	Email  string `json:"email" v:"email#邮箱地址格式不正确"`
}

// AvatarGetImageNoHashRes
//
// # 获取用户头像
//
// 返回相应的数据
type AvatarGetImageNoHashRes struct {
	g.Meta `mime:"image/webp"`
}
