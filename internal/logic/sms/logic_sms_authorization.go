/*
 * ***********************************************************
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ***********************************************************
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ***********************************************************
 */

package sms

import (
	"XiaoService/internal/model/rdo"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// CheckIfAuthorizationIsAvailable
//
// # 检查授权是否可用
//
// 检查授权是否可用，用于检查授权是否可用，根据用户的机器 IP 以及 UserAgent 生成授权码；
// 若返回的内容没有报错的内容则说明可以继续获取授权码；
// 在对授权码的可用性检查过程中，会操作 Redis 缓存，并不是单纯判断进行检查是否可用。
//
// # 参数
//   - ctx		上下文(context.Context)
//
// # 返回
//   - err		错误(error)
func (s *sSms) CheckIfAuthorizationIsAvailable(ctx context.Context, authorizationCode string) (err error) {
	g.Log().Noticef(ctx, "[LOGIC] 检查授权是否可用 [CheckIfAuthorizationIsAvailable]")
	getRequest := g.RequestFromCtx(ctx)
	// 获取基本信息
	ip := getRequest.GetClientIp()
	userAgent := getRequest.GetHeader("User-Agent")
	// 生成固定授权码
	makeCacheUUID := butil.MakeUUIDByString(ip + userAgent)
	var redisSmsAuthorization *rdo.RedisSmsAuthorization
	getMap, err := g.Redis().HGetAll(ctx, "sms:auth:"+makeCacheUUID.String())
	err = getMap.Structs(&redisSmsAuthorization)
	if err == nil {
		if redisSmsAuthorization.LastSendAt == nil {
			// 生成授权码
			sendingUUID, _ := uuid.NewV7()
			_ = g.Redis().HMSet(ctx, "sms:auth:"+makeCacheUUID.String(), gconv.Map(rdo.RedisSmsAuthorization{
				LastSendAt:  gtime.Now(),
				Frequency:   1,
				SendingUUID: sendingUUID.String(),
			}))
			_, _ = g.Redis().Expire(ctx, "sms:auth:"+makeCacheUUID.String(), 900)
		} else {
			// 获取最后发送时间
			timeLeft := gtime.Now().Timestamp() - redisSmsAuthorization.LastSendAt.Timestamp()
			if timeLeft > 60 {
				// 并且重复发送小于 5 次
				if redisSmsAuthorization.Frequency < 5 {
					// 更新发送
					sendingUUID, _ := uuid.NewV7()
					_, _ = g.Redis().HSetNX(
						ctx,
						"sms:auth:"+makeCacheUUID.String(),
						"sending_uuid",
						sendingUUID.String(),
					)
					_, _ = g.Redis().HIncrBy(ctx, "sms:auth:"+makeCacheUUID.String(), "frequency", 1)
				} else {
					_, _ = g.Redis().ExpireAt(ctx, "sms:auth:"+makeCacheUUID.String(), gtime.Now().Add(time.Second*900).Time)
					err = berror.NewError(bcode.OperationFailed, "发送次数过多，请稍后再试")
				}
			} else {
				if authorizationCode != redisSmsAuthorization.SendingUUID {
					// 重新设置有效期
					err = berror.NewError(bcode.OperationFailed,
						"发送时间间隔过短，请稍后再试",
						"剩余 "+strconv.FormatInt(60-timeLeft, 10)+" 秒",
					)
				}
			}
		}
	} else {
		err = berror.NewErrorHasError(bcode.ServerInternalError, err, "获取授权码失败")
	}
	// 返回错误
	return err
}

// GetSmsAuthorization
//
// # 获取短信验证码授权码
//
// 获取短信验证码授权码，用于获取短信验证码授权码；
//
// # 参数
//   - ctx		上下文(context.Context)
//
// # 返回
func (s *sSms) GetSmsAuthorization(ctx context.Context) (authorizationCode *rdo.RedisSmsAuthorization, err error) {
	g.Log().Noticef(ctx, "[LOGIC] 获取短信验证码授权码 [GetSmsAuthorization]")
	getRequest := g.RequestFromCtx(ctx)
	// 获取基本信息
	ip := getRequest.GetClientIp()
	userAgent := getRequest.GetHeader("User-Agent")
	// 生成固定授权码
	makeCacheUUID := butil.MakeUUIDByString(ip + userAgent)
	var redisSmsAuthorization *rdo.RedisSmsAuthorization
	getMap, err := g.Redis().HGetAll(ctx, "sms:auth:"+makeCacheUUID.String())
	err = getMap.Structs(&redisSmsAuthorization)
	if err == nil {
		if redisSmsAuthorization == nil {
			err = berror.NewError(bcode.OperationFailed, "授权码不存在")
		}
	} else {
		err = berror.NewErrorHasError(bcode.ServerInternalError, err, "获取授权码失败")
	}
	// 返回错误
	return redisSmsAuthorization, err
}
