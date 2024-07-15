/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package v1

import "github.com/gogf/gf/v2/frame/g"

type MailSendReq struct {
	g.Meta  `path:"/mail/send" method:"Post" tags:"邮件控制器" summary:"发送邮件" dc:"发送邮件，用于发送邮件验证码，用户注册等操作"`
	Referer string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	Mail    string `json:"mail" v:"required|email#请输入邮箱|邮箱格式不正确" summary:"邮箱"`
}

type MailSendRes struct {
	g.Meta `mime:"application/json"`
}
