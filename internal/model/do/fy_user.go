// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FyUser is the golang structure of table fy_user for DAO operations like Where/Data.
type FyUser struct {
	g.Meta           `orm:"table:fy_user, do:true"`
	UserUuid         interface{} //
	Username         interface{} //
	Email            interface{} //
	Phone            interface{} //
	Role             interface{} //
	Permissions      interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
	PasswordHash     interface{} //
	EmailVerifiedAt  *gtime.Time //
	PhoneVerifiedAt  *gtime.Time //
	TwoFactorEnabled interface{} //
	TwoFactorSecret  interface{} //
	Nickname         interface{} //
	AvatarUrl        interface{} //
	Gender           interface{} //
	BirthDate        *gtime.Time //
	Bio              interface{} //
	QqEmail          interface{} //
	Status           interface{} //
	LastLoginAt      *gtime.Time //
	LastLoginIp      interface{} //
	RegistrationIp   interface{} //
	DeletedAt        *gtime.Time //
}
