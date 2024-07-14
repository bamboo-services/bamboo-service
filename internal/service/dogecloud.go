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
	IDogeCloud interface {
		// Api
		//
		// # 多吉云 API 请求
		//
		// 通过该接口可以请求多吉云的 API 接口；
		// 该接口会自动处理签名等信息；
		// 该接口会自动处理 JSON 数据；
		// 该接口会自动处理请求头等信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - apiPath		API 路径(string)
		//   - data			请求数据(map[string]interface{})
		//   - jsonMode		是否为 JSON 模式(bool)
		//
		// # 返回
		//   - ret			返回数据(map[string]interface{})
		//   - err			错误信息(error)
		Api(ctx context.Context, apiPath string, data map[string]interface{}, jsonMode bool) (ref map[string]interface{}, err error)
		// GetAccessTokenApi
		//
		// # 获取多吉云存储 Token 权限
		//
		// 该借口操作不建议单独直接进行调用，在 task 模块进行定时循环调用；
		// 目的为获取权限 Token，用于后续的操作；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//
		// # 返回
		//   - bucket		多吉云存储 Token 信息(*rdo.DogeCloudBucketRDO)
		//   - err			错误信息(error)
		GetAccessTokenApi(ctx context.Context) (bucket *rdo.DogeCloudBucketRDO, err error)
		// GetToken
		//
		// # 获取多吉云存储 Token 权限
		//
		// 该接口为从缓存调用令牌信息；
		// 用于获取多吉云存储 Token 权限；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//
		// # 返回
		//   - bucket		多吉云存储 Token 信息(*rdo.DogeCloudBucketRDO)
		//   - err			错误信息(error)
		GetToken(ctx context.Context) (bucket *rdo.DogeCloudBucketRDO, err error)
	}
)

var (
	localDogeCloud IDogeCloud
)

func DogeCloud() IDogeCloud {
	if localDogeCloud == nil {
		panic("implement not found for interface IDogeCloud, forgot register?")
	}
	return localDogeCloud
}

func RegisterDogeCloud(i IDogeCloud) {
	localDogeCloud = i
}
