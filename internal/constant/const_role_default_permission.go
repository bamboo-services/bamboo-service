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

// ------------------------------------------------------------
// 系统角色权限,用于系统初始化时的角色权限初始化,不可更改
// ------------------------------------------------------------

const (
	PermissionChangePassword    = "auth:change_password" // 修改密码
	PermissionChangePhone       = "auth:change_phone"    // 修改手机号
	PermissionChangeEmail       = "auth:change_email"    // 修改邮箱
	PermissionChangeUsername    = "auth:change_username" // 修改用户名
	PermissionChangeAvatar      = "auth:change_avatar"   // 修改头像
	PermissionChangeInfo        = "auth:change_info"     // 修改信息
	AdminPermissionAcgurlDelete = "acgurl:admin-delete"  // (管理员)删除图库
	AdminPermissionAcgurlEdit   = "acgurl:admin-edit"    // (管理员)编辑图库
)
