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
	var eventMessage MessageEvent
	err := json2.Unmarshal(b, &eventMessage)
	if err != nil {
		return errors.New("解析消息事件时出错：" + err.Error())
	}

	return err
}
