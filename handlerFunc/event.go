package handlerFunc

import (
	"dsl/engine"
	json2 "encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Event struct {
	Schema string                 `json:"schema"`
	Header EventHeader            `json:"header"`
	Event  map[string]interface{} `json:"event"`
}

type EventHeader struct {
	EventId    string `json:"event_id"`
	Token      string `json:"token"`
	CreateTime string `json:"create_time"`
	EventType  string `json:"event_type"`
	TenantKey  string `json:"tenant_key"`
	AppId      string `json:"app_id"`
}

// EventHandlerFunc 处理飞书机器人订阅的消息
func EventHandlerFunc(c *gin.Context) {
	json := make(map[string]interface{})
	err := c.BindJSON(&json)
	if err != nil {
		panic(errors.New("处理事件订阅请求时，json解析失败"))
	}

	// 处理飞书服务器的订阅地址验证Challenge
	if json["type"] != nil && json["type"] == "url_verification" {
		challenge := json["challenge"]
		c.JSON(http.StatusOK, gin.H{
			"challenge": challenge,
		})
	}

	// 处理响应事件
	if json["schema"] != nil && json["schema"] == "2.0" {
		var event Event
		bytes, err := json2.Marshal(json)
		if err != nil {
			panic(errors.New("重编码json时出错"))
		}
		err = json2.Unmarshal(bytes, &event)
		if err != nil {
			panic(errors.New("解析json到event时出错"))
		}

		typ := event.Header.EventType

		// 如果是接受到了消息
		if typ == "im.message.receive_v1" {
			var err = engine.HandleMessageEvent(event.Event)
			panic(err)
		}
	}

	print(json)
}
