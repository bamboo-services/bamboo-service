package token

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// GenerateProxyToken 生成新的代理令牌。
//
// 参数:
//   - ctx: 上下文信息，用于请求追踪和控制。
//   - userUUID: 用户的UUID。
//   - name: 代理令牌的名称。
//   - desc: 代理令牌的描述信息。
//   - expiredAt: 代理令牌的过期时间，不能超过当前时间的一年后。
//
// 返回:
//   - *entity.ProxyToken: 包含生成的代理令牌信息。
//   - *berror.ErrorCode: 错误代码，当方法执行失败时返回。
//
// 错误:
//   - berror.ErrInvalidParameters: 参数无效，例如有效期超过一年。
func (s *sToken) GenerateProxyToken(ctx context.Context, userUUID, name, desc string, expiredAt *gtime.Time) (*dto.ProxyTokenDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GenerateProxyToken", "生成代理令牌 %s", name)

	// 生成令牌
	generateToken, errorCode := dao.ProxyToken.CreateNewToken(ctx, userUUID, name, desc, expiredAt)
	if errorCode != nil {
		return nil, errorCode
	}

	// 将生成的令牌转换为DTO
	var newToken *dto.ProxyTokenDTO
	operateErr := gconv.Struct(generateToken, &newToken)
	if operateErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr)
	}

	return newToken, nil
}
