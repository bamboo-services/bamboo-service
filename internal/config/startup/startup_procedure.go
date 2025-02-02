/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package startup

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

// initialDatabaseStartup
//
// # 初始化数据库
//
// 初始化数据库，进行数据库的初始化操作；若检查数据库中没有数据则进行初始化操作；
func (s *systemStart) initialDatabaseStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查数据库")
	// 检查数据库是否存在
	createDatabase(s.ctx, "fy_info")
	createDatabase(s.ctx, "fy_permission")
	createDatabase(s.ctx, "fy_resource")
	createDatabase(s.ctx, "fy_role")
	createDatabase(s.ctx, "fy_vip")
	createDatabase(s.ctx, "fy_user")
	createDatabase(s.ctx, "fy_location_ipv4")
	createDatabase(s.ctx, "fy_location_ipv6")
	createDatabase(s.ctx, "fy_album")
	createDatabase(s.ctx, "fy_photo_list")
}

// initialTableContentStartup
//
// # 初始化数据表内容
//
// 初始化数据表内容，进行数据表内容的初始化操作；
func (s *systemStart) initialTableContentStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查数据表内容")

	// Info 表
	g.Log().Infof(s.ctx, "\t检查 fy_info 表数据")
	// 检查数据表
	checkInfoTableValue(s.ctx, "system_name", "竹业")
	checkInfoTableValue(s.ctx, "system_version", "v1.0.0")
	checkInfoTableValue(s.ctx, "system_author", "筱锋xiao_lfeng")
	checkInfoTableValue(s.ctx, "system_referer", "")
	checkInfoTableValue(s.ctx, "has_initial_mode", "1")
	checkInfoTableValue(s.ctx, "web_name", "竹业")
	checkInfoTableValue(s.ctx, "web_description", "竹业，一个基于 Golang 开发的后端服务")
	checkInfoTableValue(s.ctx, "web_keywords", "竹业, Golang, 后端服务")
	checkInfoTableValue(s.ctx, "web_logo", "https://cdn.fengyuecloud.com/2021/09/20210909100000.png")
	checkInfoTableValue(s.ctx, "web_favicon", "https://cdn.fengyuecloud.com/2021/09/20210909100000.png")
	checkInfoTableValue(s.ctx, "web_icp", "粤ICP备2023097009号-2")
	checkInfoTableValue(s.ctx, "web_record", "粤公网安备 44030502003107号")
	checkInfoTableValue(s.ctx, "web_copy", "© 2022-2024 竹业")

	// 邮件内容
	checkInfoTableValue(s.ctx, "mail_host", "smtp.qiye.aliyun.com")
	checkInfoTableValue(s.ctx, "mail_port", "465")
	checkInfoTableValue(s.ctx, "mail_user", "noreplay@x-lf.cn")
	checkInfoTableValue(s.ctx, "mail_password", "password")
	checkInfoTableValue(s.ctx, "mail_name", "锋楪服务")

	// 插入 Aliyun 相关数据表
	checkInfoTableValue(s.ctx, "aliyun_sms_sign_name", "锋楪")
	checkInfoTableValue(s.ctx, "aliyun_sms_code_template", "SMS_468930484")
	checkInfoTableValue(s.ctx, "aliyun_sms_endpoint", "dysmsapi.aliyuncs.com")

	// 全局配置变量存储
	checkInfoTableValue(s.ctx, "ip_4_import_time", "")
	checkInfoTableValue(s.ctx, "ip_4_import_spending", "")
	checkInfoTableValue(s.ctx, "ip_4_ctx", "")
	checkInfoTableValue(s.ctx, "ip_4_period_time", "")
	checkInfoTableValue(s.ctx, "ip_4_upload_total", "")
	checkInfoTableValue(s.ctx, "ip_6_import_time", "")
	checkInfoTableValue(s.ctx, "ip_6_import_spending", "")
	checkInfoTableValue(s.ctx, "ip_6_ctx", "")
	checkInfoTableValue(s.ctx, "ip_6_period_time", "")
	checkInfoTableValue(s.ctx, "ip_6_upload_total", "")
}

// initialRoleStartup
//
// # 初始化角色
//
// 初始化角色，进行角色的初始化操作；
func (s *systemStart) initialRoleStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查角色")

	initializeRole(s.ctx, "admin", "管理员", constant.AdminRolePermission, "管理员角色，用于管理全系统级别权限管理等")
	initializeRole(s.ctx, "user", "用户", constant.UserRolePermission, "用户角色，用于管理用户级别权限管理等")
	initializeRole(s.ctx, "bad", "黑名单", constant.BadRolePermission, "黑名单角色，用于管理黑名单级别权限管理等")
}

// initialSuperAdminStartup
//
// # 初始化超级管理员
//
// 初始化超级管理员，进行超级管理员的初始化操作；
func (s *systemStart) initialSuperAdminStartup() {
	g.Log().Noticef(s.ctx, "[STAR] 检查超级管理员")
	// 检查超级管理员是否存在
	if !hasSuperAdmin(s.ctx) {
		// 创建超级管理员
		createSuperAdmin(s.ctx)

		g.Log().Noticef(s.ctx, "\t用户名：%s", "superAdmin")
		g.Log().Noticef(s.ctx, "\t密码：%s", "admin")
	}
}

// getConstantStorage
//
// # 获取常量存储
//
// 获取常量存储，用于获取常量存储；
func (s *systemStart) getConstantStorage() {
	g.Log().Noticef(s.ctx, "[STAR] 检查常量存储")

	// 获取阿里云授权密钥
	constant.AliyunSmsSignName = getInfoForDB(s.ctx, "aliyun_sms_sign_name")
	constant.AliyunSmsCodeTemplateCode = getInfoForDB(s.ctx, "aliyun_sms_code_template")
	constant.AliyunSmsEndpoint = getInfoForDB(s.ctx, "aliyun_sms_endpoint")

	// 获取系统信息
	constant.InitializeMode = gconv.Bool(getInfoForDB(s.ctx, "has_initial_mode"))
	constant.SystemReferer = getInfoForDB(s.ctx, "system_referer")
	constant.WebFavicon = getInfoForDB(s.ctx, "web_favicon")
	constant.WebLogo = getInfoForDB(s.ctx, "web_logo")
	constant.WebName = getInfoForDB(s.ctx, "web_name")
	constant.WebDescription = getInfoForDB(s.ctx, "web_description")
	constant.WebKeywords = getInfoForDB(s.ctx, "web_keywords")
	constant.WebICP = getInfoForDB(s.ctx, "web_icp")
	constant.WebRecord = getInfoForDB(s.ctx, "web_record")
	constant.WebCopy = getInfoForDB(s.ctx, "web_copy")

	// 获取邮件信息
	constant.MailSMTPHost = getInfoForDB(s.ctx, "mail_host")
	constant.MailSMTPPort = getInfoForDB(s.ctx, "mail_port")
	constant.MailUser = getInfoForDB(s.ctx, "mail_user")
	constant.MailPassword = getInfoForDB(s.ctx, "mail_password")
	constant.MailNickname = getInfoForDB(s.ctx, "mail_name")
}

// getAliyunAuthorizationKey
//
// # 获取阿里云授权密钥
//
// 获取阿里云授权密钥，用于获取阿里云授权密钥；
func (s *systemStart) getAliyunAuthorizationKey() {
	g.Log().Noticef(s.ctx, "[STAR] 获取阿里云授权密钥")
	json := gjson.New(gfile.GetContents("access.json"))
	constant.AliyunAccessKey = json.Get("AliyunKey.AccessKeyID").String()
	constant.AliyunSecretKey = json.Get("AliyunKeyKey.AccessKeySecret").String()
	constant.DogeCloudAccessKey = json.Get("DogeCloudKey.AccessKey").String()
	constant.DogeCloudSecretKey = json.Get("DogeCloudKey.SecretKey").String()
}

// dogeCloudKey
//
// # 获取多吉云授权密钥
//
// 获取多吉云授权密钥，用于获取多吉云授权密钥；
func (s *systemStart) dogeCloudKey() {
	g.Log().Noticef(s.ctx, "[STAR] 获取多吉云授权密钥")
	bucket, err := service.DogeCloud().GetAccessTokenAPI(s.ctx)
	if err != nil {
		g.Log().Panic(s.ctx, err.Error())
	}
	// 缓存数据重写
	err = g.Redis().HMSet(s.ctx, "global:dc:bucket", gconv.Map(bucket))
	if err != nil {
		g.Log().Panic(s.ctx, err.Error())
	}
	// 设置过期时间
	_, err = g.Redis().ExpireAt(s.ctx, "global:dc:bucket", bucket.ExpiredAt.Time)
	if err != nil {
		g.Log().Panic(s.ctx, err.Error())
	}
}

// checkFolder
//
// # 检查文件夹
//
// 检查文件夹，用于检查文件夹是否存在；
func (s *systemStart) checkFolder() {
	g.Log().Noticef(s.ctx, "[STAR] 检查文件夹")

	checkAndCreateFolder(s.ctx, "upload")
	checkAndCreateFolder(s.ctx, "upload/ip_location")
}
