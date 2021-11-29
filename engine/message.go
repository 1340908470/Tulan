// Package engine.message 是以接受者的角度，处理用户发来的消息
package engine

import (
	"dsl/def"
	"dsl/model"
	"dsl/web"
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

	// TODO: 通过engine/trigger提供等函数，解析消息并得到对应的process

	// 根据 sender 获取对应的上下文
	sessionCtx, isNew := GetSessionCtx(messageEvent.Sender.SenderId.UserId)
	if isNew {
		// 如果是新上下文，则应该：触发trigger - 找到process - 设置状态为wait - 给用户发送"触发事务"消息
		// TODO：处理用户 "触发事务"消息 的响应，接受则更新上下文，并进入首个guide；否则不进行任何操作；更新卡片内容
		process, processIndex, isFound := FindProcess(messageEvent.Message.Content)
		// 如果没有找到，则不应发送触发事务的消息；否则设置上下文状态
		if isFound {
			sessionCtx.NowType = WAIT
			sessionCtx.ProcessName = process.Name
			sessionCtx.ProcessIndex = processIndex
			// 同时将processName加入到参数列表里
			sessionCtx.Params["process_name"] = process.Name
			UpdateSessionCtx(messageEvent.Sender.SenderId.UserId, sessionCtx)
			SendMessageTrigger(messageEvent)
		}
	} else {
		// 否则，更新状态为 handle 并将消息作为参数传递给 handler
	}

	return err
}

// SendMessageTrigger 向用户发送触发事务的消息
func SendMessageTrigger(messageEvent MessageEvent) {
	sessionCtx, _ := GetSessionCtx(messageEvent.Sender.SenderId.UserId)
	print(sessionCtx.Params["process_name"])

	process := def.GetProcesses()[sessionCtx.ProcessIndex]
	file, _ := json2.Marshal(process.Trigger.TriggerCard)
	ParseJson(&file, messageEvent.Sender.SenderId.UserId)
	var messageCard model.MessageCard
	err := json2.Unmarshal(file, &messageCard)
	if err != nil {
		return
	}

	message := model.Message{
		ChatId:  messageEvent.Message.ChatId,
		MsgType: "interactive",
		Card:    messageCard,
	}

	paras := make(map[string]string)
	paras["receive_id_type"] = "user_id"

	json := web.Request(web.ApiSendMessageCard, message, paras)
	var res web.ApiSendMessageCardRes
	json2.Unmarshal(json, &res)

	json, err = json2.Marshal(message)
	if err != nil {
		return
	}
}

func ParseJson(json *[]byte, userId string) {
	str := string(*json)
	r, _ := regexp.Compile("@@[\\s\\S]*?@@")
	indexes := r.FindAllIndex(*json, -1)
	sessionCtx, _ := GetSessionCtx(userId)
	for _, index := range indexes {
		for key, value := range sessionCtx.Params {
			if index[1]-index[0] > 3 && str[index[0]+2:index[1]-2] == key {
				str = fmt.Sprintf("%v%v%v", str[:index[0]], value, str[index[1]:])
			}
		}
	}
	*json = []byte(str)
}
