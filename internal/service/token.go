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

	"github.com/google/uuid"
)

type (
	IToken interface {
		// MakeToken
		//
		// # 生成Token
		//
		// 生成Token，用于生成Token操作；
		// 生成Token需要用户UUID，其中 halfDay 为有效天数（1 为半天），不填写默认有效期为 1 天（范围为正整数 int8 范围，不含 0）
		// 生成Token操作失败会返回错误信息；
		// 生成的 Token 会返回所生成 Token 的 uuid.UUID 信息，用于后续操作，并且将会自动存入缓存中；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - userUUID		用户UUID(uuid.UUID)
		//   - halfDay		有效天数(*uint8)
		//
		// # 返回
		//   - err		错误信息(error)
		MakeToken(ctx context.Context, userUUID uuid.UUID, halfDay *uint8) (token *uuid.UUID, err error)
		// VerifyToken
		//
		// # 验证Token
		//
		// 验证Token，用于验证Token操作；
		// 验证Token需要用户UUID和Token，其中 user 为用户UUID，token 为 Token；
		// 验证Token操作失败会返回错误信息；
		// 验证 Token 会返回错误信息，如果验证成功则返回 nil；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - user			用户UUID(uuid.UUID)
		//   - token		Token(uuid.UUID)
		//
		// # 返回
		//   - err		错误信息(error)
		VerifyToken(ctx context.Context, user, token uuid.UUID) (err error)
		// RemoveToken
		//
		// # 移除Token
		//
		// 移除Token，用于移除Token操作；
		// 移除Token需要 Token，其中 token 为 Token；
		// 移除Token操作失败会返回错误信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - token		Token(uuid.UUID)
		//
		// # 返回
		//   - err		错误信息(error)
		RemoveToken(ctx context.Context, token uuid.UUID) (err error)
	}
)

var (
	localToken IToken
)

func Token() IToken {
	if localToken == nil {
		panic("implement not found for interface IToken, forgot register?")
	}
	return localToken
}

func RegisterToken(i IToken) {
	localToken = i
}
