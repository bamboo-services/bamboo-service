// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmailCodeDao is the data access object for the table fy_email_code.
type EmailCodeDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  EmailCodeColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// EmailCodeColumns defines and stores column names for the table fy_email_code.
type EmailCodeColumns struct {
	CodeUuid  string // 验证码主键
	Email     string // 邮箱地址
	Code      string // 验证码
	Purpose   string // 验证码用途：register-注册,reset-重置密码,bind-绑定
	ExpiredAt string // 过期时间
	CreatedAt string // 创建时间
}

// emailCodeColumns holds the columns for the table fy_email_code.
var emailCodeColumns = EmailCodeColumns{
	CodeUuid:  "code_uuid",
	Email:     "email",
	Code:      "code",
	Purpose:   "purpose",
	ExpiredAt: "expired_at",
	CreatedAt: "created_at",
}

// NewEmailCodeDao creates and returns a new DAO object for table data access.
func NewEmailCodeDao(handlers ...gdb.ModelHandler) *EmailCodeDao {
	return &EmailCodeDao{
		group:    "default",
		table:    "fy_email_code",
		columns:  emailCodeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmailCodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmailCodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmailCodeDao) Columns() EmailCodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmailCodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmailCodeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EmailCodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
