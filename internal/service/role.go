// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/entity"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IRole interface {
		// GetRoleByUUID 根据角色 UUID 获取角色信息。
		//
		// 参数:
		//   - ctx: 上下文对象，用于控制请求的生命周期。
		//   - roleUUID: 角色的唯一标识符。
		//
		// 返回:
		//   - *entity.Role: 匹配的角色信息，若未找到则为 nil。
		//   - *berror.ErrorCode: 错误代码，若没有错误则为 nil。
		GetRoleByUUID(ctx context.Context, roleUUID string) (*entity.Role, *berror.ErrorCode)
		// GetRoleByName 根据角色名称获取角色信息。
		//
		// 参数:
		//   - ctx: 上下文对象，用于控制请求的生命周期。
		//   - roleName: 角色的名称。
		//
		// 返回:
		//   - *entity.Role: 匹配的角色信息，若未找到则为 nil。
		//   - *berror.ErrorCode: 错误代码，若没有错误则为 nil。
		GetRoleByName(ctx context.Context, roleName string) (*entity.Role, *berror.ErrorCode)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
