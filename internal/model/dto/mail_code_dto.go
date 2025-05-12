package dto

import "github.com/gogf/gf/v2/os/gtime"

// MailCodeDTO 邮件验证码传输对象，用于封装 email 及创建时间等信息。
type MailCodeDTO struct {
	Email     string      `json:"email"`
	CreatedAt *gtime.Time `json:"created_at"`
}
