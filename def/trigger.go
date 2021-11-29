package def

import "dsl/model"

// Trigger 触发器，当检测到关键词后，会进入对应 Process，并根据 GuideIndex 进入首个指引
type Trigger struct {
	Keywords           []string          `json:"keywords"`
	GuideIndex         int               `json:"guide_index"`
	TriggerCard        model.MessageCard `json:"trigger_card"`         // 触发后的提示卡片，用户确认后进入第一个guide
	TriggerCancelCard  model.MessageCard `json:"trigger_cancel_card"`  // 用户在提示卡片中取消后显示的卡片
	TriggerConfirmCard model.MessageCard `json:"trigger_confirm_card"` // 用户在提示卡片中取消后显示的卡片
}
