package v1

import (
	"bamboo-service/internal/model/dto"
	"github.com/XiaoLFeng/bamboo-utils/bmodels"
	"github.com/gogf/gf/v2/frame/g"
)

type MailSendCodeReq struct {
	g.Meta   `path:"/mail/code/send" method:"post" sm:"发送验证码" tags:"邮件控制器" description:"用作需要邮件验证的地方进行发送验证码「验证码有效期 15分钟」"`
	Email    string `json:"email" v:"required|email#邮箱不能为空|邮箱格式不正确" dc:"邮箱"`
	Purpose  string `json:"purpose" v:"required#请输入验证码用途" dc:"验证码用途"`
	Template string `json:"template" v:"required#请输入验证码模板" dc:"验证码模板"`
}

type MailSendCodeRes struct {
	g.Meta `mime:"application/json;utf-8"`
	*bmodels.ResponseDTO[*dto.MailCodeDTO]
}
