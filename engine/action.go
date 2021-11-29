package engine

import (
	"dsl/def"
	"dsl/model"
	json2 "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

			// 然后触发第一个guide
			// 更新上下文
			sessionCtx.NowType = TYPE_GUIDE
			sessionCtx.NowIndex = process.Trigger.GuideIndex
			UpdateSessionCtx(action.UserId, sessionCtx)
			SendMessageGuide(action.UserId, process.Trigger.GuideIndex)
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

	// 如果字段为 move_to_guide: xxx ，则将当前卡片替换为索引为xxx的卡片
	if action.Key == "move_to_guide" {
		sessionCtx, _ := GetSessionCtx(action.UserId)
		index, _ := strconv.ParseInt(action.Value, 10, 64)
		// 更新状态到 guide xxx
		sessionCtx.NowIndex = int(index)
		UpdateSessionCtx(action.UserId, sessionCtx)
		SendMessageGuide(action.UserId, int(index))
	}
}
