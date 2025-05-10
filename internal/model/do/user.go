// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table fy_user for DAO operations like Where/Data.
type User struct {
	g.Meta           `orm:"table:fy_user, do:true"`
	UserUuid         interface{} // 用户唯一标识符
	Username         interface{} // 用户名
	Email            interface{} // 电子邮箱
	Phone            interface{} // 手机号码
	Role             interface{} // 用户角色
	Permissions      *gjson.Json // 用户权限
	CreatedAt        *gtime.Time // 记录创建时间
	UpdatedAt        *gtime.Time // 记录更新时间
	PasswordHash     interface{} // 密码哈希值
	EmailVerifiedAt  *gtime.Time // 邮箱验证时间
	PhoneVerifiedAt  *gtime.Time // 手机验证时间
	TwoFactorEnabled interface{} // 是否启用两因素认证
	TwoFactorSecret  interface{} // 两因素认证密钥
	Nickname         interface{} // 用户昵称
	AvatarUrl        interface{} // 头像URL
	AvatarBase64     interface{} // 头像Base64编码
	Gender           interface{} // 性别
	BirthDate        *gtime.Time // 出生日期
	Bio              interface{} // 个人简介
	QqEmail          interface{} // QQ邮箱
	Status           interface{} // 用户账户状态
	LastLoginAt      *gtime.Time // 最后登录时间
	LastLoginIp      interface{} // 最后登录IP地址
	RegistrationIp   interface{} // 注册IP地址
	DeletedAt        *gtime.Time // 删除时间（软删除）
}
