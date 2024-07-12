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
)

func (s *sSms) AliyunSmsSend(ctx context.Context, to string, code string) (err error) {
	g.Log().Notice(ctx, "[SERV] sms.AliyunSmsSend | 发送短信")

	// TODO-[24071201] 对于发送者的 to 需要进行正则表达式判断，并且从缓存读取是否可以再次发送，需要整理 Json 发送短信
	// 创建客户端
	smsParam := gmap.NewHashMap(true)
	smsParam.Set("code", code)
	backCode, err := smsSendCode(ctx, to, gjson.New(smsParam).MustToJsonString())
	if err != nil {
		return err
	} else {
		if backCode != nil {
			return berror.NewError(bcode.OperationFailed, "短信发送失败")
		} else {
			return nil
		}
	}
}

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
