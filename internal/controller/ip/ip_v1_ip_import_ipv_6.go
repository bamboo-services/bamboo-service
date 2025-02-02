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
	"bamboo-service/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"

	"bamboo-service/api/ip/v1"
)

// IPImportIpv6
//
// # 导入IPv6数据库
//
// 导入IPv6数据库，用于导入IPv6数据库操作；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - req		请求(*v1.IPImportIpv6Req)
//
// # 返回
//   - res		响应(*v1.IPImportIpv6Res)
//   - err		错误信息(error)
func (c *ControllerV1) IPImportIpv6(
	ctx context.Context,
	req *v1.IPImportIpv6Req,
) (res *v1.IPImportIpv6Res, err error) {
	g.Log().Notice(ctx, "[CONT] 导入IPv6数据库")
	// 授权检查
	err = service.Auth().CheckUserHasSuperAdmin(ctx, req.Authorization)
	if err != nil {
		return nil, err
	}
	// 多线程导入IPv6数据库
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		newCTX := gctx.New()
		err = service.IP().IPv6FileImport(newCTX)
		if err != nil {
			g.Log().Warningf(newCTX, "[CONT] 导入IPv6数据库失败：%s", err.Error())
		}
	}()
	return nil, nil
}
