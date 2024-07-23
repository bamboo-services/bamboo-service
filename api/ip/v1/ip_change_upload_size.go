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

// IPChangeUploadFileSizeReq
//
// # 修改上传文件大小
//
// 修改上传文件大小，用于修改上传文件大小操作；
//
// # 参数
//   - Referer			来源地址(string)
//   - Authorization	授权码(string)
//   - Length			文件大小(int64)
type IPChangeUploadFileSizeReq struct {
	g.Meta        `path:"/ip/upload/size" method:"Put" summary:"修改上传文件大小" tags:"地址控制器"`
	Referer       string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Authorization string `json:"Authorization" v:"required#请输入授权码" in:"header"`
	Length        int64  `json:"Length" v:"required#请输入文件大小" default:"8388608"`
}

// IPChangeUploadFileSizeRes
//
// # 修改上传文件大小
//
// 返回相应的数据
type IPChangeUploadFileSizeRes struct {
	g.Meta `mime:"application/json"`
}
