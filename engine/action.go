package engine

import (
	"dsl/def"
	"dsl/model"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Action struct {
	OpenId        string `json:"open_id"`
	UserId        string `json:"user_id"`
	OpenMessageId string `json:"open_message_id"`
	TenantKey     string `json:"tenant_key"`
	Token         string `json:"token"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

func HandleAction(c *gin.Context, action Action) {
	// 首先判断是否是 trigger_action，并依此判断是否进入事务
	if action.Key == "trigger_action" {
		// 用户确认则进入第一个 guide
		if action.Value == "yes" {
			// 返回 trigger_cancel_card 中配置的确认卡片
			sessionCtx, _ := GetSessionCtx(action.UserId)
			process := def.GetProcesses()[sessionCtx.ProcessIndex]
			file, _ := json2.Marshal(process.Trigger.TriggerConfirmCard)
			ParseJson(&file, action.UserId)
			var card model.MessageCard
			err := json2.Unmarshal(file, &card)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, card)
		}
		// 用户取消则重置 sessionCtx
		if action.Value == "no" {
			// 返回 trigger_cancel_card 中配置的取消卡片
			sessionCtx, _ := GetSessionCtx(action.UserId)
			process := def.GetProcesses()[sessionCtx.ProcessIndex]
			file, _ := json2.Marshal(process.Trigger.TriggerCancelCard)
			ParseJson(&file, action.UserId)
			var card model.MessageCard
			err := json2.Unmarshal(file, &card)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, card)
			// 重置上下文，防止卡在状态机的中间，无法触发下一次的trigger
			ResetSessionCtx(action.UserId)
		}
	}
}
