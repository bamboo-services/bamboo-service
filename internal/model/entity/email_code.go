// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// EmailCode is the golang structure for table email_code.
type EmailCode struct {
	CodeUuid  string      `json:"code_uuid"  orm:"code_uuid"  description:"验证码主键"`                                // 验证码主键
	Email     string      `json:"email"      orm:"email"      description:"邮箱地址"`                                 // 邮箱地址
	Code      string      `json:"code"       orm:"code"       description:"验证码"`                                  // 验证码
	Purpose   string      `json:"purpose"    orm:"purpose"    description:"验证码用途：register-注册,reset-重置密码,bind-绑定"` // 验证码用途：register-注册,reset-重置密码,bind-绑定
	ExpiredAt *gtime.Time `json:"expired_at" orm:"expired_at" description:"过期时间"`                                 // 过期时间
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:"创建时间"`                                 // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:"更新时间"`                                 // 更新时间
}
