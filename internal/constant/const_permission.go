/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package constant

var (
	// AdminRolePermission
	//
	// # 管理员角色权限
	//
	// 管理员角色权限，用于管理管理员角色的权限；
	AdminRolePermission = []string{
		PermissionChangePassword,
		AdminPermissionAcgurlDelete,
		AdminPermissionAcgurlEdit,
	}

	// UserRolePermission
	//
	// # 用户角色权限
	//
	// 用户角色权限，用于管理用户角色的权限；
	UserRolePermission = []string{
		PermissionChangePassword,
	}

	// BadRolePermission
	//
	// # 黑名单角色权限
	//
	// 黑名单角色权限，用于管理黑名单角色的权限；
	BadRolePermission = []string{
		PermissionChangePassword,
	}
)
