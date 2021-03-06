// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"net/http"
	"net/url"
)

// 微信服务器推送过来的消息(事件)处理接口
type MessageHandler interface {
	ServeMessage(http.ResponseWriter, *Request)
}

type MessageHandlerFunc func(http.ResponseWriter, *Request)

func (fn MessageHandlerFunc) ServeMessage(w http.ResponseWriter, r *Request) {
	fn(w, r)
}

// 消息(事件)请求信息
type Request struct {
	HttpRequest *http.Request // 可以为 nil, 因为某些 http 框架没有提供此参数

	// 下面的字段必须提供

	QueryValues  url.Values // 回调请求 URL 中的查询参数集合
	MsgSignature string     // 回调请求 URL 中的消息体签名: msg_signature
	EncryptType  string     // 回调请求 URL 中的加密方式: encrypt_type
	Timestamp    int64      // 回调请求 URL 中的时间戳: timestamp
	Nonce        string     // 回调请求 URL 中的随机数: nonce

	RawMsgXML []byte        // 消息的"明文"XML 文本
	MixedMsg  *MixedMessage // RawMsgXML 解析后的消息

	AESKey [32]byte // 当前消息 AES 加密的 key
	Random []byte   // 当前消息加密时所用的 random, 16 bytes

	AppId string // 请求消息所属第三方平台的 AppId
	Token string // 请求消息所属第三方平台的 Token
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId      string `xml:"AppId"      json:"AppId"`
	CreateTime int64  `xml:"CreateTime" json:"CreateTime"`
	InfoType   string `xml:"InfoType"   json:"InfoType"`

	VerifyTicket    string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"`
	AuthorizerAppId string `xml:"AuthorizerAppid"       json:"AuthorizerAppid"`
}
