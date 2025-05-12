package consts

//
// 邮件发送模板字符串
//

type MailTemplate struct {
	Name        string `json:"name" dc:"模板英文名字"`
	Description string `json:"description" dc:"模板描述"`
	Value       uint   `json:"value" dc:"模板对应值内容Ω"`
}

var MailTemplateForUserForgetPassword = []MailTemplate{
	{Name: "forget_password", Value: 1, Description: "用户忘记密码所发送的邮箱验证码"},
}
