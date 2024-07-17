/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package ip

import (
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

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
func (s *sIP) IPv4FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv4FileUpload | 上传IPv4数据库接口")
	// 检查原文件是否存在
	err = gfile.Remove("upload/ip_location/database_location_ipv4.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件删除失败")
	}
	file.Filename = "database_location_ipv4.scv"
	_, err = file.Save("upload/ip_location/")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件写入失败")
	}
	return nil
}

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
func (s *sIP) IPv6FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv6FileUpload | 上传IPv6数据库接口")
	// 检查原文件是否存在
	err = gfile.Remove("upload/ip_location/database_location_ipv6.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件删除失败")
	}
	file.Filename = "database_location_ipv6.scv"
	_, err = file.Save("upload/ip_location/")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件写入失败")
	}
	return nil
}
