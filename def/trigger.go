package def

import "dsl/model"

// Trigger 触发器，当检测到关键词后，会进入对应 Process，并根据 GuideIndex 进入首个指引
type Trigger struct {
	Keywords    []string          `json:"keywords"`
	GuideIndex  int               `json:"guide_index"`
	TriggerCard model.MessageCard `json:"trigger_card"`
}
