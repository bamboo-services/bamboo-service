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
	"bamboo-service/internal/constant"
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"
)

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
func (s *sSms) AliyunSmsSend(ctx context.Context, to string, code string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.AliyunSmsSend | 发送短信")

	// 检查是否是国内的手机号
	match, err := regexp.Match(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, []byte(to))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	if !match {
		return berror.NewError(bcode.OperationFailed, "手机号格式不正确")
	}

	// 创建客户端
	smsParam := gmap.NewHashMap(true)
	smsParam.Set("code", code)
	backCode, err := smsSendCode(ctx, to, gjson.New(smsParam).MustToJsonString())
	if err != nil {
		return err
	} else {
		if *backCode != "OK" {
			return berror.NewError(bcode.OperationFailed, "短信发送失败")
		} else {
			return nil
		}
	}
}

// smsSendCode
//
// # 发送短信
//
// 发送短信，用于发送短信验证码；该接口为内部接口，不对外开放；
// 接口实现了通过阿里云的 SDK 发送短信验证码；
//
// # 请求
//   - ctx			上下文(context.Context)
//   - to			接收者手机号(string)
//   - code			验证码(string)
//
// # 响应
//   - backCode		返回码，短信发送状态(string)
//   - err			错误信息(error)
func smsSendCode(ctx context.Context, to, code string) (backCode *string, err error) {
	g.Log().Debugf(ctx, "[SERV] Aliyun 发送短信")
	config := &openapi.Config{
		AccessKeyId:     &constant.AliyunAccessKey,
		AccessKeySecret: &constant.AliyunSecretKey,
	}
	config.Endpoint = &constant.AliyunSmsEndpoint
	// 创建客户端
	client, err := dysmsapi.NewClient(config)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}

	// 发送短信
	response, err := client.SendSms(&dysmsapi.SendSmsRequest{
		SignName:      tea.String(constant.AliyunSmsSignName),
		TemplateCode:  tea.String(constant.AliyunSmsCodeTemplateCode),
		PhoneNumbers:  tea.String(to),
		TemplateParam: tea.String(code),
	})
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err)
	}
	// 对 Response 进行配置
	if *response.Body.Code != "OK" {
		return nil, berror.NewError(bcode.OperationFailed, *response.Body.Message)
	}
	g.Log().Infof(ctx, "[SERV] 短信已发送 \n\tBizId 码：%s\n\t请求码：%s\n\t状态码：%s\n\t消息描述：%s",
		*response.Body.BizId,
		*response.Body.RequestId,
		*response.Body.Code,
		*response.Body.Message,
	)
	return response.Body.Code, nil
}
