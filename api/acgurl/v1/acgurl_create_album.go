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

// AcgurlCreateAlbumReq
//
// # 创建图库
//
// 创建一个图库，用于创建一个图库操作；
//
// # 参数
//   - Referer			来源地址(string)
//   - Authorization	授权码(string)
//   - AlbumName		图库名称(string)
//   - AlbumDesc		图库描述(string)
//   - AlbumCover		图库封面(string)
type AcgurlCreateAlbumReq struct {
	g.Meta        `path:"/acgurl/album" method:"Post" tags:"图库控制器" summary:"创建图库" dc:"创建一个图库"`
	Referer       string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Authorization string `json:"Authorization" v:"required#请输入授权码" in:"header"`
	AlbumName     string `json:"name" v:"required|max-length:30#请输入图库名称|图库名称长度不能超过30个字符"`
	AlbumDisplay  string `json:"display" v:"required|max-length:30#请输入图库展示名称|图库展示名称长度不能超过30个字符"`
	AlbumDesc     string `json:"desc" v:"required|max-length:1024#请输入图库描述|图库描述长度不能超过1024个字符"`
	AlbumCover    string `json:"cover" v:"required|url#请输入图库封面|图库封面格式不正确"`
	AlbumOpen     bool   `json:"open" v:"required|boolean#请输入图库开放状态|图库开放状态格式不正确"`
}

// AcgurlCreateAlbumRes
//
// # 创建图库
//
// 返回指定的信息
type AcgurlCreateAlbumRes struct {
	g.Meta `mime:"application/json"`
}
