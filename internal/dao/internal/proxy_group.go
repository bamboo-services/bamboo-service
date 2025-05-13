// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProxyGroupDao is the data access object for the table fy_proxy_group.
type ProxyGroupDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ProxyGroupColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ProxyGroupColumns defines and stores column names for the table fy_proxy_group.
type ProxyGroupColumns struct {
	GroupUuid   string // 代理组UUID
	UserUuid    string // 用户UUID
	Name        string // 代理组名称
	FileName    string // 代理组文件名
	Description string // 代理组描述
	Proxy       string // 代理组代理
	Partition   string // 代理组分区
	Rule        string // 代理组规则
	IsEnabled   string // 代理组是否启用
	CreatedAt   string // 代理组创建时间
	UpdatedAt   string // 代理组更新时间
}

// proxyGroupColumns holds the columns for the table fy_proxy_group.
var proxyGroupColumns = ProxyGroupColumns{
	GroupUuid:   "group_uuid",
	UserUuid:    "user_uuid",
	Name:        "name",
	FileName:    "file_name",
	Description: "description",
	Proxy:       "proxy",
	Partition:   "partition",
	Rule:        "rule",
	IsEnabled:   "is_enabled",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewProxyGroupDao creates and returns a new DAO object for table data access.
func NewProxyGroupDao(handlers ...gdb.ModelHandler) *ProxyGroupDao {
	return &ProxyGroupDao{
		group:    "default",
		table:    "fy_proxy_group",
		columns:  proxyGroupColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProxyGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProxyGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProxyGroupDao) Columns() ProxyGroupColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProxyGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProxyGroupDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProxyGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
