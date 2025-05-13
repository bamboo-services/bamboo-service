package role

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
)

// GetRoleByUUID 根据角色 UUID 获取角色信息。
//
// 参数:
//   - ctx: 上下文对象，用于控制请求的生命周期。
//   - roleUUID: 角色的唯一标识符。
//
// 返回:
//   - *entity.Role: 匹配的角色信息，若未找到则为 nil。
//   - *berror.ErrorCode: 错误代码，若没有错误则为 nil。
func (s *sRole) GetRoleByUUID(ctx context.Context, roleUUID string) (*entity.Role, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRoleByUUID", "根据UUID %s 获取角色信息", roleUUID)
	getRole, errorCode := dao.Role.GetRoleByUUID(ctx, roleUUID)
	if errorCode != nil {
		return nil, errorCode
	}
	return getRole, nil
}

// GetRoleByName 根据角色名称获取角色信息。
//
// 参数:
//   - ctx: 上下文对象，用于控制请求的生命周期。
//   - roleName: 角色的名称。
//
// 返回:
//   - *entity.Role: 匹配的角色信息，若未找到则为 nil。
//   - *berror.ErrorCode: 错误代码，若没有错误则为 nil。
func (s *sRole) GetRoleByName(ctx context.Context, roleName string) (*entity.Role, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "GetRoleByName", "根据名称 %s 获取角色信息", roleName)
	getRole, errorCode := dao.Role.GetRoleByName(ctx, roleName)
	if errorCode != nil {
		return nil, errorCode
	}
	return getRole, nil
}
