// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProxySubscriptionDao is the data access object for the table fy_proxy_subscription.
type ProxySubscriptionDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  ProxySubscriptionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// ProxySubscriptionColumns defines and stores column names for the table fy_proxy_subscription.
type ProxySubscriptionColumns struct {
	SubscriptionUuid string // 代理订阅地址唯一标识符
	ProxyGroupUuid   string // 代理组唯一标识符
	UserUuid         string // 用户唯一标识符
	Name             string // 代理订阅地址名称
	Merchant         string // 代理订阅地址商户
	Description      string // 代理订阅地址描述
	Url              string // 代理订阅地址URL
	OriginalContent  string // 代理订阅地址所订阅获取的原始内容
	CreatedAt        string // 代理订阅地址创建时间
	UpdatedAt        string // 代理订阅地址更新时间
	SubscribeAt      string // 代理订阅地址订阅时间
}

// proxySubscriptionColumns holds the columns for the table fy_proxy_subscription.
var proxySubscriptionColumns = ProxySubscriptionColumns{
	SubscriptionUuid: "subscription_uuid",
	ProxyGroupUuid:   "proxy_group_uuid",
	UserUuid:         "user_uuid",
	Name:             "name",
	Merchant:         "merchant",
	Description:      "description",
	Url:              "url",
	OriginalContent:  "original_content",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	SubscribeAt:      "subscribe_at",
}

// NewProxySubscriptionDao creates and returns a new DAO object for table data access.
func NewProxySubscriptionDao(handlers ...gdb.ModelHandler) *ProxySubscriptionDao {
	return &ProxySubscriptionDao{
		group:    "default",
		table:    "fy_proxy_subscription",
		columns:  proxySubscriptionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ProxySubscriptionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ProxySubscriptionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ProxySubscriptionDao) Columns() ProxySubscriptionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ProxySubscriptionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ProxySubscriptionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ProxySubscriptionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
