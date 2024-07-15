/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"
)

type (
	IMail interface {
		// SendCodeMail
		//
		// # 发送验证码邮件
		//
		// 发送验证码邮件，发送验证码邮件到指定邮箱；
		// 验证码将发送到指定 mail 邮箱中，验证码为 code；
		// 验证码的有效期为 5 分钟；该接口将会为输入的 code 进行存入缓存中，用于后续的验证；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - mail		邮箱(string)
		//   - code		验证码(string)
		//
		// # 返回
		//   - error	错误信息(error)
		SendCodeMail(ctx context.Context, mail, code string) (err error)
		// SendMail
		//
		// # 发送邮件
		//
		// 发送邮件，发送邮件到指定邮箱；
		//
		// # 参数
		//   - ctx		上下文(context.Context)
		//   - mail		邮箱(string)
		//   - title	标题(string)
		//   - tpl		模板(string)
		//   - value	自定义参数([]dto.MailVariableDTO)
		//
		// # 返回
		//   - error	错误信息(error)
		SendMail(ctx context.Context, mail, title, tpl string, value []dto.MailVariableDTO) (err error)
	}
)

var (
	localMail IMail
)

func Mail() IMail {
	if localMail == nil {
		panic("implement not found for interface IMail, forgot register?")
	}
	return localMail
}

func RegisterMail(i IMail) {
	localMail = i
}
