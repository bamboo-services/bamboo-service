// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-09 15:09:16
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table fy_user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table fy_user.
type UserColumns struct {
	UserUuid         string // 用户唯一标识符
	Username         string // 用户名
	Email            string // 电子邮箱
	Phone            string // 手机号码
	Role             string // 用户角色
	Permissions      string // 用户权限
	CreatedAt        string // 记录创建时间
	UpdatedAt        string // 记录更新时间
	PasswordHash     string // 密码哈希值
	EmailVerifiedAt  string // 邮箱验证时间
	PhoneVerifiedAt  string // 手机验证时间
	TwoFactorEnabled string // 是否启用两因素认证
	TwoFactorSecret  string // 两因素认证密钥
	Nickname         string // 用户昵称
	AvatarUrl        string // 头像URL
	Gender           string // 性别
	BirthDate        string // 出生日期
	Bio              string // 个人简介
	QqEmail          string // QQ邮箱
	Status           string // 用户账户状态
	LastLoginAt      string // 最后登录时间
	LastLoginIp      string // 最后登录IP地址
	RegistrationIp   string // 注册IP地址
	DeletedAt        string // 删除时间（软删除）
}

// userColumns holds the columns for the table fy_user.
var userColumns = UserColumns{
	UserUuid:         "user_uuid",
	Username:         "username",
	Email:            "email",
	Phone:            "phone",
	Role:             "role",
	Permissions:      "permissions",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	PasswordHash:     "password_hash",
	EmailVerifiedAt:  "email_verified_at",
	PhoneVerifiedAt:  "phone_verified_at",
	TwoFactorEnabled: "two_factor_enabled",
	TwoFactorSecret:  "two_factor_secret",
	Nickname:         "nickname",
	AvatarUrl:        "avatar_url",
	Gender:           "gender",
	BirthDate:        "birth_date",
	Bio:              "bio",
	QqEmail:          "qq_email",
	Status:           "status",
	LastLoginAt:      "last_login_at",
	LastLoginIp:      "last_login_ip",
	RegistrationIp:   "registration_ip",
	DeletedAt:        "deleted_at",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "fy_user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
