/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package avatar

import (
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
)

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
//   - avatar	头像内容(dto.AvatarApiDTO)
//   - err		错误信息(error)
func (s *sAvatar) AvatarGetAPI(ctx context.Context, hash string) (avatar *dto.AvatarApiDTO, err error) {
	g.Log().Notice(ctx, "[SERV] avatar.AvatarGetAPI | 获取头像接口")
	// 从 CDN 获取头像检查是否存在
	client := g.Client()
	getCDNResponse, err := client.Get(ctx, "http://b-cdn.api-fy.cn/avatar/"+hash+".png")
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取头像失败")
	}
	if getCDNResponse.StatusCode == 200 {
		return &dto.AvatarApiDTO{
			Avatar: "http://b-cdn.api-fy.cn/avatar/" + hash + ".png",
		}, nil
	}

	client = g.Client()
	response, err := client.Get(ctx, "https://cravatar.cn/avatar/"+hash+"?s=2048")
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取头像失败")
	}
	if response.StatusCode == 200 {
		// 数据保存上传
		err := service.DogeCloud().UploadData(ctx, "/avatar/", hash+".png", response.ReadAll())
		if err != nil {
			return nil, err
		}
		// 数据返回
		return &dto.AvatarApiDTO{
			Avatar: "https://cravatar.cn/avatar/" + hash,
		}, nil
	}
	return nil, nil
}

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
func (s *sAvatar) AvatarGet(ctx context.Context, hash string) (avatar []byte) {
	g.Log().Notice(ctx, "[SERV] avatar.AvatarGet | 获取头像")
	// 从 CDN 获取头像检查是否存在
	getCDNResponse, err := g.Client().Get(ctx, "http://b-cdn.api-fy.cn/avatar/"+hash+".png")
	if err != nil {
		return nil
	}
	if getCDNResponse.StatusCode == 200 {
		return getCDNResponse.ReadAll()
	}
	// 从 Cravatar 获取
	response, err := g.Client().Get(ctx, "https://cravatar.cn/avatar/"+hash+"?s=2048")
	if err != nil {
		return nil
	}
	if response.StatusCode == 200 {
		// 数据保存上传
		err := service.DogeCloud().UploadData(ctx, "/avatar/", hash+".png", response.ReadAll())
		if err != nil {
			return nil
		}
		return response.ReadAll()
	}
	return nil
}
