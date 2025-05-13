// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyGroup is the golang structure of table fy_proxy_group for DAO operations like Where/Data.
type ProxyGroup struct {
	g.Meta      `orm:"table:fy_proxy_group, do:true"`
	GroupUuid   interface{} // 代理组UUID
	UserUuid    interface{} // 用户UUID
	Name        interface{} // 代理组名称
	FileName    interface{} // 代理组文件名
	Description interface{} // 代理组描述
	Proxy       *gjson.Json // 代理组代理
	Partition   *gjson.Json // 代理组分区
	Rule        *gjson.Json // 代理组规则
	IsEnabled   interface{} // 代理组是否启用
	CreatedAt   *gtime.Time // 代理组创建时间
	UpdatedAt   *gtime.Time // 代理组更新时间
}
