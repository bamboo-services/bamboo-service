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
	"bamboo-service/internal/model/rdo"
	"context"
	"fmt"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

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
func (s *sSms) GetPhoneCode(ctx context.Context, phone string) (code *rdo.SmsPhoneCodeRDO, err error) {
	g.Log().Notice(ctx, "[SERV] sms.GetPhoneCode | 获取手机验证码")
	hmGet, err := g.Redis().HGetAll(ctx, "sms:code:"+phone)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取手机验证码失败")
	}
	var getCode *rdo.SmsPhoneCodeRDO
	err = hmGet.Scan(&getCode)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "数据解析失败")
	}
	return getCode, nil
}

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
func (s *sSms) SetPhoneCode(ctx context.Context, phone, code string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.SetPhoneCode | 设置手机验证码")
	err = g.Redis().HMSet(ctx, "sms:code:"+phone, gconv.Map(rdo.SmsPhoneCodeRDO{
		Code:      code,
		CreatedAt: gtime.Now(),
		ExpiredAt: gtime.Now().Add(5 * time.Minute),
	}))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "设置验证码失败")
	}
	_, err = g.Redis().Expire(ctx, "sms:code:"+phone, int64(5*time.Minute))
	if err != nil {
		_, err := g.Redis().Del(ctx, "sms:code:"+phone)
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err, "删除验证码失败")
		}
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "设置验证码时间失败")
	}
	return nil
}

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
func (s *sSms) DelPhoneCode(ctx context.Context, phone string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.DelPhoneCode | 删除手机验证码")
	_, err = g.Redis().Del(ctx, "sms:code:"+phone)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "删除验证码失败")
	}
	return nil
}

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
func (s *sSms) PhoneCodeAbleResend(ctx context.Context, phone string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.PhoneCodeAbleResend | 手机验证码可重发")
	code, err := s.GetPhoneCode(ctx, phone)
	if err != nil {
		return err
	}
	fmt.Println(gjson.New(code).MustToJsonString())
	if code.Code == "" {
		return nil
	}
	if code.CreatedAt.Add(2 * time.Minute).After(gtime.Now()) {
		return berror.NewError(bcode.OperationFailed, "验证码依然有效")
	}
	return nil
}

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
func (s *sSms) PhoneCodeVerify(ctx context.Context, phone, code string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.PhoneCodeVerify | 手机验证码验证")
	getCode, err := s.GetPhoneCode(ctx, phone)
	if err != nil {
		return err
	}
	if getCode.Code != code {
		return berror.NewError(bcode.VerifyFailed, "验证码错误")
	}
	// 验证成功删除验证码
	err = s.DelPhoneCode(ctx, phone)
	if err != nil {
		return err
	}
	return nil
}
