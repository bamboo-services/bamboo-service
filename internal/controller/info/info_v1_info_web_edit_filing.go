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
	"bamboo-service/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/info/v1"
)

// InfoWebEditFiling
//
// # 修改网站备案信息
//
// 可以通过该接口修改网站的备案信息，包括备案号、备案名称信息；
// 方便直接进行热修改，而无需修改代码；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.InfoWebEditFilingReq)
//
// # 返回
//   - res			响应(*v1.InfoWebEditFilingRes)
//   - err			错误信息(error)
func (c *ControllerV1) InfoWebEditFiling(
	ctx context.Context,
	req *v1.InfoWebEditFilingReq,
) (res *v1.InfoWebEditFilingRes, err error) {
	g.Log().Notice(ctx, "[CONT] 修改网站备案信息")
	// 修改网站备案信息
	err = service.Info().WebEditFiling(ctx, req.WebICP, req.WebRecord)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
