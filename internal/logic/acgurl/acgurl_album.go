/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package acgurl

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/entity"
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

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
func (s *sAcgurl) CheckAlbumExist(ctx context.Context, name string) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.CheckAlbumExist | 检查图库是否存在")
	var getAlbum *entity.Album
	if getAlbum, err = s.getAlbumByName(ctx, name); err != nil {
		return err
	}
	// 检查是否存在
	if getAlbum != nil {
		return berror.NewError(bcode.AlreadyExists, "图库已存在")
	}
	return nil
}

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
func (s *sAcgurl) CreateAlbum(
	ctx context.Context,
	authorization, name, displayName, description, cover string,
	visible bool,
) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.CreateAlbum | 创建图库")
	// 解析登录用户的信息
	getUser, err := service.Auth().GetUserByAuthorization(ctx, authorization)
	if err != nil {
		return err
	}
	// 检查图库是否存在
	if err = s.CheckAlbumExist(ctx, name); err != nil {
		return err
	}
	// 创建图库
	album := &entity.Album{
		AlbumUuid:   butil.StringToUUID(name).String(),
		Uuid:        getUser.Uuid,
		Name:        name,
		DisplayName: displayName,
		Description: description,
		Cover:       cover,
		Visible:     visible,
	}
	if _, err = dao.Album.Ctx(ctx).Insert(album); err != nil {
		return err
	}
	return nil
}

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
func (s *sAcgurl) DeleteAlbum(ctx context.Context, albumUUID uuid.UUID) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.DeleteAlbum | 删除图库")
	// 检查图库是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, albumUUID)
	if err != nil {
		return err
	}
	if getAlbum == nil {
		return berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 删除图库
	err = s.deleteAlbum(ctx, butil.StringToUUID(getAlbum.AlbumUuid))
	if err != nil {
		return err
	}
	return nil
}
