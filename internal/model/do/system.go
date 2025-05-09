// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-10 03:00:04
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// System is the golang structure of table fy_system for DAO operations like Where/Data.
type System struct {
	g.Meta     `orm:"table:fy_system, do:true"`
	SystemUuid interface{} // 系统唯一标识符
	Key        interface{} // 键
	Value      interface{} // 值
	Version    interface{} // 版本
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
