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
	"github.com/gogf/gf/v2/frame/g"
	"os"
)

// IPUploadIPv6Req
//
// # 上传IPv6数据库
//
// 上传IPv6数据库，用于上传IPv6数据库操作；
//
// # 参数
//   - Referer			来源地址(string)
//   - Authorization	授权码(string)
//   - File				文件(string)
type IPUploadIPv6Req struct {
	g.Meta        `path:"/ip/upload/ipv6" method:"Post" summary:"上传IPv6数据库" tags:"地址控制器" mime:"multipart/form-data"`
	Referer       string  `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Authorization string  `json:"Authorization" v:"required#请输入授权码" in:"header"`
	File          os.File `json:"file" v:"required#请输入文件" dc:"文件"`
}

// IPUploadIPv6Res
//
// # 上传IPv6数据库
//
// 返回相应的数据
type IPUploadIPv6Res struct {
	g.Meta `mime:"application/json"`
}
