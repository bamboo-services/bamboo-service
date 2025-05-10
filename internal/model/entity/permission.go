// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure for table permission.
type Permission struct {
	PermissionKey         string      `json:"permission_key"         orm:"permission_key"         description:"权限标识"` // 权限标识
	PermissionName        string      `json:"permission_name"        orm:"permission_name"        description:"权限名称"` // 权限名称
	PermissionDescription string      `json:"permission_description" orm:"permission_description" description:"权限描述"` // 权限描述
	PermissionStatus      bool        `json:"permission_status"      orm:"permission_status"      description:"权限状态"` // 权限状态
	CreatedAt             *gtime.Time `json:"created_at"             orm:"created_at"             description:"创建时间"` // 创建时间
	UpdatedAt             *gtime.Time `json:"updated_at"             orm:"updated_at"             description:"更新时间"` // 更新时间
}
