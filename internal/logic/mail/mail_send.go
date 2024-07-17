/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package mail

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/model/dto"
	"bamboo-service/utility"
	"context"
	"fmt"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gtime"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
)

// SendCodeMail
//
// # 发送验证码邮件
//
// 发送验证码邮件，发送验证码邮件到指定邮箱；
// 验证码将发送到指定 mail 邮箱中，验证码为 code；
// 验证码的有效期为 5 分钟；该接口将会为输入的 code 进行存入缓存中，用于后续的验证；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - mail		邮箱(string)
//   - code		验证码(string)
//
// # 返回
//   - error	错误信息(error)
func (s *sMail) SendCodeMail(ctx context.Context, mail, code string) (err error) {
	g.Log().Notice(ctx, "[SERV] mail.SendCodeMail | 发送验证码邮件接口")
	// 检查是否可以再次发送验证码
	getExpireTime, err := g.Redis().ExpireTime(ctx, "mail:code:"+utility.StringToMD5(mail))
	g.Log().Debug(ctx, getExpireTime)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "获取验证码缓存失败")
	}
	// 检查是否小于 4 分钟
	if getExpireTime.GTime().Timestamp()-gtime.Now().Timestamp() > 240 {
		second := getExpireTime.GTime().Timestamp() - gtime.Now().Timestamp() - 240
		return berror.NewError(
			bcode.OperationNotAllow,
			"验证码发送时间过短，剩余 "+strconv.FormatInt(second, 10)+" 秒",
		)
	}
	// 发送验证码
	_, err = g.Redis().Set(ctx, "mail:code:"+utility.StringToMD5(mail), code)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "验证码缓存设置失败")
	}
	_, err = g.Redis().Expire(ctx, "mail:code:"+utility.StringToMD5(mail), 300)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "验证码缓存时间设置失败")
	}
	// 设置变量
	value := make([]dto.MailVariableDTO, 0)
	value = append(value, dto.MailVariableDTO{
		Key:   "code",
		Value: code,
	})
	err = s.sendMail(ctx, mail, constant.WebName+"-验证码", "mail_code", value)
	if err != nil {
		return err
	}
	return nil
}

// sendMail
//
// # 发送邮件
//
// 发送邮件，发送邮件到指定邮箱；
//
// # 参数
//   - ctx		上下文(context.Context)
//   - mail		邮箱(string)
//   - title	标题(string)
//   - tpl		模板(string)
//   - value	自定义参数([]dto.MailVariableDTO)
//
// # 返回
//   - error	错误信息(error)
func (s *sMail) sendMail(ctx context.Context, mail, title, tpl string, value []dto.MailVariableDTO) (err error) {
	g.Log().Notice(ctx, "[SERV] mail.sendMail | 发送邮件接口")
	// 发送邮件
	plainAuth := smtp.PlainAuth("", constant.MailUser, constant.MailPassword, constant.MailSMTPHost)
	// 获取发送邮件内容
	if !gres.Contains("resource/template/mail/" + tpl + ".html") {
		return berror.NewError(bcode.NotExist, "不存在 "+tpl+" 模板")
	}
	getContent := gres.GetContent("resource/template/mail/" + tpl + ".html")
	// 对模板进行参数替换
	getReplace := strings.ReplaceAll(string(getContent), "%title%", title)
	getReplace = strings.ReplaceAll(getReplace, "%email%", mail)
	getReplace = strings.ReplaceAll(getReplace, "%footer_message%", constant.WebCopy)
	// 去除版权
	getReplace = string(regexp.MustCompile("<!--(\\S| |\\n)+-->\n").ReplaceAll([]byte(getReplace), []byte("")))
	// 自定义参数批量替换
	for _, getValue := range value {
		getReplace = strings.ReplaceAll(getReplace, "%"+getValue.Key+"%", getValue.Value)
	}
	// 检查是否存在剩余未赋值变量
	matched, err := regexp.MatchString("%\\S+%", getReplace)
	if err != nil {
		return berror.NewErrorHasError(bcode.UnknownError, err, "正则匹配失败")
	}
	if matched {
		// 返回仍有未复制变量的值
		compile := regexp.MustCompile("%\\S+%")
		findString := compile.FindString(getReplace)
		return berror.NewError(bcode.OperationNotAllow, "存在未赋值变量："+findString)
	}
	// 准备数据
	contentType := "Content-Type: text/html; charset=UTF-8"
	getReplace = fmt.Sprintf("To: %s\r\nFrom: %s <%s>\r\nSubject: %s\r\nMessage-ID: %s\r\n%s\r\n\r\n%s",
		mail, constant.MailNickname, constant.MailUser, title, gtime.Now().TimestampNanoStr(), contentType, getReplace)
	g.Log().Debug(ctx, getReplace)
	// 准备发送邮件
	address := fmt.Sprintf("%s:%s", constant.MailSMTPHost, constant.MailSMTPPort)
	err = smtp.SendMail(address, plainAuth, constant.MailUser, []string{mail}, []byte(getReplace))
	if err != nil {
		return berror.NewErrorHasError(bcode.UnknownError, err, "发送邮件失败")
	}
	return nil
}
