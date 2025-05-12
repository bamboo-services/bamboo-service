package consts

//
// 邮件发送模板字符串
//

type MailTemplate struct {
	Name        string `json:"name" dc:"模板英文名字"`
	Description string `json:"description" dc:"模板描述"`
}

type MailPurpose struct {
	Name        string `json:"name" dc:"验证码用途"`
	Description string `json:"description" dc:"用途描述"`
}

// MailTemplateList 包含所有支持的邮件模板列表，用于验证和填充邮件内容。
var MailTemplateList = []MailTemplate{
	{Name: "forget_password", Description: "用户忘记密码所发送的邮箱验证码"},
}

// MailPurposeList 包含所有支持的邮件用途列表，用于验证和填充邮件内容。
var MailPurposeList = []MailPurpose{
	{Name: "register", Description: " 用于用户注册账号时激活邮箱使用的验证码"},
	{Name: "forget_password", Description: "用户忘记密码时发送的邮箱验证码"},
	{Name: "login", Description: "用户登录时发送的邮箱验证码"},
	{Name: "bind_email", Description: "用户绑定邮箱时发送的邮箱验证码"},
}
