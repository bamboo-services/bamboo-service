package proxy

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/gogf/gf/v2/encoding/gjson"
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
	blog.ServiceInfo(ctx, "CreateProxyGroup", "由 %s 创建代理组 [%s]", userEntity.Username, name)
	// 创建代理组
	newProxyGroupEntity := &entity.ProxyGroup{
		GroupUuid:   uuid.New().String(),
		UserUuid:    userEntity.UserUuid,
		Name:        name,
		Description: description,
		Proxy:       gjson.New([]string{}),
		Partition:   gjson.New([]string{}),
		Rule:        gjson.New([]string{}),
		CreatedAt:   gtime.Now(),
		UpdatedAt:   gtime.Now(),
	}

	// 插入数据库
	_, sqlErr := dao.ProxyGroup.Ctx(ctx).Insert(&newProxyGroupEntity)
	if sqlErr != nil {
		blog.ServiceError(ctx, "CreateProxyGroup", "数据库插入失败: %v", sqlErr)
		return nil, berror.ErrorAddData(&berror.ErrDatabaseError, sqlErr.Error())
	}

	// 转换为DTO
	var newProxyGroupDTO *dto.ProxyBaseGroupDTO
	operateErr := gconv.Struct(newProxyGroupEntity, &newProxyGroupDTO)
	if operateErr != nil {
		blog.ServiceError(ctx, "CreateProxyGroup", "数据转换失败: %v", operateErr)
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr.Error())
	}
	return newProxyGroupDTO, nil
}

// ProxyGroupPage 分页检索用户的代理组信息。
//
// 参数:
//   - ctx: 上下文对象，用于控制生命周期和日志追踪。
//   - userEntity: 用户实体，包含执行操作的用户信息。
//   - page: 页码，从 1 开始。
//   - size: 每页返回的记录数。
//   - search: 搜索关键词，用于匹配代理组名称或描述。
//
// 返回:
//   - []*dto.ProxyBaseGroupDTO: 包含代理组信息列表的数据传输对象数组。
//   - *berror.ErrorCode: 错误信息对象，操作成功时为 nil。
func (s *sProxy) ProxyGroupPage(ctx context.Context, userEntity *entity.User, page, size int, search string) (*[]*dto.ProxyBaseGroupDTO, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "ProxyGroupPage", "由 %s 分页检索代理组信息", userEntity.Username)
	// 查询代理组列表
	var proxyGroupList []*entity.ProxyGroup
	model := dao.ProxyGroup.Ctx(ctx).Where(&do.ProxyGroup{UserUuid: userEntity.UserUuid})
	if search != "" {
		model.WhereLike("name", fmt.Sprintf("%%%s%%", search))
		model.WhereOrLike("description", fmt.Sprintf("%%%s%%", search))
		model.WhereOrLike("file_name", fmt.Sprintf("%%%s%%", search))
	}
	sqlErr := model.Page(page, size).Scan(&proxyGroupList)
	if sqlErr != nil {
		blog.ServiceError(ctx, "ProxyGroupPage", "数据库查询失败: %v", sqlErr)
		return nil, berror.ErrorAddData(&berror.ErrDatabaseError, sqlErr.Error())
	}

	// 转换为DTO
	var proxyGroupListDTO []*dto.ProxyBaseGroupDTO
	operateErr := gconv.Structs(proxyGroupList, &proxyGroupListDTO)
	if operateErr != nil {
		blog.ServiceError(ctx, "ProxyGroupPage", "数据转换失败: %v", operateErr)
		return nil, berror.ErrorAddData(&berror.ErrInternalServer, operateErr.Error())
	}
	return &proxyGroupListDTO, nil
}
