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

// SmsSendReq
//
// # 发送短信
//
// 发送短信，用于发送短信验证码，用户注册等操作；
//
// # 参数
//   - Phone		手机号(string)
type SmsSendReq struct {
	g.Meta `path:"/sms" method:"Post" tags:"短信控制器" summary:"发送短信" dc:"发送短信，用于发送短信验证码，用户注册等操作"`
	Phone  string `json:"phone" v:"required|length:11,11#请输入手机号|手机号长度为11位" summary:"手机号"`
}

// SmsSendRes
//
// # 发送短信
//
// 返回相应的数据；
type SmsSendRes struct {
	g.Meta `mime:"application/json"`
}
