package web

import "dsl/model"

// ApiSendMessageCard 给指定用户或者会话发送消息，支持文本、富文本、可交互的消息卡片、群名片、个人名片、图片、视频、音频、文件、表情包。
var ApiSendMessageCard = "https://open.feishu.cn/open-apis/message/v4/send/"

type ApiSendMessageCardReq model.Message
type ApiSendMessageCardRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		MessageId string `json:"message_id"`
	} `json:"data"`
}

// ApiTenantAccessToken 获取 tenant_access_token（企业自建应用）
var ApiTenantAccessToken = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

type ApiTenantAccessTokenReq struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}
type ApiTenantAccessTokenRes struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}
