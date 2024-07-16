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
	"bamboo-service/internal/logic/sms"
	"bamboo-service/internal/service"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

func TestSmsSend(t *testing.T) {
	if false {
		service.RegisterSms(sms.New())
		// 读取文件
		getJSON := gjson.New(gfile.GetContents("../access.json"))
		constant.AliyunSmsSignName = "锋楪"
		constant.AliyunSmsCodeTemplateCode = "SMS_468930484"
		constant.AliyunAccessKey = getJSON.Get("AliyunKey.AccessKeyID").String()
		constant.AliyunSecretKey = getJSON.Get("AliyunKey.AccessKeySecret").String()
		constant.AliyunSmsEndpoint = "dysmsapi.aliyuncs.com"
		g.Log().Debugf(ctx, "[TEST] %s", gjson.New(constant.AliyunAccessKey).MustToJsonString())
		err := service.Sms().AliyunSmsSend(ctx, "13316569390", butil.RandomString(6))
		if err != nil {
			t.Error(err)
		} else {
			t.Log("短信发送成功")
		}
	}
}
