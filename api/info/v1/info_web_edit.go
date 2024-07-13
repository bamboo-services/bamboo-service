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

// InfoWebEditReq
//
// # 修改网站信息
//
// 可以通过该接口修改网站的信息，包括网站的名称、描述、关键字等信息；
// 方便直接进行热修改，而无需修改代码；
//
// # 参数
//   - Referer		来源地址(string)
//   - WebName		网站名称(string)
//   - WebDesc		网站描述(string)
//   - WebKey		网站关键字(string)
//   - WebCopy		网站版权(string)
type InfoWebEditReq struct {
	g.Meta  `path:"/info/web" method:"Put" summary:"修改网站信息" tags:"信息控制器"`
	Referer string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	WebName string `json:"web_name" v:"required|length:6,30#请输入网站名称|网站名称长度为 6-30 位"`
	WebDesc string `json:"web_desc" v:"required|length:6,100#请输入网站描述|网站描述长度为 6-100 位"`
	WebKey  string `json:"web_key" v:"required|length:6,100#请输入网站关键字|网站关键字长度为 6-100 位"`
	WebCopy string `json:"web_copy" v:"required|length:6,100#请输入网站版权|网站版权长度为 6-100 位"`
	WebLogo string `json:"web_logo" v:"required|url#请输入网站Logo|网站Logo格式不正确"`
	WebFav  string `json:"web_fav" v:"required|url#请输入网站Favicon|网站Favicon格式不正确"`
}

// InfoWebEditRes
//
// # 修改网站信息
//
// 返回相应的数据
type InfoWebEditRes struct {
	g.Meta `mime:"application/json"`
}
