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

// CreateProxyGroup 创建一个新的代理组。
//
// 参数:
//   - ctx: 上下文对象，用于控制生命周期和日志追踪。
//   - userEntity: 用户实体，包含执行操作的用户信息。
//   - name: 代理组的名称。
//   - description: 代理组的描述。
//
// 返回:
//   - *dto.ProxyBaseGroupDTO: 包含新创建代理组信息的数据传输对象。
//   - *berror.ErrorCode: 错误信息对象，操作成功时为 nil。
//
// 错误:
//   - 数据库插入失败时返回 ErrDatabaseError。
//   - 数据转换失败时返回 ErrInternalServer。
func (s *sProxy) CreateProxyGroup(ctx context.Context, userEntity *entity.User, name, description string) (*dto.ProxyBaseGroupDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "CreateProxyGroup", "由 %s 创建代理组 %s", userEntity.Username, name)
	// 创建代理组
	newProxyGroupEntity := &entity.ProxyGroup{
		GroupUuid:   uuid.New().String(),
		UserUuid:    userEntity.UserUuid,
		Name:        name,
		Description: description,
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	}

	// 插入数据库
	_, sqlErr := dao.ProxyGroup.Ctx(ctx).Insert(&newProxyGroupEntity)
	if sqlErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrDatabaseError, sqlErr)
	}

	// 转换为DTO
	var newProxyGroupDTO *dto.ProxyBaseGroupDTO
	operateErr := gconv.Struct(newProxyGroupEntity, &newProxyGroupDTO)
	if operateErr != nil {
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr)
	}
	return newProxyGroupDTO, nil
}
