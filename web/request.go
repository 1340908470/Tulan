package web

import (
	"bytes"
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var appId = "cli_a128a2cdb838d00e"
var appSecret = "bUO4olqjZKkWua4KocwPDeJIzqFOQJr0"

var tokenTime int64 = 0
var tenantAccessToken = ""

func sendRequest(api string, req interface{}, header map[string]string, paras map[string]string) []byte {
	json, _ := json2.Marshal(req)
	request, err := http.NewRequest("POST", api, bytes.NewReader(json))
	if err != nil {
		panic(err)
	}

	// 添加头
	for key, value := range header {
		request.Header.Set(key, value)
	}
	request.Header.Set("Content-Type", "application/json")

	if len(paras) > 0 {
		// 添加查询参数
		query := request.URL.Query()
		for key, value := range paras {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	res, _ := ioutil.ReadAll(resp.Body)
	return res
}

func Request(api string, req interface{}, paras map[string]string) []byte {
	// 首先判断token是否过期
	now := time.Now().UnixNano() / 1e6
	// 以 1.5h 为过期期限
	if now-tokenTime > 3*30*60*1000 || tenantAccessToken == "" {
		tokenReq := ApiTenantAccessTokenReq{
			AppId:     appId,
			AppSecret: appSecret,
		}
		header := make(map[string]string)
		var tokenRes ApiTenantAccessTokenRes
		json := sendRequest(ApiTenantAccessToken, tokenReq, header, make(map[string]string))
		json2.Unmarshal(json, &tokenRes)
		tenantAccessToken = tokenRes.TenantAccessToken
	}

	// 然后正式进行请求
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %v", tenantAccessToken)
	res := sendRequest(api, req, header, paras)
	return res
}
