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

// AuthInitialReq
//
// # 用户初始化
//
// 用户初始化，用于用户初始化，需要传递用户名、手机号、密码、邮箱；
// 用作系统初始化的用户，为超级管理员账户；
// 拥有整个系统最高权限的账户；
//
// # 参数
//   - Username		用户名(string)
//   - Phone		手机号(string)
//   - Password		密码(string)
//   - Email		邮箱(string)
type AuthInitialReq struct {
	g.Meta   `path:"/auth/init" method:"Post" summary:"用户初始化" tags:"用户控制器"`
	Username string `json:"username" v:"required|length:6,30#请输入用户名|用户名长度为 6-30 位"`
	Phone    string `json:"phone" v:"required|length:11,11#请输入手机号|手机号长度为11位"`
	Password string `json:"password" v:"required|length:6,40#请输入密码|密码长度为 6-40 位"`
	Email    string `json:"email" v:"required|email#请输入邮箱|邮箱格式不正确"`
	Referer  string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
}

// AuthInitialRes
//
// # 用户初始化
//
// 返回相应的数据
type AuthInitialRes struct {
	g.Meta `mime:"application/json"`
}
