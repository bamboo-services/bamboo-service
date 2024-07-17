/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package user

import (
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// UserEntityToUserCurrent
//
// # 用户实体转用户当前
//
// 用户实体转用户当前，将用户实体转换为用户当前数据传输对象；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - user		用户实体(*entity.User)
//
// # 返回
//   - *dto.UserCurrentDTO	用户当前数据传输对象(*dto.UserCurrentDTO)
func (s *sUser) UserEntityToUserCurrent(ctx context.Context, user *entity.User) (userCurrent *dto.UserCurrentDTO) {
	g.Log().Notice(ctx, "[SERV] user.UserEntityToUserCurrent | 用户实体转用户当前")
	recordRole, err := dao.Role.Ctx(ctx).Where(do.Role{Ruuid: user.Ruuid}).One()
	if err != nil {
		g.Log().Warningf(ctx, "[SERV] 查询用户角色失败：%s", err.Error())
	}
	recordVip, err := dao.Vip.Ctx(ctx).Where(do.Vip{Vuuid: user.Vuuid}).One()
	if err != nil {
		g.Log().Warningf(ctx, "[SERV] 查询用户VIP失败：%s", err.Error())
	}
	var ruuidString, vuuidString string
	if recordRole != nil {
		ruuidString = recordRole.GMap().Get("name").(string)
	} else {
		ruuidString = ""
	}
	if recordVip != nil {
		vuuidString = recordVip.GMap().Get("name").(string)
	} else {
		vuuidString = ""
	}
	return &dto.UserCurrentDTO{
		UUID:        user.Uuid,
		Username:    user.Username,
		Phone:       user.Phone,
		Email:       user.Email,
		Ruuid:       ruuidString,
		Vuuid:       vuuidString,
		Description: user.Description,
	}
}
