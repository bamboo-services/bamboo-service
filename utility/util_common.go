/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
)

// StringToMD5
//
// # 字符串转 MD5
//
// 字符串转 MD5，将字符串转换为 MD5 加密后的字符串；
//
// # 参数
//   - str		字符串(string)
//
// # 返回
//   - string	MD5 加密后的字符串(string)
func StringToMD5(str string) string {
	encrypt, _ := gmd5.EncryptString(str)
	return encrypt
}
