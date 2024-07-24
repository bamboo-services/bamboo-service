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
	v1 "bamboo-service/api/acgurl/v1"
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/google/uuid"
)

type (
	IAcgurl interface {
		// CheckAlbumExist
		//
		// # 检查图库是否存在
		//
		// 检查图库是否存在，若存在则返回错误；
		// 若不存在则返回 nil；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - name		图库名称(string)
		//
		// # 返回
		//   - err		错误信息(error)
		CheckAlbumExist(ctx context.Context, name string) (err error)
		// CreateAlbum
		//
		// # 创建图库
		//
		// 创建一个图库，用于创建一个图库操作；
		//
		// # 参数
		//   - ctx				上下文(context.Context)
		//   - authorization	授权码(string)
		//   - name				图库名称(string)
		//   - displayName		图库展示名称(string)
		//   - description		图库描述(string)
		//   - cover			图库封面(string)
		//   - visible			图库是否可见(bool)
		//
		// # 返回
		//   - err		错误信息(error)
		CreateAlbum(ctx context.Context, authorization, name, displayName, description, cover string, visible bool) (err error)
		// DeleteAlbum
		//
		// # 删除图库
		//
		// 删除一个图库，用于删除一个图库操作；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - albumUUID	图库唯一标识(uuid.UUID)
		//
		// # 返回
		//   - err		错误信息(error)
		DeleteAlbum(ctx context.Context, albumUUID uuid.UUID, authorization string) (err error)
		// GetAlbumInfo
		//
		// # 获取图库信息
		//
		// 获取图库信息，用于获取图库信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - albumUUID	图库唯一标识(uuid.UUID)
		GetAlbumInfo(ctx context.Context, albumUUID uuid.UUID) (album *dto.AlbumInfoDTO, err error)
		// EditAlbum
		//
		// # 编辑图库
		//
		// 编辑一个图库，用于编辑一个图库操作；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - req		请求(*v1.AcgurlEditAlbumReq)
		//
		// # 返回
		//   - err		错误(error)
		EditAlbum(ctx context.Context, req *v1.AcgurlEditAlbumReq) (err error)
	}
)

var (
	localAcgurl IAcgurl
)

func Acgurl() IAcgurl {
	if localAcgurl == nil {
		panic("implement not found for interface IAcgurl, forgot register?")
	}
	return localAcgurl
}

func RegisterAcgurl(i IAcgurl) {
	localAcgurl = i
}
