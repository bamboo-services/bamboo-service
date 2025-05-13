package proxy

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

// AddSubscriptionInProxyGroup 向代理组中添加订阅地址。
//
// 参数:
//   - ctx: 上下文对象，用于控制生命周期和日志追踪。
//   - userEntity: 用户实体，包含执行操作的用户信息。
//   - proxyGroupUUID: 代理组的唯一标识符。
//   - name: 订阅地址名称。
//   - merchant: 订阅地址的商户信息。
//   - description: 订阅地址的描述信息。
//   - url: 订阅地址的链接。
//
// 返回:
//   - *dto.ProxyBaseSubscriptionDTO: 包含新添加订阅地址信息的数据传输对象。
//   - *berror.ErrorCode: 错误信息对象，操作成功时为 nil。
//
// 错误:
//   - 数据库操作失败时返回 ErrDatabaseError。
//   - 数据转换失败时返回 ErrInternalServer。
func (s *sProxy) AddSubscriptionInProxyGroup(ctx context.Context, userEntity *entity.User, proxyGroupUUID, name, merchant, description, url string) (*dto.ProxyBaseSubscriptionDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "AddSubscriptionInProxyGroup", "由 %s 在代理组 %s 中添加订阅地址 %s", userEntity.Username, proxyGroupUUID, name)
	// 创建订阅地址
	newProxySubscriptionEntity := &entity.ProxySubscription{
		SubscriptionUuid: uuid.New().String(),
		ProxyGroupUuid:   proxyGroupUUID,
		UserUuid:         userEntity.UserUuid,
		Name:             name,
		Merchant:         merchant,
		Description:      description,
		Url:              url,
		CreatedAt:        gtime.Now(),
		UpdatedAt:        gtime.Now(),
	}

	// 插入数据库
	_, sqlErr := dao.ProxySubscription.Ctx(ctx).Insert(&newProxySubscriptionEntity)
	if sqlErr != nil {
		blog.ServiceError(ctx, "AddSubscriptionInProxyGroup", "数据库操作失败: %v", sqlErr)
		return nil, berror.ErrorAddData(&berror.ErrDatabaseError, sqlErr.Error())
	}

	// 转换为DTO
	var newProxySubscriptionDTO *dto.ProxyBaseSubscriptionDTO
	operateErr := gconv.Struct(newProxySubscriptionEntity, &newProxySubscriptionDTO)
	if operateErr != nil {
		blog.ServiceError(ctx, "AddSubscriptionInProxyGroup", "数据转换失败: %v", operateErr)
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr.Error())
	}
	return newProxySubscriptionDTO, nil
}
