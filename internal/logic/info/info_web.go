/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package info

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebEditCommon
//
// # 编辑网站信息
//
// 可以通过该接口编辑网站的信息，包括网站的名称、描述、关键字等信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - webName		网站名称(string)
//   - webDesc		网站描述(string)
//   - webKeywords	网站关键字(string)
//   - webLogo		网站Logo(string)
//   - webFavicon	网站Favicon(string)
//   - webCopy		网站版权(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sInfo) WebEditCommon(
	ctx context.Context,
	webName, webDesc, webKeywords, webLogo, webFavicon, webCopy string,
) (err error) {
	g.Log().Notice(ctx, "[SERV] info.WebEditCommon | 编辑网站信息")
	// 编辑网站信息
	err = dao.Info.Transaction(ctx, func(_ context.Context, tx gdb.TX) error {
		_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "web_name"}).Update(do.Info{Value: webName})
		if err != nil {
			return err
		}
		_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "web_description"}).Update(do.Info{Value: webDesc})
		if err != nil {
			return err
		}
		_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "web_keywords"}).Update(do.Info{Value: webKeywords})
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Model(dao.Info.Table).Where(do.Info{Key: "web_logo"}).Update(do.Info{Value: webLogo})
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Model(dao.Info.Table).Where(do.Info{Key: "web_favicon"}).Update(do.Info{Value: webFavicon})
		if err != nil {
			return err
		}
		_, err = tx.Ctx(ctx).Model(dao.Info.Table).Where(do.Info{Key: "web_copy"}).Update(do.Info{Value: webCopy})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "编辑网站信息失败")
	}
	// 修改全局变量
	constant.WebName = webName
	constant.WebDescription = webDesc
	constant.WebKeywords = webKeywords
	constant.WebLogo = webLogo
	constant.WebFavicon = webFavicon
	constant.WebCopy = webCopy

	return nil
}

// WebEditFiling
//
// # 编辑网站备案信息
//
// 可以通过该接口编辑网站的备案信息，包括备案号、备案名称信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - webICP		网站备案号(string)
//   - webRecord	网站网安备案号(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sInfo) WebEditFiling(ctx context.Context, webICP, webRecord string) (err error) {
	g.Log().Notice(ctx, "[SERV] info.WebEditFiling | 编辑网站备案信息")
	// 编辑网站备案信息
	err = dao.Info.Transaction(ctx, func(_ context.Context, _ gdb.TX) error {
		_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "web_icp"}).Update(do.Info{Value: webICP})
		if err != nil {
			return err
		}
		_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "web_record"}).Update(do.Info{Value: webRecord})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "编辑网站备案信息失败")
	}
	// 修改全局变量
	constant.WebICP = webICP
	constant.WebRecord = webRecord

	return nil
}
