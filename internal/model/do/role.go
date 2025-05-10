// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Role is the golang structure of table fy_role for DAO operations like Where/Data.
type Role struct {
	g.Meta          `orm:"table:fy_role, do:true"`
	RoleUuid        interface{} //
	RoleName        interface{} // 角色名称
	RoleNickname    interface{} // 角色昵称
	RoleDescription interface{} // 角色描述
	RolePermission  *gjson.Json // 角色权限
	RoleStatus      interface{} // 角色状态(开启和关闭)
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
