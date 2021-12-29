package def

import "Tulan/model"

// Guide 指引，以消息卡片的形式，指引用户进行下一步操作
type Guide struct {
	Index              int               `json:"index"` // 指引的索引
	SuccessHandleIndex int               `json:"success_handle_index"`
	GuideCard          model.MessageCard `json:"guide_card"`
	Regexp             string            `json:"regexp"` // 可选字段，表示指引后用户输入内容的正则表达式
	RegexpErrCard      model.MessageCard `json:"regexp_err_card"`
}
