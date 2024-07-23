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

// IPImportIpv4Req
//
// # 导入IPv4数据库
//
// 导入IPv4数据库，用于导入IPv4数据库操作；
type IPImportIpv4Req struct {
	g.Meta        `path:"/ip/import/ipv4" method:"Get" summary:"导入IPv4数据库" tags:"地址控制器"`
	Authorization string `json:"Authorization" v:"required#请输入授权码" in:"header"`
}

// IPImportIpv4Res
//
// # 导入IPv4数据库
//
// 返回相应的数据
type IPImportIpv4Res struct {
	g.Meta `mime:"application/json"`
}
