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
	"bamboo-service/internal/logic/dogecloud"
	"bamboo-service/internal/service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

var ctx = gctx.GetInitCtx()

func TestDogeCloudApi(t *testing.T) {
	if false {
		service.RegisterDogeCloud(dogecloud.New())
		getJson := gjson.New(gfile.GetContents("../access.json"))
		constant.DogeCloudAccessKey = getJson.Get("DogeCloudKey.AccessKey").String()
		constant.DogeCloudSecretKey = getJson.Get("DogeCloudKey.SecretKey").String()
		//data := gconv.Map(g.Map{
		//	"channel": "OSS_FULL",
		//	"scopes":  "bamboo-service",
		//})
		//ref, err := service.DogeCloud().Api(ctx, "/auth/tmp_token.json", data, true)
		// 处理
		rdo, err := service.DogeCloud().GetAccessTokenApi(ctx)
		if err != nil {
			t.Error(err)
		}
		t.Log(gjson.New(rdo).MustToJsonString())
	}
}
