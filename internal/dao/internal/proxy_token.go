// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProxyTokenDao is the data access object for the table fy_proxy_token.
type ProxyTokenDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProxyTokenColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProxyTokenColumns defines and stores column names for the table fy_proxy_token.
type ProxyTokenColumns struct {
	ProxyTokenUuid string // 代理令牌UUID
	UserUuid       string // 用户UUID
	Name           string // 代理令牌名称
	Description    string // 代理令牌描述
	CreatedAt      string // 代理令牌创建时间
	ExpiredAt      string // 代理令牌过期时间
}

// proxyTokenColumns holds the columns for the table fy_proxy_token.
var proxyTokenColumns = ProxyTokenColumns{
	ProxyTokenUuid: "proxy_token_uuid",
	UserUuid:       "user_uuid",
	Name:           "name",
	Description:    "description",
	CreatedAt:      "created_at",
	ExpiredAt:      "expired_at",
}

// NewProxyTokenDao creates and returns a new DAO object for table data access.
func NewProxyTokenDao(handlers ...gdb.ModelHandler) *ProxyTokenDao {
	return &ProxyTokenDao{
		group:    "default",
		table:    "fy_proxy_token",
		columns:  proxyTokenColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProxyTokenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProxyTokenDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProxyTokenDao) Columns() ProxyTokenColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProxyTokenDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProxyTokenDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProxyTokenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
