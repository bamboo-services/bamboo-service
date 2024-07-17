/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package startup

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/glog"
)

// systemStart
//
// # 系统启动
//
// 系统启动，用于系统启动时的初始化操作；
type systemStart struct {
	ctx context.Context
}

func newSystem(ctx context.Context) *systemStart {
	return &systemStart{
		ctx: ctx,
	}
}

// SystemStartUp
//
// # 系统启动
//
// 系统启动，用于系统启动时的初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
func SystemStartUp(ctx context.Context) {
	glog.Noticef(ctx, "[STAR] 开始检查系统完整性...")

	/*
	 * 初始化
	 */
	// 初始化准备
	var system = newSystem(ctx)

	glog.Noticef(ctx, "==================================================")
	// 系统初始化
	system.initialDatabaseStartup()
	system.initialTableContentStartup()
	system.initialRoleStartup()
	system.initialSuperAdminStartup()
	system.getConstantStorage()
	system.getAliyunAuthorizationKey()
	system.dogeCloudKey()
	system.checkFolder()

	/*
	 * 系统准备完成
	 */
	glog.Noticef(ctx, "==================================================")

	fmt.Println("\033[1;35m" + `
	   _  ___           ____             _        
	  | |/_(_)__ ____  / __/__ _____  __(_)______ 
	 _>  </ / _ ` + "`" + `/ _ \_\ \/ -_) __/ |/ / / __/ -_)
	/_/|_/_/\_,_/\___/___/\__/_/  |___/_/\__/\__/ `)
	fmt.Println("\033[32m	   ::: XiaoService :::\033[33m	   	v1.0.0")
	fmt.Println("\033[0m")
}
