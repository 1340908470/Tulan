package web

import (
	"bytes"
	json2 "encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var appId = os.Getenv("FEISHU_TULAN_APPID")
var appSecret = os.Getenv("FEISHU_TULAN_APPSECRET")

var tokenTime int64 = 0
var tenantAccessToken = ""

func sendRequest(api string, req interface{}, header map[string]string) interface{} {
	json, _ := json2.Marshal(req)
	request, err := http.NewRequest("POST", api, bytes.NewReader(json))
	if err != nil {
		panic(err)
	}
	for key, value := range header {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	res, _ := ioutil.ReadAll(resp.Body)
	return res
}

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

func Request(api string, req interface{}) interface{} {
	// 首先判断token是否过期
	now := time.Now().UnixNano() / 1e6
	// 以 1.5h 为过期期限
	if now-tokenTime > 3*30*60*1000 || tenantAccessToken == "" {
		tokenReq := ApiTenantAccessTokenReq{
			AppId:     appId,
			AppSecret: appSecret,
		}
		header := make(map[string]string)
		tokenRes := sendRequest(ApiTenantAccessToken, tokenReq, header).(ApiTenantAccessTokenRes)
		tenantAccessToken = tokenRes.TenantAccessToken
	}

	// 然后正式进行请求
	header := make(map[string]string)
	header["Authorization"] = tenantAccessToken
	res := sendRequest(api, req, header)
	return res
}
