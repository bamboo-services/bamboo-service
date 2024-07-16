/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package v2

import (
	"bamboo-service/internal/model/dto"
	"github.com/gogf/gf/v2/frame/g"
)

// AvatarGetNoHashReq
//
// # 获取用户头像(非Hash)
//
// 可以通过该接口获取用户的头像信息；
// 头像可以选择 UserID 或者 Email 进行获取；
// 头像若没有默认上传头像则返回 Gravatar 头像；若不存在 Gravatar 头像将会获取 QQ 号码进行获取 QQ 头像；
// 若都不存在则返回默认头像；
//
// # 参数
//   - UserID		用户ID(int64)
//   - Email		邮箱地址(string)
type AvatarGetNoHashReq struct {
	g.Meta `path:"/avatar" method:"Get" summary:"获取用户头像(非Hash)" tags:"头像控制器"`
	UserID int64  `json:"user_id" v:"regex:^[0-9]+$#请输入用户ID|用户ID格式不正确"`
	Email  string `json:"email" v:"email#邮箱地址格式不正确"`
}

// AvatarGetNoHashRes
//
// # 获取用户头像
//
// 返回相应的数据
type AvatarGetNoHashRes struct {
	g.Meta `mime:"application/json"`
	dto.AvatarAPIDTO
}
