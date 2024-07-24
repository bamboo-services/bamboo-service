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
	v1 "bamboo-service/api/acgurl/v1"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
func (s *sAcgurl) DeleteAlbum(ctx context.Context, albumUUID uuid.UUID, authorization string) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.DeleteAlbum | 删除图库")
	// 检查图库是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, albumUUID)
	if err != nil {
		return err
	}
	if getAlbum == nil {
		return berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 检查用户权限
	err = service.Auth().IsUserCanDo(ctx, authorization, getAlbum.Uuid, "acgurl:admin-delete")
	if err != nil {
		return err
	}
	// 删除图库
	err = s.deleteAlbum(ctx, butil.StringToUUID(getAlbum.AlbumUuid))
	if err != nil {
		return err
	}
	return nil
}

// GetAlbumInfo
//
// # 获取图库信息
//
// 获取图库信息，用于获取图库信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - albumUUID	图库唯一标识(uuid.UUID)
func (s *sAcgurl) GetAlbumInfo(ctx context.Context, albumUUID uuid.UUID) (album *dto.AlbumInfoDTO, err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.GetAlbumInfo | 获取图库信息")
	// 检查图库是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, albumUUID)
	if err != nil {
		return nil, err
	}
	if getAlbum == nil {
		return nil, berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 获取用户的信息
	getUser, err := service.User().GetUserByUUID(ctx, butil.StringToUUID(getAlbum.Uuid))
	if err != nil {
		return nil, err
	}
	pattern := "未知模式"
	switch getAlbum.ExcludePattern {
	case 0:
		pattern = "常规模式"
	case 1:
		pattern = "黑名单模式"
	case 2:
		pattern = "白名单模式"
	}
	// 对数据进行可视化整理
	err = gconv.Struct(getAlbum, &album, map[string]string{"User": getUser.Username, "Pattern": pattern})
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "数据转换失败")
	}
	return album, nil
}

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
func (s *sAcgurl) EditAlbum(ctx context.Context, req *v1.AcgurlEditAlbumReq) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.EditAlbum | 编辑图库")
	// 检查图库是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, butil.StringToUUID(req.AlbumUUID))
	if err != nil {
		return err
	}
	if getAlbum == nil {
		return berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 检查用户权限
	err = service.Auth().IsUserCanDo(ctx, req.Authorization, getAlbum.Uuid, "acgurl:admin-edit")
	if err != nil {
		return err
	}
	// 编辑图库
	getAlbum.DisplayName = req.AlbumDisplay
	getAlbum.Description = req.AlbumDesc
	getAlbum.Visible = req.AlbumOpen
	getAlbum.Cover = req.AlbumCover
	getAlbum.ExcludePattern = gconv.Int(req.AlbumPattern)
	getAlbum.MatchAddress = gjson.New(req.MatchAddress)
	err = s.editAlbum(ctx, getAlbum)
	if err != nil {
		return err
	}
	return nil
}
