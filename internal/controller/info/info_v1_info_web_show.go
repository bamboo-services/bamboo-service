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
	"bamboo-service/internal/model/dto"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/info/v1"
)

// InfoWebShow
//
// # 获取网站信息
//
// 可以通过该接口获取网站的信息，包括网站的名称、描述、关键字、备案号等信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.InfoWebShowReq)
//
// # 返回
//   - res			响应(*v1.InfoWebShowRes)
//   - err			错误信息(error)
func (c *ControllerV1) InfoWebShow(
	ctx context.Context,
	_ *v1.InfoWebShowReq,
) (res *v1.InfoWebShowRes, err error) {
	g.Log().Notice(ctx, "[CONT] 获取网站信息")
	return &v1.InfoWebShowRes{
		InfoWebDTO: dto.InfoWebDTO{
			WebName:        constant.WebName,
			WebDescription: constant.WebDescription,
			WebKeywords:    constant.WebKeywords,
			WebLogo:        constant.WebLogo,
			WebFavicon:     constant.WebFavicon,
			WebICP:         constant.WebICP,
			WebRecord:      constant.WebRecord,
			WebCopy:        constant.WebCopy,
		},
	}, nil
}
