/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms

import (
	"context"

	"bamboo-service/api/sms/v2"
)

type ISmsV2 interface {
	SmsMakeAuthorization(ctx context.Context, req *v2.SmsMakeAuthorizationReq) (res *v2.SmsMakeAuthorizationRes, err error)
	SmsSendHasUser(ctx context.Context, req *v2.SmsSendHasUserReq) (res *v2.SmsSendHasUserRes, err error)
	SmsSendNoUser(ctx context.Context, req *v2.SmsSendNoUserReq) (res *v2.SmsSendNoUserRes, err error)
}
