// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"

	"github.com/XiaoLFeng/bamboo-utils/berror"
)

type (
	IProxy interface {
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
		CreateProxyGroup(ctx context.Context, userEntity *entity.User, name string, description string) (*dto.ProxyBaseGroupDTO, *berror.ErrorCode)
	}
)

var (
	localProxy IProxy
)

func Proxy() IProxy {
	if localProxy == nil {
		panic("implement not found for interface IProxy, forgot register?")
	}
	return localProxy
}

func RegisterProxy(i IProxy) {
	localProxy = i
}
