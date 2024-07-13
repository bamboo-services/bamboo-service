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

import (
	"bamboo-service/internal/model/dto"
	"github.com/gogf/gf/v2/frame/g"
)

// InfoWebShowReq
//
// # 获取网站信息
//
// 可以通过该接口获取网站的信息，包括网站的名称、描述、关键字、备案号等信息；
// 方便直接进行热修改，而无需修改代码；
//
// # 参数
//   - Referer		来源地址(string)
type InfoWebShowReq struct {
	g.Meta  `path:"/info/web" method:"Get" summary:"获取网站信息" tags:"信息控制器"`
	Referer string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
}

// InfoWebShowRes
//
// # 获取网站信息
//
// 返回相应的数据
type InfoWebShowRes struct {
	g.Meta `mime:"application/json"`
	dto.InfoWebDTO
}
