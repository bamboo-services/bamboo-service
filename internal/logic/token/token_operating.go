/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package token

import (
	"bamboo-service/utility"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"strings"
)

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
func (s *sToken) MakeToken(ctx context.Context, userUUID uuid.UUID, halfDay *uint8) (token *uuid.UUID, err error) {
	g.Log().Notice(ctx, "[SERV] token.MakeToken | 生成令牌接口")
	// 生成 Token
	getToken := butil.GenerateRandUUID()
	// 检查有效期
	if halfDay == nil {
		halfDay = new(uint8)
		*halfDay = 2
	}
	// 存入缓存
	_, err = g.Redis().Set(ctx, "token:"+utility.UUIDReplaceDash(getToken), userUUID.String())
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "存入缓存失败")
	}
	_, err = g.Redis().Expire(ctx, "token:"+utility.UUIDReplaceDash(getToken), gconv.Int64(halfDay)*43200)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "设置有效期失败")
	}
	return &getToken, nil
}

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
func (s *sToken) VerifyToken(ctx context.Context, user, token uuid.UUID) (err error) {
	g.Log().Notice(ctx, "[SERV] token.VerifyToken | 验证令牌接口")
	// 读取令牌
	getUser, err := g.Redis().Get(ctx, "token:"+utility.UUIDReplaceDash(token))
	// 检查用户令牌是否存在
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "获取令牌失败")
	}
	if getUser.IsNil() {
		return berror.NewError(bcode.Unauthorized, "令牌不存在")
	}
	// 匹配成功
	if strings.EqualFold(user.String(), getUser.String()) {
		return nil
	} else {
		return berror.NewError(bcode.Unauthorized, "令牌验证失败")
	}
}

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
func (s *sToken) RemoveToken(ctx context.Context, token uuid.UUID) (err error) {
	g.Log().Notice(ctx, "[SERV] token.RemoveToken | 移除令牌接口")
	// 移除令牌
	_, err = g.Redis().Del(ctx, "token:"+utility.UUIDReplaceDash(token))
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "移除令牌失败")
	}
	return nil
}
