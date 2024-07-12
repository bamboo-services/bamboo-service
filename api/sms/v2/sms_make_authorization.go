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

// SmsMakeAuthorizationReq
//
// # 生成短信验证码授权码
//
// 生成短信验证码授权码，用于发送短信验证码；
//
// # 请求
//   - authorization_code		授权码，传入后若正确则通知任然有效(string?)
type SmsMakeAuthorizationReq struct {
	g.Meta            `path:"/api/v2/sms/authorization" method:"Get" summary:"生成短信验证码授权码" tags:"短信验证码控制器"`
	AuthorizationCode string `json:"authorization_code" v:"regex:^(|[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12})$#授权码格式不正确" summary:"授权码"`
}

// SmsMakeAuthorizationRes
//
// # 生成短信验证码授权码
//
// 返回相应的数据
//
// # 返回
//   - authorization_code		授权码(string)
type SmsMakeAuthorizationRes struct {
	g.Meta            `mime:"application/json"`
	AuthorizationCode string `json:"authorization_code"`
}
