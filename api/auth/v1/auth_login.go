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

import (
	"bamboo-service/internal/model/dto"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthLoginReq
//
// # 用户登录
//
// 用户登录，用于用户登录操作；
// 用户登录需要用户名和密码；
//
// # 参数
//   - User		用户名/邮箱/手机号(string)
//   - Pass		密码(string)
type AuthLoginReq struct {
	g.Meta  `path:"/auth/login" method:"Post" summary:"用户登录" tags:"认证控制器"`
	Referer string `json:"Referer" v:"required|url#请输入来源地址|来源地址格式不正确" in:"header"`
	User    string `json:"user" v:"required#请输入用户名" dc:"用户名/邮箱/手机号"`
	Pass    string `json:"pass" v:"required#请输入密码" dc:"密码"`
}

// AuthLoginRes
//
// # 用户登录
//
// 返回相应的数据
type AuthLoginRes struct {
	g.Meta `mime:"application/json"`
	User   dto.UserCurrentDTO `json:"user" orm:"user" dc:"用户信息"`
	Token  string             `json:"token" orm:"token" dc:"令牌"`
}
