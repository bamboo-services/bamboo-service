package dto

import "github.com/gogf/gf/v2/os/gtime"

// MailCodeDTO 邮件验证码的传输数据对象。
//
// 封装邮箱、验证码及其创建时间，用于传输邮件验证码相关信息。
type MailCodeDTO struct {
	Email     string      `json:"email" dc:"邮箱地址"`
	Code      string      `json:"code,omitempty" dc:"验证码"`
	CreatedAt *gtime.Time `json:"created_at" dc:"创建时间"`
}
