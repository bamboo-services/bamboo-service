// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EmailCode is the golang structure of table fy_email_code for DAO operations like Where/Data.
type EmailCode struct {
	g.Meta    `orm:"table:fy_email_code, do:true"`
	CodeUuid  interface{} // 验证码主键
	Email     interface{} // 邮箱地址
	Code      interface{} // 验证码
	Purpose   interface{} // 验证码用途：register-注册,reset-重置密码,bind-绑定
	ExpiredAt *gtime.Time // 过期时间
	CreatedAt *gtime.Time // 创建时间
}
