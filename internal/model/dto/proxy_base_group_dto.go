package dto

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyBaseGroupDTO 表示代理基础组的数据传输对象。
//
// 包含代理组的基本信息和配置属性，用于在系统中传递代理组相关数据。
type ProxyBaseGroupDTO struct {
	GroupUuid   string      `json:"group_uuid"  description:"代理组UUID"`
	UserUuid    string      `json:"user_uuid"   description:"用户UUID"`
	Name        string      `json:"name"        description:"代理组名称"`
	FileName    string      `json:"file_name"   description:"代理组文件名"`
	Description string      `json:"description" description:"代理组描述"`
	IsEnabled   bool        `json:"is_enabled"  description:"代理组是否启用"`
	CreatedAt   *gtime.Time `json:"created_at"  description:"代理组创建时间"`
	UpdatedAt   *gtime.Time `json:"updated_at"  description:"代理组更新时间"`
}
