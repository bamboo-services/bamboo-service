package mail

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// GetMailCode 根据邮箱地址和用途生成并返回邮件验证码。
//
// 功能概述:
// 检查邮件验证码的发送频率，避免频繁发送；若符合条件，则生成新的验证码并缓存相关信息。
//
// 参数:
//   - ctx: 上下文，用于控制生命周期和传递信息。
//   - email: 邮箱地址，接收验证码的目标地址。
//   - purpose: 验证码的用途，如注册、密码重置等。
//
// 返回:
//   - *dto.MailCodeDTO: 封装邮箱、验证码及创建时间的传输数据对象。
//   - *berror.ErrorCode: 错误信息，包含缓存、频率限制或验证码生成失败的原因。
func (s *sMail) GetMailCode(ctx context.Context, email string, purpose string) (*dto.MailCodeDTO, *berror.ErrorCode) {
	// 检查最后一个发送的邮件是否少于 1 分钟
	getLastSendTime, redisErr := g.Redis().GetEX(ctx, fmt.Sprintf(consts.RedisMailCodeSendTime, gmd5.MustEncrypt(email)))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
	}

	// 重复发送的事件在一分钟内「禁止重复发送」
	if !getLastSendTime.IsNil() && !getLastSendTime.IsEmpty() {
		sendTime := getLastSendTime.GTime()
		if gtime.Now().Sub(sendTime).Milliseconds() < 60000 {
			var errorBack = make(map[string]interface{})
			errorBack["last_send_time"] = sendTime.String()
			errorBack["resend_milliseconds"] = 60000 - gtime.Now().Sub(sendTime).Milliseconds()
			return nil, berror.ErrorAddData(*custom.ErrMailCodeSentTooFrequently, errorBack)
		}
	}

	// 创建新的验证码
	mailCodeEntity, errorCode := dao.EmailCode.CreateMailCode(ctx, email, purpose)
	if errorCode != nil {
		return nil, errorCode
	}
	// 缓存创建最后创建的验证码时间
	redisErr = g.Redis().SetEX(ctx, fmt.Sprintf(consts.RedisMailCodeSendTime, gmd5.MustEncrypt(email)), gtime.Now(), int64(30*time.Minute))
	if redisErr != nil {
		return nil, berror.ErrorAddData(berror.ErrCacheError, redisErr)
	}
	return &dto.MailCodeDTO{
		Email:     mailCodeEntity.Email,
		CreatedAt: mailCodeEntity.CreatedAt,
	}, nil
}
