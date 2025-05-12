// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IMail interface {
		// GetMailCode 根据邮箱地址和用途生成并返回邮件验证码。
		//
		// 功能概述:
		// 检查邮件验证码的发送频率，避免频繁发送；若符合条件，则生成新的验证码并缓存相关信息。
		//
		// 参数:
		//   - ctx: 上下文，用于控制生命周期和传递信息。
		//   - email: 邮箱地址，接收验证码的目标地址。
		//   - purpose: 验证码的用途，如注册、密码重置等。
		//
		// 返回:
		//   - *dto.MailCodeDTO: 封装邮箱、验证码及创建时间的传输数据对象。
		//   - *berror.ErrorCode: 错误信息，包含缓存、频率限制或验证码生成失败的原因。
		GetMailCode(ctx context.Context, email string, purpose string) (*dto.MailCodeDTO, *berror.ErrorCode)
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
