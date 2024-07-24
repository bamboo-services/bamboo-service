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

// AcgurlEditAlbumReq
//
// # 编辑图库
//
// 编辑一个图库，用于编辑一个图库操作；
//
// # 参数
//   - Referer			来源地址(string)
//   - Authorization	授权码(string)
//   - AlbumUUID		图库ID(string)
//   - AlbumName		图库名称(string)
//   - AlbumDisplay		图库展示名称(string)
//   - AlbumDesc		图库描述(string)
//   - AlbumCover		图库封面(string)
//   - AlbumOpen		图库开放状态(bool)
//   - AlbumPattern		图库模式(uint8)
//   - MatchAddress		匹配地址([]string)
type AcgurlEditAlbumReq struct {
	g.Meta        `path:"/acgurl/album" method:"Put" tags:"图库控制器" summary:"编辑图库" dc:"编辑一个图库"`
	Referer       string   `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Authorization string   `json:"Authorization" v:"required#请输入授权码" in:"header"`
	AlbumUUID     string   `json:"album_uuid" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$#请输入图库ID|图库ID格式不正确"` //nolint:lll
	AlbumDisplay  string   `json:"display" v:"required|max-length:30#请输入图库展示名称|图库展示名称长度不能超过30个字符"`
	AlbumDesc     string   `json:"desc" v:"required|max-length:1024#请输入图库描述|图库描述长度不能超过1024个字符"`
	AlbumCover    string   `json:"cover" v:"required|url#请输入图库封面|图库封面格式不正确"`
	AlbumOpen     bool     `json:"open" v:"required|boolean#请输入图库开放状态|图库开放状态格式不正确"`
	AlbumPattern  uint8    `json:"pattern" v:"required|uint8#请输入图库模式|图库模式格式不正确"`
	MatchAddress  []string `json:"match_address" v:"required#请输入匹配地址"`
}

// AcgurlEditAlbumRes
//
// # 编辑图库
//
// 返回指定的信息
type AcgurlEditAlbumRes struct {
	g.Meta `mime:"application/json"`
}
