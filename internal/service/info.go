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
	"context"
)

type (
	IInfo interface {
		// WebEditCommon
		//
		// # 编辑网站信息
		//
		// 可以通过该接口编辑网站的信息，包括网站的名称、描述、关键字等信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - webName		网站名称(string)
		//   - webDesc		网站描述(string)
		//   - webKeywords	网站关键字(string)
		//   - webLogo		网站Logo(string)
		//   - webFavicon	网站Favicon(string)
		//   - webCopy		网站版权(string)
		//
		// # 返回
		//   - err			错误信息(error)
		WebEditCommon(ctx context.Context, webName, webDesc, webKeywords, webLogo, webFavicon, webCopy string) (err error)
		// WebEditFiling
		//
		// # 编辑网站备案信息
		//
		// 可以通过该接口编辑网站的备案信息，包括备案号、备案名称信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - webICP		网站备案号(string)
		//   - webRecord	网站网安备案号(string)
		//
		// # 返回
		//   - err			错误信息(error)
		WebEditFiling(ctx context.Context, webICP, webRecord string) (err error)
	}
)

var (
	localInfo IInfo
)

func Info() IInfo {
	if localInfo == nil {
		panic("implement not found for interface IInfo, forgot register?")
	}
	return localInfo
}

func RegisterInfo(i IInfo) {
	localInfo = i
}
