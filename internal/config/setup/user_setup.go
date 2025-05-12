package setup

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

// CheckUserSuperAdminExist 检查是否存在超级管理员用户。
//
// 如果不存在超级管理员用户，则创建一个默认的超级管理员，并更新系统的超级管理员UUID值。
func (s *Setup) CheckUserSuperAdminExist() {
	g.Log().Info(s.ctx, "[INIT] 检查超级管理员")
	var system *entity.System
	err := dao.System.Ctx(s.ctx).Where(do.System{Key: consts.SystemSuperAdminUUID}).Scan(&system)
	if err != nil {
		panic(err)
	}
	if system.Value == "" {
		newUserUUID := uuid.New().String()
		createUser(s.ctx, newUserUUID)
		// 修改系统超级管理员UUID
		_, err = dao.System.Ctx(s.ctx).Where(do.System{Key: consts.SystemSuperAdminUUID}).Data(do.System{Value: newUserUUID}).Update()
		if err != nil {
			panic(err)
		}
		g.Log().Info(s.ctx, "初始超级管理员创建成功「请尽快登录并修改密码」")
		g.Log().Info(s.ctx, "\t账号: admin")
		g.Log().Info(s.ctx, "\t密码: admin")
	} else {
		// 检查是否已存在超级管理员
		var user *entity.User
		err := dao.User.Ctx(s.ctx).Where(do.User{UserUuid: system.Value}).Scan(&user)
		if err != nil {
			panic(err)
		}
		if user == nil {
			createUser(s.ctx, system.Value)
		} else {
			g.Log().Debug(s.ctx, "超级管理员存在，忽略创建")
		}
	}
}

func createUser(ctx context.Context, userUUID string) {
	// 获取角色列表
	var role *entity.Role
	err := dao.Role.Ctx(ctx).Where(do.Role{RoleName: consts.RoleSuperAdmin}).Scan(&role)
	if err != nil {
		panic(err)
	}
	// 自动创建超级管理员
	newUser := &do.User{
		UserUuid:     userUUID,
		Username:     "admin",
		Email:        "admin@x-lf.cn",
		Role:         role.RoleUuid,
		PasswordHash: butil.PasswordEncode("123456"),
		Nickname:     "SuperAdmin",
	}
	_, err = dao.User.Ctx(ctx).Insert(newUser)
	if err != nil {
		panic(err)
	}
}
