package mail

import (
	"bamboo-service/internal/consts"
	"bamboo-service/internal/custom"
	"bamboo-service/internal/model/dto"
	"context"
	"fmt"
	"github.com/XiaoLFeng/bamboo-utils/berror"
	"github.com/XiaoLFeng/bamboo-utils/blog"
	"github.com/XiaoLFeng/bamboo-utils/butil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
)

// CheckMailTemplate 检查提供的邮件模板名称是否有效。
//
// 参数:
//   - ctx: 上下文，用于控制生命周期和日志记录。
//   - template: 模板名称，待检查的邮件模板标识。
//
// 返回:
//   - *berror.ErrorCode: 错误信息，如果模板无效或为空则返回相应错误。
func (s *sMail) CheckMailTemplate(ctx context.Context, template string) *berror.ErrorCode {
	blog.ServiceDebug(ctx, "CheckMailTemplate", "检查模板 %s", template)
	if template == "" {
		return berror.ErrorAddData(berror.ErrInvalidParameters, "模板名字不能为空")
	}
	// 循环列表，检查是否是允许的模板类型
	for _, mailTemplate := range consts.MailTemplateList {
		if mailTemplate.Name == template {
			return nil
		}
	}
	return berror.ErrorAddData(berror.ErrInvalidParameters, "模板名字不存在")
}

// SendMail 发送邮件至指定邮箱并填充模板数据。
//
// 参数:
//   - ctx: 用于控制请求生命周期和传递上下文信息。
//   - template: 模板名称，用于指定邮件内容格式。
//   - mailTemplate: 包含邮件模板所需的补充数据。
//
// 返回:
//   - *berror.ErrorCode: 错误信息，表示模板检查失败、邮件发送失败等可能原因。
func (s *sMail) SendMail(ctx context.Context, template string, mailTemplate *dto.MailSendTemplateDTO) *berror.ErrorCode {
	blog.ServiceInfo(ctx, "SendMail", "发送邮件 %s 邮件模板 %s", mailTemplate.Email, template)
	// 检查模板
	errorCode := s.CheckMailTemplate(ctx, template)
	if errorCode != nil {
		return errorCode
	}

	// 模板数据补充构建
	mailTemplate.AuthorName = consts.SystemAuthorNameValue
	mailTemplate.AuthorLink = consts.SystemAuthorLinkValue
	mailTemplate.Minute = 30
	mailTemplate.WebName = consts.SystemNameValue
	mailTemplate.WebLink = consts.SystemLinkValue

	// 获取模板
	getResult, errorCode := getMailTemplate(ctx, template, mailTemplate)
	if errorCode != nil {
		return nil
	}

	// 发送邮件
	errorCode = send(ctx, mailTemplate.Email, fmt.Sprintf("%s - 验证码", consts.SystemNameValue), *getResult)
	return errorCode
}

// getMailTemplate 根据指定名称获取邮件模板并替换变量占位符。
//
// 参数:
//   - ctx: 上下文，用于日志记录及控制生命周期。
//   - template: 模板名称，用于定位邮件内容格式。
//   - mailTemplate: 包含需要替换占位符的数据结构。
//
// 返回:
//   - *string: 处理后的邮件模板内容。
//   - *berror.ErrorCode: 错误信息，若模板无法获取或处理失败则返回对应错误。
func getMailTemplate(ctx context.Context, template string, mailTemplate *dto.MailSendTemplateDTO) (*string, *berror.ErrorCode) {
	blog.ServiceDebug(ctx, "getMailTemplate", "获取模板 %s", template)
	getResourceByte := gres.GetContent("template/mail/" + template + ".html")
	if len(getResourceByte) == 0 {
		return nil, berror.ErrorAddData(berror.ErrNotFound, fmt.Sprintf("模板文件 %s 不存在", template))
	}
	getResource := string(getResourceByte)
	for key, val := range butil.StructToMap(mailTemplate) {
		getResource = strings.ReplaceAll(getResource, "{{."+key+"}}", g.NewVar(val).String())
	}
	return &getResource, nil
}

// send 发送邮件至指定邮箱。
//
// 参数:
//   - targetMail: 接收邮件的目标邮箱地址。
//   - subject: 邮件的主题内容。
//   - result: 邮件的HTML内容。
//
// 返回:
//   - *berror.ErrorCode: 错误信息，表示发送邮件失败的原因。
func send(ctx context.Context, targetMail string, subject string, result string) *berror.ErrorCode {
	getMailConfig, configErr := g.Cfg().GetWithEnv(ctx, "custom.mail")
	if configErr != nil {
		blog.ServiceError(ctx, "send", "获取邮件配置失败 %v", configErr)
		return berror.ErrorAddData(*custom.ErrorMailConfigFailed, configErr)
	}
	mailConfig := getMailConfig.Map()
	newEmail := email.NewEmail()
	newEmail.From = fmt.Sprintf("%s <%s>", mailConfig["name"], mailConfig["user"])
	newEmail.To = []string{targetMail}
	newEmail.Subject = subject
	newEmail.HTML = []byte(result)
	mailErr := newEmail.Send(fmt.Sprintf("%s:%s", mailConfig["host"], mailConfig["port"]), smtp.PlainAuth("", mailConfig["user"].(string), mailConfig["pass"].(string), mailConfig["host"].(string)))
	if mailErr != nil {
		blog.ServiceError(context.Background(), "send", "发送邮件失败 %v", mailErr)
		return berror.ErrorAddData(*custom.ErrorMailSendFailed, mailErr)
	}
	return nil
}
