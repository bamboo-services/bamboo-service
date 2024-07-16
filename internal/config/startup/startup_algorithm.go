/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package startup

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
)

// createDatabase
//
// # 创建数据库
//
// 创建数据库，根据传入的 schema 名称进行创建数据库；
//
// # 参数
//   - schema		数据库名称(string)
func createDatabase(ctx context.Context, schema string) {
	record, err := g.DB().Ctx(ctx).
		Model("information_schema.tables").
		Where("table_name=?", schema).
		One()
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据表检查失败：%s", err.Error())
	}
	if record.IsEmpty() {
		g.Log().Debugf(ctx, "[STAR] 数据表不存在，创建 %s 数据表", schema)
		// 读取 packed 文件并且根据分号拆分
		gres.Dump()
		getContent := gres.GetContent("resource/schema/" + schema + ".sql")
		if len(getContent) == 0 {
			g.Log().Panicf(ctx, "[STAR] 数据表创建失败：%s 文件不存在", schema)
		}
		_, err := g.DB().Exec(ctx, string(getContent))
		if err != nil {
			g.Log().Panicf(ctx, "[STAR] 数据表创建失败：%s", err.Error())
		}
	}
}

// checkInfoTableValue
//
// # 检查 fy_info 表数据
//
// 检查 fy_info 表数据，检查是否存在数据；若不存在则进行初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
//   - key			键(string)
//   - value		值(string)
func checkInfoTableValue(ctx context.Context, key string, value string) {
	var getInfo *entity.Info
	err := dao.Info.Ctx(ctx).Where(do.Info{Key: key}).Limit(1).Scan(&getInfo)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据表检查失败：%s", err.Error())
	}
	if getInfo == nil {
		g.Log().Debugf(ctx, "[STAR] 数据 %s 不存在，创建 [%s]%s 数据", key, key, value)
		// 读取文件并且根据分号拆分
		_, err := dao.Info.Ctx(ctx).Data(do.Info{SystemUuid: butil.GenerateRandUUID(), Key: key, Value: value}).
			OnConflict(dao.Info.Columns().SystemUuid).
			Save()
		if err != nil {
			g.Log().Panicf(ctx, "[STAR] 数据表创建失败：%s", err.Error())
		}
	}
}

// initializeRole
//
// # 初始化角色
//
// 初始化角色，用于系统初始化时创建角色；
//
// # 请求
//   - ctx			上下文(context.Context)
//   - roleName		角色名称(string)
//   - roleDisplay	角色显示名称(string)
//   - permissions	权限列表([]string)
//   - description	描述(string)
func initializeRole(
	ctx context.Context,
	roleName string,
	roleDisplay string,
	permissions []string,
	description string,
) {
	// 检查角色是否存在
	var getRole *entity.Role
	err := dao.Role.Ctx(ctx).Where(do.Role{Name: roleName}).Scan(&getRole)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据库出现错误：%s", err.Error())
	}
	if getRole != nil {
		return
	}
	// 权限由数组转为 Json
	permissionJSON, err := gjson.Encode(permissions)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 权限转换失败：%s", err.Error())
	}
	_, err = dao.Role.Ctx(ctx).Data(do.Role{
		Ruuid:       butil.GenerateRandUUID(),
		Name:        roleName,
		DisplayName: roleDisplay,
		Description: description,
		Permission:  gjson.New(permissionJSON),
	}).Insert()
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 角色创建失败：%s", err.Error())
	}
}

// hasSuperAdmin
//
// # 检查超级管理员
//
// 检查超级管理员，检查是否存在超级管理员；若不存在则进行初始化操作；
//
// # 请求
//   - ctx			上下文(context.Context)
func hasSuperAdmin(ctx context.Context) bool {
	var getInfo *entity.Info
	err := dao.Info.Ctx(ctx).Where(do.Info{Key: "system_admin"}).Scan(&getInfo)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据库出现错误：%s", err.Error())
	}
	// 检查数据是否存在
	return getInfo != nil
}

// createSuperAdmin
//
// # 创建超级管理员
//
// 创建超级管理员，用于系统初始化时创建超级管理员；
//
// # 请求
//   - ctx			上下文(context.Context)
func createSuperAdmin(ctx context.Context) {
	// 获取管理员角色组
	var getRole *entity.Role
	err := dao.Role.Ctx(ctx).Where(do.Role{Name: "admin"}).Scan(&getRole)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据库出现错误：%s", err.Error())
	}
	if getRole == nil {
		g.Log().Panicf(ctx, "[STAR] 角色不存在")
	}
	// 创建超级管理员
	adminUUID := butil.GenerateRandUUID()
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Uuid:     adminUUID,
		Username: "superAdmin",
		Phone:    "18888888888",
		Password: butil.PasswordEncode("admin"),
		Ruuid:    getRole.Ruuid,
		OtpUuid:  butil.GenerateRandUUID(),
	}).Insert()
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 超级管理员创建失败：%s", err.Error())
	}
	checkInfoTableValue(ctx, "system_admin", adminUUID.String())
}

// getInfoForDB
//
// # 获取数据库信息
//
// 获取数据库信息，用于获取数据库中的信息；用于获取数据库中的信息，存入常量中；
//
// # 请求
//   - ctx			上下文(context.Context)
//   - key			键(string)
func getInfoForDB(ctx context.Context, key string) string {
	var getInfo *entity.Info
	err := dao.Info.Ctx(ctx).Where(do.Info{Key: key}).Scan(&getInfo)
	if err != nil {
		g.Log().Panicf(ctx, "[STAR] 数据库出现错误：%s", err.Error())
	}
	if getInfo == nil {
		g.Log().Panicf(ctx, "[STAR] 数据 %s 不存在", key)
	}
	return getInfo.Value
}
