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

// InfoWebEditFilingReq
//
// # 修改网站备案信息
//
// 可以通过该接口修改网站的备案信息，包括网站的备案号、公安备案信息等信息；
// 方便直接进行热修改，而无需修改代码；
//
// # 参数
//   - Referer		来源地址(string)
//   - WebICP		网站备案号(string)
//   - WebRecord	网站公安备案信息(string
type InfoWebEditFilingReq struct {
	g.Meta    `path:"/info/web/filing" method:"Put" summary:"修改网站备案信息" tags:"信息控制器"`
	Referer   string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	WebICP    string `json:"web_icp" v:"required|length:6,100#请输入网站备案号|网站备案号长度为 6-100 位"`
	WebRecord string `json:"web_record" v:"required|length:6,100#请输入网站公安备案信息|网站备案信息长度为 6-100 位"`
}

// InfoWebEditFilingRes
//
// # 修改网站备案信息
//
// 返回相应的数据
type InfoWebEditFilingRes struct {
	g.Meta `mime:"application/json"`
}
