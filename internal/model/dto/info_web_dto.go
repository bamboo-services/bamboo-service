/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package dto

type InfoWebDTO struct {
	WebName        string `json:"web_name" dc:"网站名称"`         // 网站名称
	WebDescription string `json:"web_description" dc:"网站描述"`  // 网站描述
	WebKeywords    string `json:"web_keywords" dc:"网站关键字"`    // 网站关键字
	WebLogo        string `json:"web_logo" dc:"网站Logo"`       // 网站Logo
	WebFavicon     string `json:"web_favicon" dc:"网站Favicon"` // 网站Favicon
	WebICP         string `json:"web_icp" dc:"网站备案号"`         // 网站备案号
	WebRecord      string `json:"web_record" dc:"网站备案地址"`     // 网站备案地址
	WebCopy        string `json:"web_copy" dc:"网站版权"`         // 网站版权
}
