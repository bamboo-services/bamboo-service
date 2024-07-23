/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IIP interface {
		// IPv4FileUpload
		//
		// # 上传IPv4数据库
		//
		// 上传IPv4数据库，用于上传IPv4数据库操作；
		// 该接口将会对上传的文件进行解码，解码成功后将会将文件写入到 upload/ip_location/database_location_ipv4.scv 文件中；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - file			文件(string)
		//
		// # 返回
		//   - err			错误信息(error)
		IPv4FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error)
		// IPv6FileUpload
		//
		// # 上传IPv6数据库
		//
		// 上传IPv6数据库，用于上传IPv6数据库操作；
		// 该接口将会对上传的文件进行解码，解码成功后将会将文件写入到 upload/ip_location/database_location_ipv6.scv 文件中；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - file			文件(string)
		//
		// # 返回
		//   - err			错误信息(error)
		IPv6FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error)
		// IPv4FileImport
		//
		// # 导入IPv4数据库
		//
		// 导入IPv4数据库，用于导入IPv4数据库操作；
		// 该接口将会从 upload/ip_location/database_location_ipv4.scv 文件中导入数据到数据库中；
		// 该接口将会清空原有的数据；
		IPv4FileImport(ctx context.Context) (err error)
		// IPv6FileImport
		//
		// # 导入IPv6数据库
		//
		// 导入IPv6数据库，用于导入IPv6数据库操作；
		// 该接口将会从 upload/ip_location/database_location_ipv6.scv 文件中导入数据到数据库中；
		// 该接口将会清空原有的数据；
		IPv6FileImport(ctx context.Context) (err error)
	}
)

var (
	localIP IIP
)

func IP() IIP {
	if localIP == nil {
		panic("implement not found for interface IIP, forgot register?")
	}
	return localIP
}

func RegisterIP(i IIP) {
	localIP = i
}
