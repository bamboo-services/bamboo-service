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
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

var ctx = gctx.GetInitCtx()

// TestDogeCloudApi
//
// # 测试多吉云 API
//
// 测试多吉云 API 接口；
// 该接口会自动处理签名等信息；
func TestDogeCloudApi(t *testing.T) {
	if false {
		service.RegisterDogeCloud(dogecloud.New())
		getJSON := gjson.New(gfile.GetContents("../access.json"))
		constant.DogeCloudAccessKey = getJSON.Get("DogeCloudKey.AccessKey").String()
		constant.DogeCloudSecretKey = getJSON.Get("DogeCloudKey.SecretKey").String()
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

// TestDogeCloudUploadData
//
// # 测试多吉云上传数据
//
// 测试多吉云上传数据接口；
// 该接口会自动处理签名等信息；
func TestDogeCloudUploadData(t *testing.T) {
	if false {
		service.RegisterDogeCloud(dogecloud.New())
		getJSON := gjson.New(gfile.GetContents("../access.json"))
		constant.DogeCloudAccessKey = getJSON.Get("DogeCloudKey.AccessKey").String()
		constant.DogeCloudSecretKey = getJSON.Get("DogeCloudKey.SecretKey").String()
		// 获取 Cravatar gm@x-lf.cn 的头像
		response, err := g.Client().Get(ctx, "https://cravatar.cn/avatar/76602d1259d6a5e0796933f5d0ff9b84?s=2048")
		if err != nil {
			t.Error(err)
		}
		err = service.DogeCloud().UploadData(
			ctx,
			"/avatar/",
			"76602d1259d6a5e0796933f5d0ff9b84.png",
			response.ReadAll(),
		)
		if err != nil {
			t.Error(err)
		}
	}
}
