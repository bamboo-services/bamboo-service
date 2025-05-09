// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-10 03:00:04
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure of table fy_permission for DAO operations like Where/Data.
type Permission struct {
	g.Meta                `orm:"table:fy_permission, do:true"`
	PermissionKey         interface{} // 权限标识
	PermissionName        interface{} // 权限名称
	PermissionDescription interface{} // 权限描述
	PermissionStatus      interface{} // 权限状态
	CreatedAt             *gtime.Time // 创建时间
	UpdatedAt             *gtime.Time // 更新时间
}
