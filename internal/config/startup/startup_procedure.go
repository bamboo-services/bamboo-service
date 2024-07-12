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
	"github.com/gogf/gf/v2/frame/g"
)

// initialDatabaseStartup
//
// # 初始化数据库
//
// 初始化数据库，进行数据库的初始化操作；若检查数据库中没有数据则进行初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
func (s *systemStart) initialDatabaseStartup(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 初始化数据库...")
	// 检查数据库是否存在
	createDatabase(ctx, "fy_info")
	createDatabase(ctx, "fy_permission")
	createDatabase(ctx, "fy_resource")
	createDatabase(ctx, "fy_role")
	createDatabase(ctx, "fy_vip")
	createDatabase(ctx, "fy_user")
}

// initialTableContentStartup
//
// # 初始化数据表内容
//
// 初始化数据表内容，进行数据表内容的初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
func (s *systemStart) initialTableContentStartup(ctx context.Context) {
	g.Log().Noticef(ctx, "[STAR] 初始化数据表内容...")

	// Info 表
	g.Log().Infof(ctx, "[STAR] 检查 fy_info 表数据...")
	// 检查数据表
	checkInfoTableValue(ctx, "system_name", "XiaoService")
	checkInfoTableValue(ctx, "system_version", "v1.0.0")
	checkInfoTableValue(ctx, "system_author", "筱锋xiao_lfeng")
}
