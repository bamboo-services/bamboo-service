/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package mail

import (
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"

	"bamboo-service/api/mail/v1"
)

// MailSend
//
// # 发送邮件
//
// 发送邮件，用于发送邮件验证码，用户注册等操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - req			请求(*v1.MailSendReq)
//
// # 返回
//   - res			响应(*v1.MailSendRes)
//   - err			错误信息(error)
func (c *ControllerV1) MailSend(
	ctx context.Context,
	req *v1.MailSendReq,
) (res *v1.MailSendRes, err error) {
	g.Log().Notice(ctx, "[CONT] 发送邮件")
	// 生成验证码
	randomString := butil.RandomString(6)
	err = service.Mail().SendCodeMail(ctx, req.Mail, randomString)
	return nil, err
}
