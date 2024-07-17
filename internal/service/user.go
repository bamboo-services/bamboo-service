/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"bamboo-service/internal/model/dto"
	"bamboo-service/internal/model/entity"
	"context"

	"github.com/google/uuid"
)

type (
	IUser interface {
		// UserChangePassword
		//
		// # 用户修改密码
		//
		// 用户修改密码，用于用户修改密码，需要传递用户的 UUID 和新密码；
		// 该接口将会对用户的密码进行修改，修改成功后将会返回 nil 信息；
		// 新的密码将会重新加密存入数据库中，旧密码将会保存在数据库旧密码位置；
		// 修改密码将会检查不允许修改当前密码以及上一次密码，以及重置密码阶段可以对上一次的密码进行密码找回操作；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - getUUID		用户 UUID(uuid.UUID)
		//   - newPassword	新密码(string)
		//
		// # 返回
		//   - err			错误信息(error)
		UserChangePassword(ctx context.Context, getUUID uuid.UUID, newPassword string) (err error)
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
		UserEntityToUserCurrent(ctx context.Context, user *entity.User) (userCurrent *dto.UserCurrentDTO)
		// UserExistByUsername
		//
		// # 用户是否存在
		//
		// 用户是否存在，用于检查用户是否存在；检查用户名是否重复，如果重复则返回错误；
		// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
		// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - username		用户名(string)
		//
		// # 返回
		//   - getUser		用户信息(*entity.User)
		//   - err			错误信息(error)
		UserExistByUsername(ctx context.Context, username string) (getUser *entity.User, err error)
		// UserExistByPhone
		//
		// # 用户是否存在
		//
		// 用户是否存在，用于检查用户是否存在；检查手机号是否重复，如果重复则返回错误；
		// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
		// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		手机号(string)
		//
		// # 返回
		//   - getUser		用户信息(*entity.User)
		//   - err			错误信息(error)
		UserExistByPhone(ctx context.Context, phone string) (getUser *entity.User, err error)
		// UserExistByEmail
		//
		// # 用户是否存在
		//
		// 用户是否存在，用于检查用户是否存在；检查邮箱是否重复，如果重复则返回错误；
		// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
		// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - email		邮箱(string)
		//
		// # 返回
		//   - getUser		用户信息(*entity.User)
		//   - err			错误信息(error)
		UserExistByEmail(ctx context.Context, email string) (getUser *entity.User, err error)
		// UserExistByUUID
		//
		// # 用户是否存在
		//
		// 用户是否存在，用于检查用户是否存在；检查 UUID 是否存在，如果存在则返回错误；
		// 若用户不存在将返回 nil 信息；若有其他错误信息返回错误信息；
		// 若用户存在返回 bcode.AlreadyExists 的 error 信息，以及 *entity.User 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - getUUID		用户 UUID(string)
		//
		// # 返回
		//   - getUser		用户信息(*entity.User)
		//   - err			错误信息(error)
		UserExistByUUID(ctx context.Context, getUUID uuid.UUID) (getUser *entity.User, err error)
		// GetUserByUUID
		//
		// # 获取用户信息
		//
		// 获取用户信息，用于获取用户信息，需要传递用户的 UUID；
		// 该接口将会对用户的 UUID 进行查询，查询成功后将会返回 *entity.User 信息；
		// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
		// 若有其他错误信息返回错误信息；
		// 优先从缓存读取数据信息，若缓存中不存在则从数据库中读取，并且读取的数据将会存入缓存中，缓存有效期 6 小时；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - userUUID		用户 UUID(uuid.UUID)
		//
		// # 返回
		//   - user			用户信息(*entity.User)
		//   - err			错误信息(error)
		GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (user *entity.User, err error)
		// GetUserByEmail
		//
		// # 获取用户信息
		//
		// 获取用户信息，用于获取用户信息，需要传递用户的 Email；
		// 该接口将会对用户的 Email 进行查询，查询成功后将会返回 *entity.User 信息；
		// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
		// 若有其他错误信息返回错误信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - email		用户 Email(string)
		//
		// # 返回
		//   - user			用户信息(*entity.User)
		//   - err			错误信息(error)
		GetUserByEmail(ctx context.Context, email string) (user *entity.User, err error)
		// GetUserByPhone
		//
		// # 获取用户信息
		//
		// 获取用户信息，用于获取用户信息，需要传递用户的 Phone；
		// 该接口将会对用户的 Phone 进行查询，查询成功后将会返回 *entity.User 信息；
		// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
		// 若有其他错误信息返回错误信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - phone		用户 Phone(string)
		//
		// # 返回
		//   - user			用户信息(*entity.User)
		//   - err			错误信息(error)
		GetUserByPhone(ctx context.Context, phone string) (user *entity.User, err error)
		// GetUserByUsername
		//
		// # 获取用户信息
		//
		// 获取用户信息，用于获取用户信息，需要传递用户的 Username；
		// 该接口将会对用户的 Username 进行查询，查询成功后将会返回 *entity.User 信息；
		// 若用户不存在将会返回 bcode.NotExist 的 error 信息；
		// 若有其他错误信息返回错误信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - username		用户 Username(string)
		//
		// # 返回
		//   - user			用户信息(*entity.User)
		//   - err			错误信息(error)
		GetUserByUsername(ctx context.Context, username string) (user *entity.User, err error)
		// GetUserIntoRedis
		//
		// # 获取用户信息存入缓存
		//
		// 获取用户信息存入缓存，用于获取用户信息，需要传递用户 uuid.UUID 或 entity.User；
		// uuid.UUID 和 entity.User 务必选择其一写入，否则将产生错误；
		// 执行成功后将不返回结果，若产生错误将返回错误信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - getUUID		用户 UUID(*uuid.UUID)
		//   - user			用户信息(*entity.User)
		//
		// # 返回
		//   - err			错误信息(error)
		GetUserIntoRedis(ctx context.Context, getUUID *uuid.UUID, user *entity.User) (err error)
		// UserRegister
		//
		// # 用户注册
		//
		// 用户注册，用于用户注册，需要传递用户名、邮箱、手机号、密码；
		// 该接口将会对用户的数据进行注册，注册成功后将会返回 nil 信息；
		//
		// # 参数
		//   - ctx			上下文(context.Context)
		//   - username		用户名(string)
		//   - email		邮箱(string)
		//   - phone		手机号(string)
		//   - password		密码(string)
		//
		// # 返回
		//   - err			错误信息(error)
		UserRegister(ctx context.Context, username, email, phone, password string) (err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
