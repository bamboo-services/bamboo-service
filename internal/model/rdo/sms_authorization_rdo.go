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

// RedisSmsAuthorization
//
// # Redis 短信授权码
//
// Redis 短信授权码，用于存储短信授权码的内容；
//
// # 参数
//   - LastSendAt		最后发送时间(*gtime.Time)
//   - Frequency		频率(int)
//   - SendingUUID		发送 UUID(string)
type RedisSmsAuthorization struct {
	LastSendAt  *gtime.Time `json:"last_send_at"`
	Frequency   int64       `json:"frequency"`
	SendingUUID string      `json:"sending_uuid"`
}
