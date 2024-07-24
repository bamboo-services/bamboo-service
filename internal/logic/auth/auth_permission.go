/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package auth

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
)

// IsUserCanDo
//
// # 检查用户是否有权限
//
// 检查用户是否有权限，用于检查用户是否有权限；
// 用于服务层之间进行用户是否有权限操作；
// 一般用作个人业务，个人有权限进行操作，有权限的人也可以进行操作，超级管理员也可以进行操作；
//
// # 参数
//   - ctx				上下文(context.Context)
//   - authorization	用户唯一令牌(string)
//   - userUUID			用户唯一标识(string)
//   - permission		权限(string)
//
// # 返回
//   - error	错误信息
func (s *sAuth) IsUserCanDo(ctx context.Context, authorization, userUUID string, permission string) (err error) {
	g.Log().Notice(ctx, "[SERV] auth.IsUserCanDo | 检查用户是否有权限")
	// 获取用户信息
	getUser, err := s.GetUserByAuthorization(ctx, authorization)
	if err != nil {
		return err
	}
	// 检查用户是否有权限
	if getUser.Uuid != userUUID {
		// 检查是否是管理员
		var getRole *entity.Role
		if err := dao.Role.Ctx(ctx).Where(do.Role{Ruuid: getUser.Ruuid}).Scan(&getRole); err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err, "查询用户角色失败")
		}
		if permission != "" {
			// 如果权限允许该操作
			if getRole.Permission.Contains(permission) {
				return nil
			} else {
				return berror.NewError(bcode.OperationNotPermission, "用户无权限")
			}
		}
		// 若不存在该权限，检查是否是超级管理员
		if s.CheckUserHasSuperAdmin(ctx, authorization) == nil {
			return nil
		}
		return berror.NewError(bcode.OperationNotSupport, "未传递权限")
	} else {
		return nil
	}
}
