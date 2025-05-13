// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ProxyToken is the golang structure of table fy_proxy_token for DAO operations like Where/Data.
type ProxyToken struct {
	g.Meta         `orm:"table:fy_proxy_token, do:true"`
	ProxyTokenUuid interface{} // 代理令牌UUID
	UserUuid       interface{} // 用户UUID
	Name           interface{} // 代理令牌名称
	Description    interface{} // 代理令牌描述
	CreatedAt      *gtime.Time // 代理令牌创建时间
	ExpiredAt      *gtime.Time // 代理令牌过期时间
}
