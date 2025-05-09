// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-05-10 03:00:04
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemDao is the data access object for the table fy_system.
type SystemDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SystemColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SystemColumns defines and stores column names for the table fy_system.
type SystemColumns struct {
	SystemUuid string // 系统唯一标识符
	Key        string // 键
	Value      string // 值
	Version    string // 版本
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// systemColumns holds the columns for the table fy_system.
var systemColumns = SystemColumns{
	SystemUuid: "system_uuid",
	Key:        "key",
	Value:      "value",
	Version:    "version",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewSystemDao creates and returns a new DAO object for table data access.
func NewSystemDao(handlers ...gdb.ModelHandler) *SystemDao {
	return &SystemDao{
		group:    "default",
		table:    "fy_system",
		columns:  systemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SystemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SystemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SystemDao) Columns() SystemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SystemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SystemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SystemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
