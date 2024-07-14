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

// DogeCloudBucketRDO
//
// # 多吉云存储桶令牌实体
//
// 多吉云存储桶令牌实体，用于存储多吉云存储桶的认证信息；
//
// # 参数
//   - Name				存储桶名称(string)
//   - Bucket			存储桶(string)
//   - Endpoint			存储桶地址(string)
//   - EndpointHost		存储桶地址(string)
//   - AccessKeyID		多吉云认证ID(string)
//   - SecretAccessKey	多吉云认证密钥(string)
//   - SessionToken		多吉云认证Token(string)
type DogeCloudBucketRDO struct {
	Name            string      `json:"name" dc:"存储桶名称"`
	Bucket          string      `json:"bucket" dc:"存储桶"`
	Endpoint        string      `json:"endpoint" dc:"存储桶地址"`
	EndpointHost    string      `json:"endpoint_host" dc:"存储桶地址"`
	AccessKeyID     string      `json:"access_key_id" dc:"多吉云认证ID"`
	SecretAccessKey string      `json:"secret_access_key" dc:"多吉云认证密钥"`
	SessionToken    string      `json:"session_token" dc:"多吉云认证Token"`
	ExpiredAt       *gtime.Time `json:"expired_at" dc:"过期时间"`
}
