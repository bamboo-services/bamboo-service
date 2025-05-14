package proxy

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/model/entity"
	"bamboo-service/internal/service"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bresult"
	"github.com/gogf/gf/v2/net/ghttp"

	"bamboo-service/api/proxy/v1"
)

// ProxyGroupList 获取代理组列表。
//
// 此函数用于获取代理组列表信息。它首先验证用户身份和权限，然后根据用户角色
// 和请求参数返回相应的代理组列表数据。超级管理员可以查看所有用户的代理组，
// 普通用户只能查看自己的代理组。
//
// 参数:
//   - ctx: 上下文对象，用于传递操作范围和数据。
//   - req: 请求对象，包含筛选条件。
//
// 返回:
//   - res: 包含代理组列表数据的响应对象。
//   - err: 可能的错误结果。
//
// 错误:
//   - 若用户信息获取失败或用户不存在，将返回对应错误。
//   - 若代理组列表查询失败，将返回对应错误。
func (c *ControllerV1) ProxyGroupList(ctx context.Context, req *v1.ProxyGroupListReq) (res *v1.ProxyGroupListRes, err error) {
	blog.ControllerInfo(ctx, "ProxyGroupList", "获取代理组列表")

	// 获取用户信息
	request := ghttp.RequestFromCtx(ctx)
	iUser := service.User()
	getUser, errorCode := iUser.GetUserByUUID(ctx, request.GetHeader(consts.HeaderUserUUID))
	if errorCode != nil {
		return nil, errorCode
	}
	if getUser == nil {
		return nil, custom.ErrorUserNotExist
	}

	// 获取用户所在角色
	iRole := service.Role()
	getRole, errorCode := iRole.GetRoleByUUID(ctx, getUser.Role)
	if errorCode != nil {
		return nil, errorCode
	}
	if getRole == nil {
		return nil, custom.ErrorRoleNotExist
	}

	// 检查角色是否是超级管理员
	var userEntity *entity.User
	if getRole.RoleName != consts.RoleSuperAdmin {
		if req.UserUUID != "" {
			return nil, berror.ErrorAddData(&berror.ErrNoPermission, "没有权限查阅他人信息")
		}
		userEntity = getUser
	} else {
		if req.UserUUID != "" {
			userEntity, errorCode = iUser.GetUserByUUID(ctx, req.UserUUID)
		} else {
			userEntity = getUser
		}
	}

	// 获取代理组列表
	iProxy := service.Proxy()
	proxyGroupList, errorCode := iProxy.ProxyGroupList(ctx, userEntity, req.Search)
	if errorCode != nil {
		return nil, errorCode
	}
	return &v1.ProxyGroupListRes{
		ResponseDTO: bresult.SuccessHasData(ctx, "获取代理组列表成功", proxyGroupList),
	}, nil
}
