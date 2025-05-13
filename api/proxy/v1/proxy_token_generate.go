package v1

import (
	"bamboo-service/internal/model/dto"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProxyTokenGenerateReq struct {
	g.Meta      `path:"/proxy/token" method:"Post" sm:"生成代理授权令牌" tags:"代理控制器" dc:"用于生成代理授权令牌「允许远程调用接口使用」"`
	Name        string      `json:"name" v:"required#请输入令牌名称" dc:"请输入令牌名称"`
	Description string      `json:"description" v:"length:0,255#令牌描述长度不能超过255个字符" dc:"令牌描述"`
	ExpiredAt   *gtime.Time `json:"expired_at" v:"required|regex:^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$#请输入令牌过期时间|令牌过期时间格式不正确" dc:"请输入令牌过期时间"`
}

type ProxyTokenGenerateRes struct {
	g.Meta `mime:"application/json;charset=utf-8"`
	*bmodels.ResponseDTO[*dto.ProxyTokenDTO]
}
