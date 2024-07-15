/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package sms

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/sms/v1"
)

// SmsSend
//
// # 发送短信
//
// 发送短信，用于发送短信验证码，用户注册等操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.SmsSendReq)
//
// # 返回
//   - res			响应(*v1.SmsSendRes)
//   - err			错误信息(error)
func (c *ControllerV1) SmsSend(
	ctx context.Context,
	req *v1.SmsSendReq,
) (res *v1.SmsSendRes, err error) {
	g.Log().Notice(ctx, "[CONT] 发送短信")
	// 检查手机号是否可重发
	err = service.Sms().PhoneCodeAbleResend(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	// 删除旧验证码
	err = service.Sms().DelPhoneCode(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	// 生成随机验证码
	code := butil.RandomString(6)
	// 保存验证码
	err = service.Sms().SetPhoneCode(ctx, req.Phone, code)
	if err != nil {
		return nil, err
	}
	// 发送短信
	err = service.Sms().AliyunSmsSend(ctx, req.Phone, code)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
