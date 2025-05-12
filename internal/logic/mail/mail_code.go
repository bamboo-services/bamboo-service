package mail

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

// GenerateMailCode 生成邮件验证码并限制频繁发送。
//
// 参数:
//   - ctx: 上下文，用于控制生命周期和日志记录。
//   - email: 邮箱地址，用于接收验证码的用户邮箱。
//   - purpose: 验证码的用途描述。
//
// 返回:
//   - *dto.MailCodeDTO: 包含生成的验证码及其相关信息。
//   - *berror.ErrorCode: 错误信息，如发送频率限制或缓存操作失败。
func (s *sMail) GenerateMailCode(ctx context.Context, email string, purpose string) (*dto.MailCodeDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GenerateMailCode", "生成 %s 的验证码", email)
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
		Code:      mailCodeEntity.Code,
		CreatedAt: mailCodeEntity.CreatedAt,
	}, nil
}
