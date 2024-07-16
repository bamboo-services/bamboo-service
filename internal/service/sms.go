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
	"bamboo-service/internal/model/rdo"
	"context"
)

type (
	ISms interface {
		// AliyunSmsSend
		//
		// # 阿里云短信发送
		//
		// 阿里云短信发送，用于发送短信验证码；该接口会对手机号进行正则表达式判断，判断是否是国内手机号；
		// 若为国内手机号则发送短信验证码，否则返回错误；
		// 当该手机号多次发送短信验证码时，对是否可重复发送判断由阿里云平台进行判断；
		// https://dysms.console.aliyun.com/msgsetting/frequency
		//
		// # 请求
		//   - ctx			上下文(context.Context)
		//   - to			接收者手机号(string)
		//   - code			验证码(string)
		//
		// # 响应
		//   - err			错误信息(error)
		AliyunSmsSend(ctx context.Context, to string, code string) (err error)
		// GetPhoneCode
		//
		// # 获取手机验证码
		//
		// 获取手机验证码，用于获取手机验证码；用于获取缓存中的手机验证码；
		// 该验证码会在 5 分钟后自动失效；失效后无法获取数据
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//
		// # 返回
		//   - code		验证码(*rdo.SmsPhoneCodeRDO)
		//   - err		错误信息(error)
		GetPhoneCode(ctx context.Context, phone string) (code *rdo.SmsPhoneCodeRDO, err error)
		// SetPhoneCode
		//
		// # 设置手机验证码
		//
		// 设置手机验证码，用于设置手机验证码；用于暂时存储短信验证码；
		// 该验证码会在 5 分钟后自动删除；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//   - code			验证码(string)
		//
		// # 返回
		//   - err			错误信息(error)
		SetPhoneCode(ctx context.Context, phone, code string) (err error)
		// DelPhoneCode
		//
		// # 删除手机验证码
		//
		// 删除手机验证码，用于删除手机验证码；用于删除缓存中的手机验证码；
		// 该验证码会在 5 分钟后自动失效；在失效前可以进行删除，失效后删除将无作用
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//
		// # 返回
		//   - err			错误信息(error)
		DelPhoneCode(ctx context.Context, phone string) (err error)
		// PhoneCodeAbleResend
		//
		// # 手机验证码可重发
		//
		// 手机验证码可重发，用于判断手机验证码是否可重发；用于判断手机验证码是否可重发；
		// 验证码的有效期为五分钟，若验证码超过生成的两分钟则可重发；
		// 否则系统将会拒绝重发验证码；
		// 保证验证码系统，防止盗刷行为以及错误发送行为；
		// 当 error 为 nil 时，表示可以重发验证码；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//
		// # 返回
		//   - err			错误信息(error)
		PhoneCodeAbleResend(ctx context.Context, phone string) (err error)
		// PhoneCodeVerify
		//
		// # 手机验证码验证
		//
		// 手机验证码验证，用于验证手机验证码；
		// 该验证码会在 5 分钟后自动失效，失效后无法获取数据；
		// 验证成功后将会返回 nil 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//   - code			验证码(string)
		//
		// # 返回
		//   - err			错误信息(error)
		PhoneCodeVerify(ctx context.Context, phone, code string) (err error)
	}
)

var (
	localSms ISms
)

func Sms() ISms {
	if localSms == nil {
		panic("implement not found for interface ISms, forgot register?")
	}
	return localSms
}

func RegisterSms(i ISms) {
	localSms = i
}
