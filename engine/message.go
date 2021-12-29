// Package engine.message 是以接受者的角度，处理用户发来的消息
package engine

import (
	"Tulan/def"
	"Tulan/model"
	"Tulan/web"
	json2 "encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type MessageEvent struct {
	Sender  MessageEventSender  `json:"sender"`
	Message MessageEventMessage `json:"message"`
}

type MessageEventSender struct {
	SenderId   model.UserKey `json:"sender_id"`
	SenderType string        `json:"sender_type"`
	TenantKey  string        `json:"tenant_key"`
}

type MessageEventMessage struct {
	MessageId   string                       `json:"message_id"` // 消息的 open_message_id
	RootId      string                       `json:"root_id"`
	ParentId    string                       `json:"parent_id"`
	CreateTime  string                       `json:"create_time"`
	ChatId      string                       `json:"chat_id"`      // 消息所在的群组 id
	ChatType    string                       `json:"chat_type"`    // 消息所在的群组类型
	MessageType string                       `json:"message_type"` // 消息类型
	Content     string                       `json:"content"`      // 消息内容
	Mentions    []MessageEventMessageMention `json:"mentions"`
}

type MessageEventMessageContent struct {
	Text string `json:"text"`
}

// MessageEventMessageMention 被提及用户的信息
type MessageEventMessageMention struct {
	Key       string        `json:"key"`
	Id        model.UserKey `json:"id"`
	Name      string        `json:"name"`
	TenantKey string        `json:"tenant_key"`
}

// HandleMessageEvent 处理消息事件
func HandleMessageEvent(event map[string]interface{}) error {
	b, _ := json2.Marshal(event)
	var messageEvent MessageEvent
	err := json2.Unmarshal(b, &messageEvent)
	if err != nil {
		return errors.New("解析消息事件时出错：" + err.Error())
	}

	// TODO: 会话过期等导致等创建新的process应向用户发送一段提示消息

	// 根据 sender 获取对应的上下文
	sessionCtx, isNew := GetSessionCtx(messageEvent.Sender.SenderId.UserId)
	if isNew {
		var content MessageEventMessageContent
		json2.Unmarshal([]byte(messageEvent.Message.Content), &content)
		// 如果是新上下文，则应该：触发trigger - 找到process - 设置状态为wait - 给用户发送"触发事务"消息
		process, processIndex, isFound := FindProcess(content.Text)
		// 如果没有找到，则不应发送触发事务的消息；否则设置上下文状态
		if isFound {
			sessionCtx.NowType = TYPE_WAIT
			sessionCtx.ProcessName = process.Name
			sessionCtx.ProcessIndex = processIndex
			// 同时将processName加入到参数列表里
			sessionCtx.Params["process_name"] = process.Name
			sessionCtx.ChatId = messageEvent.Message.ChatId
			UpdateSessionCtx(messageEvent.Sender.SenderId.UserId, sessionCtx)
			SendMessageTrigger(messageEvent.Sender.SenderId.UserId)
		}
	} else {
		// 否则，如果当前上下文处于guide，则将消息内容作为 guide_<index>_response 的值存到 paras 中，并更新状态为 handle，进行处理
		if sessionCtx.NowType == TYPE_GUIDE {
			paraKey := fmt.Sprintf("guide_%v_response", sessionCtx.NowIndex)
			var content MessageEventMessageContent
			json2.Unmarshal([]byte(messageEvent.Message.Content), &content)
			sessionCtx.Params[paraKey] = content.Text

			sessionCtx.NowType = TYPE_HANDLE
			process := def.GetProcesses()[sessionCtx.ProcessIndex]
			guide, _ := FindGuideByIndex(process, sessionCtx.NowIndex)

			// 如果guide中regexp字段非空，则进行正则匹配，判断是否符合需求
			if guide.Regexp != "" {
				match, _ := regexp.MatchString(guide.Regexp, content.Text)
				if !match {
					// 如果不匹配，则
					SendMessageRegexpErr(messageEvent.Sender.SenderId.UserId)
					return nil
				}
			}

			sessionCtx.NowIndex = guide.SuccessHandleIndex
			UpdateSessionCtx(messageEvent.Sender.SenderId.UserId, sessionCtx)
			DoHandle(messageEvent.Sender.SenderId.UserId, sessionCtx)
		}
	}

	return err
}

// SendMessageTrigger 向用户发送触发事务的消息
func SendMessageTrigger(userId string) {
	sessionCtx, _ := GetSessionCtx(userId)

	process := def.GetProcesses()[sessionCtx.ProcessIndex]
	file, _ := json2.Marshal(process.Trigger.TriggerCard)
	ParseJson(&file, userId)

	var messageCard model.MessageCard
	err := json2.Unmarshal(file, &messageCard)
	if err != nil {
		return
	}

	SendMessage(messageCard, sessionCtx.ChatId)
}

// SendMessageRegexpErr 向用户校验失败的消息
func SendMessageRegexpErr(userId string) {
	sessionCtx, _ := GetSessionCtx(userId)

	process := def.GetProcesses()[sessionCtx.ProcessIndex]
	guide, _ := FindGuideByIndex(process, sessionCtx.NowIndex)
	file, _ := json2.Marshal(guide.RegexpErrCard)
	ParseJson(&file, userId)

	var messageCard model.MessageCard
	err := json2.Unmarshal(file, &messageCard)
	if err != nil {
		return
	}

	SendMessage(messageCard, sessionCtx.ChatId)
}

// SendMessageGuide 发送guide中定义的消息
func SendMessageGuide(userId string, guideIndex int) {
	sessionCtx, _ := GetSessionCtx(userId)

	process := def.GetProcesses()[sessionCtx.ProcessIndex]
	var guide def.Guide
	for _, g := range process.Guides {
		if g.Index == guideIndex {
			guide = g
			break
		}
	}
	file, _ := json2.Marshal(guide.GuideCard)
	ParseJson(&file, userId)

	var messageCard model.MessageCard
	err := json2.Unmarshal(file, &messageCard)
	if err != nil {
		return
	}

	SendMessage(messageCard, sessionCtx.ChatId)
}

func SendMessage(messageCard model.MessageCard, chatId string) {
	message := model.Message{
		ChatId:  chatId,
		MsgType: "interactive",
		Card:    messageCard,
	}

	paras := make(map[string]string)
	paras["receive_id_type"] = "user_id"

	json := web.Request(web.ApiSendMessageCard, message, paras)
	var res web.ApiSendMessageCardRes
	json2.Unmarshal(json, &res)

	json, err := json2.Marshal(message)
	if err != nil {
		return
	}
}
