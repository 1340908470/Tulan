package handlerFunc

import (
	"Tulan/engine"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActionHandlerFunc(c *gin.Context) {
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

	if json["action"] != nil {
		action := engine.Action{
			OpenId:        json["open_id"].(string),
			UserId:        json["user_id"].(string),
			OpenMessageId: json["open_message_id"].(string),
			TenantKey:     json["tenant_key"].(string),
			Token:         json["token"].(string),
			Key:           json["action"].(map[string]interface{})["value"].(map[string]interface{})["key"].(string),
			Value:         json["action"].(map[string]interface{})["value"].(map[string]interface{})["value"].(string),
		}
		engine.HandleAction(c, action)
	}
}
