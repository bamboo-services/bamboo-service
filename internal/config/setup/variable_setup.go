package setup

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// GetVariableSetup 加载系统变量配置。
//
// 将系统配置表中的关键信息加载到程序的全局变量中，以供后续使用。
func (s *Setup) GetVariableSetup() {
	g.Log().Info(s.ctx, "[INIT] 加载系统变量配置")
	consts.SystemAuthorNameValue = getVariable(s.ctx, consts.SystemAuthorName)
	consts.SystemAuthorEmailValue = getVariable(s.ctx, consts.SystemAuthorEmail)
	consts.SystemAuthorChineseNameValue = getVariable(s.ctx, consts.SystemAuthorChineseName)
	consts.SystemAuthorEnglishNameValue = getVariable(s.ctx, consts.SystemAuthorEnglishName)
	consts.SystemDescriptionValue = getVariable(s.ctx, consts.SystemDescription)
	consts.SystemNameValue = getVariable(s.ctx, consts.SystemName)
	consts.SystemAuthorQQValue = getVariable(s.ctx, consts.SystemAuthorQQ)
	consts.SystemVersionValue = getVariable(s.ctx, consts.SystemVersion)
	consts.SystemAbleRegisterValue = g.NewVar(getVariable(s.ctx, consts.SystemAbleRegister)).Bool()
	consts.SystemSuperAdminUUIDValue = getVariable(s.ctx, consts.SystemSuperAdminUUID)
}

// getVariable 根据指定的键从系统设置表中获取对应的值。
//
// 参数:
//   - ctx: 上下文对象，用于控制请求生命周期。
//   - key: 系统设置表中的键。
//
// 返回:
//   - 对应键的值。如果查询失败会抛出异常。
func getVariable(ctx context.Context, key string) string {
	var system *entity.System
	sqlErr := dao.System.Ctx(ctx).Where(do.System{Key: key}).Scan(&system)
	if sqlErr != nil {
		panic(sqlErr)
	}
	return system.Value
}
