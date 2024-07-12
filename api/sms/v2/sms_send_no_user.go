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

// SmsSendNoUserReq
//
// # 向非注册用户发送短信验证码
//
// 向非注册用户发送短信验证码，用于用户注册等不需要进行用户信息验证的操作，大多数情况下不应该使用这个接口；
// 并且应当传递授权码，授权码是为了防止恶意发送短信验证码，授权码是在发送短信验证码之前需要先获取的；
//
// # 参数
//   - AuthorizationCode		授权码(string)
//   - Phone					手机号(string)
type SmsSendNoUserReq struct {
	g.Meta            `path:"/api/v2/sms/send/no-user" method:"Post" summary:"向非注册用户发送短信验证码" tags:"短信验证码控制器"`
	AuthorizationCode string `json:"authorization_code" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$#请输入授权码|授权码格式不正确"`
	Phone             string `json:"phone" v:"required|length:11,11#请输入手机号|手机号长度为11位"`
}

// SmsSendNoUserRes
//
// # 向非注册用户发送短信验证码
//
// 返回相应的数据
type SmsSendNoUserRes struct {
	g.Meta `mime:"application/json"`
}
