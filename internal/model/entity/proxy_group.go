// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyGroup is the golang structure for table proxy_group.
type ProxyGroup struct {
	GroupUuid   string      `json:"group_uuid"  orm:"group_uuid"  description:"代理组UUID"` // 代理组UUID
	UserUuid    string      `json:"user_uuid"   orm:"user_uuid"   description:"用户UUID"`  // 用户UUID
	Name        string      `json:"name"        orm:"name"        description:"代理组名称"`   // 代理组名称
	FileName    string      `json:"file_name"   orm:"file_name"   description:"代理组文件名"`  // 代理组文件名
	Description string      `json:"description" orm:"description" description:"代理组描述"`   // 代理组描述
	Proxy       *gjson.Json `json:"proxy"       orm:"proxy"       description:"代理组代理"`   // 代理组代理
	Partition   *gjson.Json `json:"partition"   orm:"partition"   description:"代理组分区"`   // 代理组分区
	Rule        *gjson.Json `json:"rule"        orm:"rule"        description:"代理组规则"`   // 代理组规则
	IsEnabled   bool        `json:"is_enabled"  orm:"is_enabled"  description:"代理组是否启用"` // 代理组是否启用
	CreatedAt   *gtime.Time `json:"created_at"  orm:"created_at"  description:"代理组创建时间"` // 代理组创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"  orm:"updated_at"  description:"代理组更新时间"` // 代理组更新时间
}
