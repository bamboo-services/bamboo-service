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
	"bamboo-service/api/auth/v1"
	"bamboo-service/internal/model/entity"
	"bamboo-service/internal/service"
	"context"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/bamboo-services/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
)

// AuthLogin
//
// # 用户登录
//
// 用户登录，用于用户登录操作；
// 用户登录需要用户名和密码；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - req		请求(*v1.AuthLoginReq)
//
// # 返回
//   - res		响应(*v1.AuthLoginRes)
//   - err		错误信息(error)
func (c *ControllerV1) AuthLogin(
	ctx context.Context,
	req *v1.AuthLoginReq,
) (res *v1.AuthLoginRes, err error) {
	g.Log().Notice(ctx, "[CONT] 用户登录")
	// 正则表达式判断匹配
	var user *entity.User
	// 判断是否为邮箱
	matchMail := gregex.IsMatchString(
		`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`,
		req.User,
	)
	matchPhone := gregex.IsMatchString(
		`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`,
		req.User,
	)
	switch {
	case matchMail:
		user, err = service.User().GetUserByEmail(ctx, req.User)
	case matchPhone:
		user, err = service.User().GetUserByPhone(ctx, req.User)
	default:
		user, err = service.User().GetUserByUsername(ctx, req.User)
	}
	if err != nil {
		return nil, err
	}
	// 检查数据是否为空
	if user == nil {
		return nil, berror.NewError(bcode.NotExist, "用户不存在")
	}
	// 密码验证
	if butil.PasswordVerify(req.Pass, user.Password) {
		// 生成授权请求
		getToken, err := service.Token().MakeToken(ctx, butil.MakeUUIDByString(req.User), nil)
		if err != nil {
			return nil, err
		}
		return &v1.AuthLoginRes{
			User:  *service.User().UserEntityToUserCurrent(ctx, user),
			Token: getToken.String(),
		}, nil
	} else {
		return nil, berror.NewError(bcode.Unauthorized, "密码错误")
	}
}
