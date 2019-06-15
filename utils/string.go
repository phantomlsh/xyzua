package utils

import (
	"math/rand"
	"time"
)

// constants

const SALT = "XYZUA_Powered_by_Phantomlsh"

// Message
type MSG struct {
	ParamErr    string
	Forbidden   string
	ContextErr  string
	LoginErr    string
	TokenErr    string
	DisabledErr string
	UnknownErr  string
	AppErr      string
	CodeErr     string
}

var Msg MSG

func init() {
	Msg.AppErr = "App Name or Secret Error"     // "应用程序或密钥错误"
	Msg.CodeErr = "Code Error"                  // "跨平台凭证错误"
	Msg.ContextErr = "Context Error"            // "上下文错误"
	Msg.DisabledErr = "Disabled Account"        // "账户被禁用"
	Msg.Forbidden = "Forbidden"                 // "拒绝访问"
	Msg.LoginErr = "Username or Password Error" // "用户名或密码错误"
	Msg.ParamErr = "Param Error"                // "参数错误"
	Msg.TokenErr = "Token Error"                // "凭证错误"
	Msg.UnknownErr = "Unknown Error"            // "未知错误"
}

// functions
// get random string with length l
func RandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
