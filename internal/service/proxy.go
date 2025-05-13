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
		ProxyGroupPage(ctx context.Context, userEntity *entity.User, page int, size int, search string) (*[]*dto.ProxyBaseGroupDTO, *berror.ErrorCode)
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
		AddSubscriptionInProxyGroup(ctx context.Context, userEntity *entity.User, proxyGroupUUID string, name string, merchant string, description string, url string) (*dto.ProxyBaseSubscriptionDTO, *berror.ErrorCode)
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
