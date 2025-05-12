package dto

// MailSendTemplateDTO 表示发送邮件模板时需要的数据结构。
//
// 包含邮箱地址、网站信息、作者信息、验证码及有效时间。
type MailSendTemplateDTO struct {
	Email      string `json:"email" dc:"邮箱"`
	WebName    string `json:"web_name" dc:"网站名称"`
	WebLink    string `json:"web_link" dc:"网站链接"`
	AuthorLink string `json:"author_link" dc:"作者链接"`
	AuthorName string `json:"author_name" dc:"作者名称"`
	Code       string `json:"code" dc:"验证码"`
	Minute     int    `json:"minute" dc:"验证码有效期时间"`
}
