package hook

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/service"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyPermissionBeforeProxy 在代理请求之前检查用户是否具有代理操作权限。
//
// 检查逻辑包括:
//   - 验证用户令牌的有效性。
//   - 根据用户 UUID 获取用户信息。
//   - 检查用户的直接权限列表中是否包含代理操作权限。
//   - 如果用户直接权限不存在，则检查用户所属角色的权限列表中是否包含代理操作权限(super_admin可直接跳过)
//
// 参数:
//   - r: 当前请求的上下文对象。
//
// 行为:
//   - 若用户或其角色未授予代理操作权限，则返回无权限响应并中止请求处理。
func ProxyPermissionBeforeProxy(r *ghttp.Request) {
	// 获取当前登录用户信息
	userUUID := r.GetHeader(consts.HeaderUserUUID)
	token := r.GetHeader(consts.HeaderToken)
	token = butil.TokenRemoveBearer(token)

	// 获取令牌
	iToken := service.Token()
	_, errorCode := iToken.GetToken(r.GetCtx(), userUUID, token)
	if errorCode != nil {
		blog.BambooError(r.GetCtx(), "ProxyPermissionBeforeProxy", "获取令牌失败", errorCode)
		r.Response.Status = int(errorCode.Code / 100)
		r.Response.WriteJson(&bmodels.ResponseDTO[interface{}]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
			Data:     &errorCode.Data,
		})
		r.ExitAll()
		return
	}

	// 获取用户
	iUser := service.User()
	getUser, errorCode := iUser.GetUserByUUID(r.GetCtx(), userUUID)
	if errorCode != nil {
		r.Response.Status = int(errorCode.Code / 100)
		r.Response.WriteJson(&bmodels.ResponseDTO[interface{}]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
			Data:     &errorCode.Data,
		})
		r.ExitAll()
		return
	}

	// 检查用户权限是否存在代理操作权限
	permissions := getUser.Permissions.Array()
	for _, permission := range permissions {
		permissionString := g.NewVar(permission).String()
		if permissionString == consts.PermissionProxyOperate {
			return
		}
	}
	// 检查用户所在角色是否存在代理操作权限
	iRole := service.Role()
	getRole, errorCode := iRole.GetRoleByUUID(r.GetCtx(), getUser.Role)
	if errorCode != nil {
		r.Response.Status = int(errorCode.Code / 100)
		r.Response.WriteJson(&bmodels.ResponseDTO[interface{}]{
			Code:     errorCode.Code,
			Message:  errorCode.Message,
			Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
			Time:     gtime.Now().TimestampMilli(),
			Data:     &errorCode.Data,
		})
		r.ExitAll()
		return
	}
	// 如果用户所在角色是超级管理员，则直接返回
	if getRole.RoleName != "super_admin" {
		for _, permission := range getRole.RolePermission.Array() {
			permissionString := g.NewVar(permission).String()
			if permissionString == consts.PermissionProxyOperate {
				return
			}
		}
	} else {
		return
	}

	// 不存在权限不允许访问
	errorCode = &berror.ErrNoPermission
	r.Response.Status = int(errorCode.Code / 100)
	r.Response.WriteJson(&bmodels.ResponseDTO[interface{}]{
		Code:     errorCode.Code,
		Message:  errorCode.Message,
		Overhead: butil.Ptr(gtime.Now().Sub(r.EnterTime).Milliseconds()),
		Time:     gtime.Now().TimestampMilli(),
		Data:     &errorCode.Data,
	})
	r.ExitAll()
}
