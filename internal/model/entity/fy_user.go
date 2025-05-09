// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FyUser is the golang structure for table fy_user.
type FyUser struct {
	UserUuid         string      `json:"userUuid"         orm:"user_uuid"          description:""` //
	Username         string      `json:"username"         orm:"username"           description:""` //
	Email            string      `json:"email"            orm:"email"              description:""` //
	Phone            string      `json:"phone"            orm:"phone"              description:""` //
	Role             string      `json:"role"             orm:"role"               description:""` //
	Permissions      string      `json:"permissions"      orm:"permissions"        description:""` //
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"         description:""` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"         description:""` //
	PasswordHash     string      `json:"passwordHash"     orm:"password_hash"      description:""` //
	EmailVerifiedAt  *gtime.Time `json:"emailVerifiedAt"  orm:"email_verified_at"  description:""` //
	PhoneVerifiedAt  *gtime.Time `json:"phoneVerifiedAt"  orm:"phone_verified_at"  description:""` //
	TwoFactorEnabled bool        `json:"twoFactorEnabled" orm:"two_factor_enabled" description:""` //
	TwoFactorSecret  string      `json:"twoFactorSecret"  orm:"two_factor_secret"  description:""` //
	Nickname         string      `json:"nickname"         orm:"nickname"           description:""` //
	AvatarUrl        string      `json:"avatarUrl"        orm:"avatar_url"         description:""` //
	Gender           string      `json:"gender"           orm:"gender"             description:""` //
	BirthDate        *gtime.Time `json:"birthDate"        orm:"birth_date"         description:""` //
	Bio              string      `json:"bio"              orm:"bio"                description:""` //
	QqEmail          string      `json:"qqEmail"          orm:"qq_email"           description:""` //
	Status           string      `json:"status"           orm:"status"             description:""` //
	LastLoginAt      *gtime.Time `json:"lastLoginAt"      orm:"last_login_at"      description:""` //
	LastLoginIp      string      `json:"lastLoginIp"      orm:"last_login_ip"      description:""` //
	RegistrationIp   string      `json:"registrationIp"   orm:"registration_ip"    description:""` //
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"         description:""` //
}
