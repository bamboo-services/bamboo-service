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
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

// getAlbumByName
//
// # 获取图库信息
//
// 获取图库信息，从缓存中获取，若缓存中不存在则从数据库中获取；
// 若数据库中也不存在则返回错误；
// 若获取成功则返回图库信息；
// 若从数据库取得信息，将存入缓存并设置 6 小时有效期；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - name		图库名称(string)
//
// # 返回
//   - album	图库信息(*entity.Album)
//   - err		错误信息(error)
func (s *sAcgurl) getAlbumByName(ctx context.Context, name string) (album *entity.Album, err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.getAlbumByName | 获取图库信息")
	return s.getAlbumByUUID(ctx, butil.MakeUUIDByString(name))
}

// getAlbumByUUID
//
// # 获取图库信息
//
// 获取图库信息，从缓存中获取，若缓存中不存在则从数据库中获取；
// 若数据库中也不存在则返回错误；
// 若获取成功则返回图库信息；
// 若从数据库取得信息，将存入缓存并设置 6 小时有效期；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - albumUUID	图库 UUID(uuid.UUID)
//
// # 返回
//   - album	图库信息(*entity.Album)
//   - err		错误信息(error)
func (s *sAcgurl) getAlbumByUUID(ctx context.Context, albumUUID uuid.UUID) (album *entity.Album, err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.getAlbumByUUID | 获取图库信息")
	var getAlbum *entity.Album
	hGetAll, err := g.Redis().HGetAll(ctx, "acgurl:album:"+albumUUID.String())
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if err = hGetAll.Scan(&getAlbum); err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 检查是否存在
	if getAlbum == nil {
		// 从数据库获取
		err := dao.Album.Ctx(ctx).Where(do.Album{Name: albumUUID}).Scan(&getAlbum)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
		}
		if getAlbum == nil {
			return nil, berror.NewError(bcode.NotExist, "图库不存在")
		} else {
			// 存入缓存
			if err := g.Redis().HMSet(ctx, "acgurl:album:"+albumUUID.String(), gconv.Map(getAlbum)); err != nil {
				return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
			}
			// 设置有效期 6 小时
			if _, err := g.Redis().Expire(ctx, "acgurl:album:"+albumUUID.String(), 21600); err != nil {
				return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
			}
			return getAlbum, nil
		}
	} else {
		return getAlbum, nil
	}
}

// addAlbum
//
// # 添加图库
//
// 添加图库，用于添加一个图库；
// 若图库已存在则返回错误；
// 若添加成功则返回 nil；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - newAlbum	新图库信息(*entity.Album)
//
// # 返回
//   - err		错误信息(error)
func (s *sAcgurl) addAlbum(ctx context.Context, newAlbum *entity.Album) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.addAlbum | 添加图库")
	// 前置操作
	getUUID := butil.MakeUUIDByString(newAlbum.Name)
	newAlbum.AlbumUuid = getUUID.String()
	if newAlbum.Uuid == "" {
		return berror.NewError(bcode.OperationFailed, "用户 UUID 不能为空")
	}
	// 检查是否存在
	var getAlbum *entity.Album
	if getAlbum, err = s.getAlbumByUUID(ctx, getUUID); err != nil {
		return err
	}
	if getAlbum != nil {
		return berror.NewError(bcode.AlreadyExists, "图库已存在")
	}
	// 添加图库
	if _, err = dao.Album.Ctx(ctx).Insert(newAlbum); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 存入缓存
	if err = dao.Album.Ctx(ctx).Where(do.Album{AlbumUuid: newAlbum.AlbumUuid}).Scan(&getAlbum); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if err = g.Redis().HMSet(ctx, "acgurl:album:"+newAlbum.AlbumUuid, gconv.Map(getAlbum)); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 设置有效期 6 小时
	if _, err = g.Redis().Expire(ctx, "acgurl:album:"+newAlbum.AlbumUuid, 21600); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}

// editAlbum
//
// # 编辑图库
//
// 编辑图库，用于编辑一个图库；
// 若图库不存在则返回错误；
// 若编辑成功则返回 nil；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - editAlbum	编辑图库信息(*entity.Album)
//
// # 返回
//   - err		错误信息(error)
func (s *sAcgurl) editAlbum(ctx context.Context, album *entity.Album) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.editAlbum | 编辑图库")
	// 检查是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, butil.StringToUUID(album.AlbumUuid))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 检查是否存在
	if getAlbum == nil {
		return berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 编辑图库
	if _, err = dao.Album.Ctx(ctx).Where(do.Album{AlbumUuid: album.AlbumUuid}).Update(album); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 更新缓存
	if err = dao.Album.Ctx(ctx).Where(do.Album{AlbumUuid: album.AlbumUuid}).Scan(&getAlbum); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if err = g.Redis().HMSet(ctx, "acgurl:album:"+album.AlbumUuid, gconv.Map(getAlbum)); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 设置有效期 6 小时
	if _, err = g.Redis().Expire(ctx, "acgurl:album:"+album.AlbumUuid, 21600); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}

// deleteAlbum
//
// # 删除图库
//
// 删除图库，用于删除一个图库；
// 若图库不存在则返回错误；
// 若删除成功则返回 nil；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - albumUUID	图库 UUID(uuid.UUID)
//
// # 返回
//   - err		错误信息(error)
func (s *sAcgurl) deleteAlbum(ctx context.Context, albumUUID uuid.UUID) (err error) {
	g.Log().Notice(ctx, "[SERV] acgurl.deleteAlbum | 删除图库")
	// 检查是否存在
	getAlbum, err := s.getAlbumByUUID(ctx, albumUUID)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 检查是否存在
	if getAlbum == nil {
		return berror.NewError(bcode.NotExist, "图库不存在")
	}
	// 删除图库
	if _, err = dao.Album.Ctx(ctx).Where(do.Album{AlbumUuid: albumUUID.String()}).Delete(); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 删除缓存
	if _, err = g.Redis().Del(ctx, "acgurl:album:"+albumUUID.String()); err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	return nil
}
