// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-09 15:09:16
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// System is the golang structure for table system.
type System struct {
	SystemUuid string      `json:"system_uuid" orm:"system_uuid" description:"系统唯一标识符"` // 系统唯一标识符
	Key        string      `json:"key"         orm:"key"         description:"键"`       // 键
	Value      string      `json:"value"       orm:"value"       description:"值"`       // 值
	Version    int64       `json:"version"     orm:"version"     description:"版本"`      // 版本
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:"创建时间"`    // 创建时间
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:"更新时间"`    // 更新时间
}
