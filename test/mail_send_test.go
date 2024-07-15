/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package test

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/logic/mail"
	"bamboo-service/internal/model/dto"
	_ "bamboo-service/internal/packed"
	"bamboo-service/internal/service"
	"github.com/bamboo-services/bamboo-utils/butil"
	"testing"
)

// TestMailSend
//
// # 测试邮件发送
//
// 测试邮件发送，测试邮件发送功能是否正常；
func TestMailSend(t *testing.T) {
	if true {
		service.RegisterMail(mail.New())
		// 注册变量
		constant.MailSMTPHost = "smtp.qiye.aliyun.com"
		constant.MailSMTPPort = "25"
		constant.MailUser = "noreplay@x-lf.cn"
		constant.MailPassword = "password"
		constant.MailNickname = "锋楪技术机器人"
		constant.WebCopy = "&copy️ 2023-2024 锋楪技术"
		// 发送邮件
		// 添加变量
		value := make([]dto.MailVariableDTO, 0)
		value = append(value, dto.MailVariableDTO{
			Key:   "code",
			Value: butil.RandomString(6),
		})
		err := service.Mail().SendMail(ctx, "gm@x-lf.cn", "测试邮件", "mail_code", value)
		if err != nil {
			t.Error(err)
		}
	}
}
