/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package startup

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

// InitialDatabaseStartup
//
// # 初始化数据库
//
// 初始化数据库，进行数据库的初始化操作；若检查数据库中没有数据则进行初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
func InitialDatabaseStartup(ctx context.Context) {
	g.Log().Noticef(ctx, "[STARTUP] 初始化数据库...")
	// 检查数据库是否存在
	createDatabase(ctx, "fy_info")
	createDatabase(ctx, "fy_permission")
	createDatabase(ctx, "fy_resource")
	createDatabase(ctx, "fy_role")
	createDatabase(ctx, "fy_vip")
	createDatabase(ctx, "fy_user")
}

// createDatabase
//
// # 创建数据库
//
// 创建数据库，根据传入的 schema 名称进行创建数据库；
//
// # 参数
//   - schema		数据库名称(string)
func createDatabase(ctx context.Context, schema string) {
	// 检查数据表
	record, err := g.Model("information_schema.tables").
		Where("table_name=?", schema).
		One()
	if err != nil {
		g.Log().Panicf(ctx, "[STARTUP] 数据库表检查失败：%s", err.Error())
	}
	if record.Map()["table_name"] != schema {
		g.Log().Debugf(ctx, "[STARTUP] 数据库表不存在，创建 %s 数据表", schema)
		// 读取文件并且根据分号拆分
		_, err := g.DB().Exec(ctx, gfile.GetContents("resource/schema/"+schema+".sql"))
		if err != nil {
			g.Log().Panicf(ctx, "[STARTUP] 数据库表创建失败：%s", err.Error())
		}
	}
}
