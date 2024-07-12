/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package v2

import "github.com/gogf/gf/v2/frame/g"

// AuthRegisterReq
//
// # 用户注册
//
// 用户注册，用于用户注册，需要传递用户名、手机号、密码、短信验证码；
//
// # 参数
//   - Username		用户名(string)
//   - Phone		手机号(string)
//   - Password		密码(string)
//   - SmsCode		短信验证码(string)
type AuthRegisterReq struct {
	g.Meta   `path:"/api/v2/auth/register" method:"Post" summary:"用户注册" tags:"用户控制器"`
	Username string `json:"username" v:"required|length:6,30#请输入用户名|用户名长度为 6-30 位"`
	Phone    string `json:"phone" v:"required|length:11,11#请输入手机号|手机号长度为11位"`
	Password string `json:"password" v:"required|length:6,30#请输入密码|密码长度为 6-30 位"`
	SmsCode  string `json:"sms_code" v:"required|regex:^[0-9]{6,10}#请输入短信验证码|短信验证码为 6-10 位"`
}

// AuthRegisterRes
//
// # 用户注册
//
// 返回相应的数据
type AuthRegisterRes struct {
	g.Meta `mime:"application/json"`
}
