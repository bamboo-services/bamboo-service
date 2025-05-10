package auth

import (
	v1 "bamboo-service/api/auth/v1"
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
)

// CheckUserExistByUsername 检查指定用户名的用户是否存在。
//
// 当用户
//
// 参数:
//   - ctx: 请求上下文信息，用于控制操作的生命周期。
//   - Username: 需要检查的用户名。
//
// 返回:
//   - err: 错误代码，表示用户不存在或其他错误情况。
func (s *sAuth) CheckUserExistByUsername(ctx context.Context, username string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByUsername", "检查用户名 %s 是否存在", username)
	getUser, errorCode := dao.GetUserByUsername(ctx, username)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}

// CheckUserExistByEmail 检查指定邮箱的用户是否存在。
//
// 参数:
//   - ctx: 请求上下文，用于控制操作生命周期。
//   - Email: 需要检查的邮箱地址。
//
// 返回:
//   - 错误代码，表示用户是否存在或其他错误情况。
func (s *sAuth) CheckUserExistByEmail(ctx context.Context, email string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByEmail", "检查邮箱 %s 是否存在", email)
	getUser, errorCode := dao.GetUserByEmail(ctx, email)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}

// CheckUserExistByPhone 检查指定手机号的用户是否存在。
//
// 参数:
//   - ctx: 请求上下文，用于控制操作生命周期。
//   - phone: 需要检查的手机号。
//
// 返回:
//   - 错误代码，表示用户是否存在或其他错误情况。
func (s *sAuth) CheckUserExistByPhone(ctx context.Context, phone string) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "CheckUserExistByPhone", "检查手机号 %s 是否存在", phone)
	getUser, errorCode := dao.GetUserByPhone(ctx, phone)
	if errorCode != nil {
		return errorCode
	}
	if getUser != nil {
		return custom.ErrorUserExist
	}
	return nil
}

// UserLogin 验证用户登录，并返回用户信息。
//
// 参数:
//   - ctx: 请求上下文。
//   - Username: 用户名。
//   - Password: 用户密码。
//
// 返回:
//   - userInfo: 用户信息数据传输对象。
//   - err: 错误代码，表示登录失败的原因。
func (s *sAuth) UserLogin(ctx context.Context, username, password string) (*dto.UserInfoDTO, *berror.ErrorCode) {
	blog.ControllerInfo(ctx, "UserLogin", "用户登录")
	getUser, errorCode := dao.GetUserByUsername(ctx, username)
	if errorCode != nil {
		return nil, errorCode
	}
	// 检查用户密码
	if butil.PasswordVerify(password, (getUser).PasswordHash) {
		userInfo := &dto.UserInfoDTO{}
		operateErr := gconv.Struct(getUser, userInfo)
		if operateErr != nil {
			return nil, berror.ErrorAddData(berror.ErrInternalServer, operateErr)
		}
		return userInfo, nil
	} else {
		return nil, custom.ErrorUserPasswordIncorrect
	}
}

// UserRegister 注册新用户。
//
// 参数:
//   - ctx: 请求上下文，用于控制操作生命周期。
//   - request: 包含用户名和密码的用户注册请求。
//
// 返回:
//   - 错误代码，表示注册失败的原因或 nil 表示成功。
func (s *sAuth) UserRegister(ctx context.Context, request *v1.AuthRegisterReq) (*entity.User, *berror.ErrorCode) {
	blog.ServiceInfo(ctx, "UserRegister", "用户注册 %s", request.Username)
	// 获取普通用户 UUID
	var getRole *entity.Role
	sqlErr := dao.Role.Ctx(ctx).Where(&entity.Role{RoleName: consts.RoleUser}).Scan(&getRole)
	if sqlErr != nil {
		return nil, berror.ErrorAddData(berror.ErrDatabaseError, sqlErr)
	}
	var newUserUUID = uuid.New().String()
	newUser := &entity.User{
		UserUuid:       newUserUUID,
		Username:       request.Username,
		Email:          request.Email,
		Phone:          request.Phone,
		Role:           getRole.RoleUuid,
		PasswordHash:   butil.PasswordEncode(request.Password),
		RegistrationIp: ghttp.RequestFromCtx(ctx).GetClientIp(),
	}
	_, sqlErr = dao.User.Ctx(ctx).Insert(newUser)
	if sqlErr != nil {
		return nil, berror.ErrorAddData(berror.ErrDatabaseError, sqlErr)
	}
	// 取得新用户信息
	user, errorCode := dao.GetUserByUUID(ctx, newUserUUID)
	if errorCode != nil {
		return nil, errorCode
	}
	return user, nil
}
