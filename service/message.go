package service

import (
	"dsl/engine"
	"dsl/model"
)

// SendMessageTrigger 向用户发送触发事务的消息
func SendMessageTrigger(messageEvent engine.MessageEvent) {
	message := model.Message{
		ChatId:  "",
		MsgType: "",
		RootId:  "",
		Card:    model.MessageCard{},
	}
}
