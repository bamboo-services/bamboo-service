// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PermissionDao is the data access object for the table fy_permission.
type PermissionDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PermissionColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PermissionColumns defines and stores column names for the table fy_permission.
type PermissionColumns struct {
	PermissionKey         string // 权限标识
	PermissionName        string // 权限名称
	PermissionDescription string // 权限描述
	PermissionStatus      string // 权限状态
	CreatedAt             string // 创建时间
	UpdatedAt             string // 更新时间
}

// permissionColumns holds the columns for the table fy_permission.
var permissionColumns = PermissionColumns{
	PermissionKey:         "permission_key",
	PermissionName:        "permission_name",
	PermissionDescription: "permission_description",
	PermissionStatus:      "permission_status",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewPermissionDao creates and returns a new DAO object for table data access.
func NewPermissionDao(handlers ...gdb.ModelHandler) *PermissionDao {
	return &PermissionDao{
		group:    "default",
		table:    "fy_permission",
		columns:  permissionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PermissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PermissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PermissionDao) Columns() PermissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PermissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PermissionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PermissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
