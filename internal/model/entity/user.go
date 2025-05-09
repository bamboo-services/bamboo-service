// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-09 21:32:25
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserUuid         string      `json:"user_uuid"          orm:"user_uuid"          description:"用户唯一标识符"`   // 用户唯一标识符
	Username         string      `json:"username"           orm:"username"           description:"用户名"`       // 用户名
	Email            string      `json:"email"              orm:"email"              description:"电子邮箱"`      // 电子邮箱
	Phone            string      `json:"phone"              orm:"phone"              description:"手机号码"`      // 手机号码
	Role             string      `json:"role"               orm:"role"               description:"用户角色"`      // 用户角色
	Permissions      *gjson.Json `json:"permissions"        orm:"permissions"        description:"用户权限"`      // 用户权限
	CreatedAt        *gtime.Time `json:"created_at"         orm:"created_at"         description:"记录创建时间"`    // 记录创建时间
	UpdatedAt        *gtime.Time `json:"updated_at"         orm:"updated_at"         description:"记录更新时间"`    // 记录更新时间
	PasswordHash     string      `json:"password_hash"      orm:"password_hash"      description:"密码哈希值"`     // 密码哈希值
	EmailVerifiedAt  *gtime.Time `json:"email_verified_at"  orm:"email_verified_at"  description:"邮箱验证时间"`    // 邮箱验证时间
	PhoneVerifiedAt  *gtime.Time `json:"phone_verified_at"  orm:"phone_verified_at"  description:"手机验证时间"`    // 手机验证时间
	TwoFactorEnabled bool        `json:"two_factor_enabled" orm:"two_factor_enabled" description:"是否启用两因素认证"` // 是否启用两因素认证
	TwoFactorSecret  string      `json:"two_factor_secret"  orm:"two_factor_secret"  description:"两因素认证密钥"`   // 两因素认证密钥
	Nickname         string      `json:"nickname"           orm:"nickname"           description:"用户昵称"`      // 用户昵称
	AvatarUrl        string      `json:"avatar_url"         orm:"avatar_url"         description:"头像URL"`     // 头像URL
	Gender           string      `json:"gender"             orm:"gender"             description:"性别"`        // 性别
	BirthDate        *gtime.Time `json:"birth_date"         orm:"birth_date"         description:"出生日期"`      // 出生日期
	Bio              string      `json:"bio"                orm:"bio"                description:"个人简介"`      // 个人简介
	QqEmail          string      `json:"qq_email"           orm:"qq_email"           description:"QQ邮箱"`      // QQ邮箱
	Status           string      `json:"status"             orm:"status"             description:"用户账户状态"`    // 用户账户状态
	LastLoginAt      *gtime.Time `json:"last_login_at"      orm:"last_login_at"      description:"最后登录时间"`    // 最后登录时间
	LastLoginIp      string      `json:"last_login_ip"      orm:"last_login_ip"      description:"最后登录IP地址"`  // 最后登录IP地址
	RegistrationIp   string      `json:"registration_ip"    orm:"registration_ip"    description:"注册IP地址"`    // 注册IP地址
	DeletedAt        *gtime.Time `json:"deleted_at"         orm:"deleted_at"         description:"删除时间（软删除）"` // 删除时间（软删除）
}
