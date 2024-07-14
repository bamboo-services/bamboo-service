/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package rdo

import "github.com/gogf/gf/v2/os/gtime"

type SmsPhoneCodeRDO struct {
	Code      string      `json:"code" summary:"验证码"`
	CreatedAt *gtime.Time `json:"created_at" summary:"创建时间"`
	ExpiredAt *gtime.Time `json:"expired_at" summary:"过期时间"`
}
