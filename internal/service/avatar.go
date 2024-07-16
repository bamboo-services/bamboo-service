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
	"bamboo-service/internal/model/dto"
	"context"
)

type (
	IAvatar interface {
		// AvatarGetAPI
		//
		// # 获取头像
		//
		// 获取头像，通过头像的 hash 值获取头像的内容；
		// 通过头像的 hash 值获取头像的内容，返回头像的内容；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - hash		头像的 hash 值(string)
		//
		// # 返回
		//   - avatar	头像内容(dto.AvatarAPIDTO)
		//   - err		错误信息(error)
		AvatarGetAPI(ctx context.Context, hash string) (avatar *dto.AvatarAPIDTO, err error)
		// AvatarGet
		//
		// # 获取头像
		//
		// 获取头像，通过头像的 hash 值获取头像的内容；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - hash		头像的 hash 值(string)
		//
		// # 返回
		//   - avatar	头像内容([]byte)
		AvatarGet(ctx context.Context, hash string) (avatar []byte)
	}
)

var (
	localAvatar IAvatar
)

func Avatar() IAvatar {
	if localAvatar == nil {
		panic("implement not found for interface IAvatar, forgot register?")
	}
	return localAvatar
}

func RegisterAvatar(i IAvatar) {
	localAvatar = i
}
