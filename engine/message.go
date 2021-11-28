// Package engine.message 是以接受者的角度，处理用户发来的消息
package engine

import (
	"dsl/model"
	json2 "encoding/json"
	"errors"
)

type MessageEvent struct {
	Sender  model.UserKey       `json:"sender"`
	Message MessageEventMessage `json:"message"`
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
	sessionCtx, isNew := GetSessionCtx(messageEvent.Sender.UserId)
	if isNew {
		// 如果是新上下文，则应该：触发trigger - 找到process - 设置状态为wait - 给用户发送"触发事务"消息
		// TODO：处理用户 "触发事务"消息 的响应，接受则更新上下文，并进入首个guide；否则不进行任何操作；更新卡片内容
		process, isFound := FindProcess(messageEvent.Message.Content)
		// 如果没有找到，则不应发送触发事务的消息；否则设置上下文状态
		if isFound {
			sessionCtx.NowType = WAIT
			sessionCtx.ProcessName = process.Name
			UpdateSessionCtx(messageEvent.Sender.UserId, sessionCtx)
			// TODO: 给用户发送"触发事务"消息
		}
	} else {
		// 否则，更新状态为 handle 并将消息作为参数传递给 handler
	}

	return err
}
