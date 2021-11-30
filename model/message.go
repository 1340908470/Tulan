// Package model.message 是以发送者的角度，向用户发送消息
package model

// Message 消息
type Message struct {
	ChatId  string      `json:"chat_id,omitempty"`
	MsgType string      `json:"msg_type,omitempty"`
	RootId  string      `json:"root_id,omitempty"`
	Card    MessageCard `json:"card,omitempty"`
}

// MessageCard 消息卡片
type MessageCard struct {
	Config   map[string]interface{} `json:"config,omitempty"`
	Header   map[string]interface{} `json:"header,omitempty"`
	Elements []interface{}          `json:"elements,omitempty"`
}
