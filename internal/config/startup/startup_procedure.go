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
	"bamboo-service/internal/constant"
	"github.com/gogf/gf/v2/frame/g"
)

// initialDatabaseStartup
//
// # 初始化数据库
//
// 初始化数据库，进行数据库的初始化操作；若检查数据库中没有数据则进行初始化操作；
func (s *systemStart) initialDatabaseStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查数据库")
	// 检查数据库是否存在
	createDatabase(s.ctx, "fy_info")
	createDatabase(s.ctx, "fy_permission")
	createDatabase(s.ctx, "fy_resource")
	createDatabase(s.ctx, "fy_role")
	createDatabase(s.ctx, "fy_vip")
	createDatabase(s.ctx, "fy_user")
}

// initialTableContentStartup
//
// # 初始化数据表内容
//
// 初始化数据表内容，进行数据表内容的初始化操作；
func (s *systemStart) initialTableContentStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查数据表内容...")

	// Info 表
	g.Log().Infof(s.ctx, "\t检查 fy_info 表数据...")
	// 检查数据表
	checkInfoTableValue(s.ctx, "system_name", "竹业")
	checkInfoTableValue(s.ctx, "system_version", "v1.0.0")
	checkInfoTableValue(s.ctx, "system_author", "筱锋xiao_lfeng")
	checkInfoTableValue(s.ctx, "has_initial_mode", "1")
}

func (s *systemStart) initialRoleStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查角色")

	initializeRole(s.ctx, "admin", "管理员", constant.AdminRolePermission, "管理员角色，用于管理全系统级别权限管理等")
	initializeRole(s.ctx, "user", "用户", constant.UserRolePermission, "用户角色，用于管理用户级别权限管理等")
	initializeRole(s.ctx, "bad", "黑名单", constant.BadRolePermission, "黑名单角色，用于管理黑名单级别权限管理等")
}

// initialSuperAdminStartup
//
// # 初始化超级管理员
//
// 初始化超级管理员，进行超级管理员的初始化操作；
func (s *systemStart) initialSuperAdminStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查超级管理员")
	// 检查超级管理员是否存在
	if !hasSuperAdmin(s.ctx) {
		// 创建超级管理员
		createSuperAdmin(s.ctx)
		g.Log().Noticef(s.ctx, "\t用户名：%s", "superAdmin")
		g.Log().Noticef(s.ctx, "\t密码：%s", "admin")
	}
}

func (s *systemStart) getConstantStorage() {

}
