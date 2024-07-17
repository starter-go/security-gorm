package vericode

import "github.com/starter-go/base/lang"

// Code 表示验证码
type Code string

// DataObject 用于存储数据的对象
type DataObject struct {

	// Nonce lang.Hex    // 序列号，要求前后端一致 (这个字段由前端保存)

	Sum   lang.Hex    // 对验证码的 hash 计算结果
	Salt  lang.Base64 // 用于对验证码的 hash 计算
	Addr  string      // 表示账号的地址(mail_addr|phone_num)
	TTL   int         // 剩余重试次数
	Time1 lang.Time   // 起效时间
	Time2 lang.Time   // 失效时间
}

// Verification 表示一个验证请求
type Verification struct {
	Nonce []byte // 序列号，要求前后端一致
	Code  Code   // 验证码
	Addr  string // 表示账号的地址(mail_addr|phone_num)
}
