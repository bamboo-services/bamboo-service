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

// SmsSendHasUserReq
//
// # 向已注册用户发送短信验证码
//
// 向已注册用户发送短信验证码，用于用户登录等需要进行用户信息验证的操作；
//
// # 参数
//   - Phone		手机号(string)
type SmsSendHasUserReq struct {
	g.Meta `path:"/api/v2/sms/send/user" method:"Post" summary:"向已注册用户发送短信验证码" tags:"短信验证码控制器"`
	Phone  string `json:"phone" v:"required|length:11,11#请输入手机号|手机号长度为11位"`
}

// SmsSendHasUserRes
//
// # 向已注册用户发送短信验证码
//
// 返回相应的数据
type SmsSendHasUserRes struct {
	g.Meta `mime:"application/json"`
}
