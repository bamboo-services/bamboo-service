// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-09 21:32:25
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure for table role.
type Role struct {
	RoleUuid        string      `json:"role_uuid"        orm:"role_uuid"        description:""`            //
	RoleName        string      `json:"role_name"        orm:"role_name"        description:"角色名称"`        // 角色名称
	RoleNickname    string      `json:"role_nickname"    orm:"role_nickname"    description:"角色昵称"`        // 角色昵称
	RoleDescription string      `json:"role_description" orm:"role_description" description:"角色描述"`        // 角色描述
	RolePermission  *gjson.Json `json:"role_permission"  orm:"role_permission"  description:"角色权限"`        // 角色权限
	RoleStatus      bool        `json:"role_status"      orm:"role_status"      description:"角色状态(开启和关闭)"` // 角色状态(开启和关闭)
	CreatedAt       *gtime.Time `json:"created_at"       orm:"created_at"       description:"创建时间"`        // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       orm:"updated_at"       description:"更新时间"`        // 更新时间
}
