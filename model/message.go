// Package model.message 是以发送者的角度，向用户发送消息
package model

// Message 消息
type Message struct {
	ChatId  string      `json:"chat_id"`
	MsgType string      `json:"msg_type"`
	RootId  string      `json:"root_id"`
	Card    MessageCard `json:"card"`
}

// MessageCard 消息卡片
type MessageCard struct {
	Config   MessageCardConfig `json:"config"`
	Header   MessageCardHeader `json:"header"`
	Elements []MessageElement  `json:"elements"`
}

// MessageCardConfig 描述消息卡片的功能属性
type MessageCardConfig struct {
	UpdateMulti    bool `json:"update_multi"`
	WideScreenMode bool `json:"wide_screen_mode"`
}

// MessageCardHeader 用于配置卡片的标题内容
type MessageCardHeader struct {
	Title MessageCardHeaderTitle `json:"title"`
}

type MessageCardHeaderTitle struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

// MessageElement 用于定义卡片的正文内容
// 消息卡片的正文内容由 MessageElement 堆砌而成
// 	- 内容模块
//  - 图片模块
//  - 交互模块
type MessageElement struct {
	// 根据Tag选择是哪一种模块
	Tag string `json:"tag"`

	// 内容模块 Tag = "div"
	Text   MessageContentText    `json:"text"` // Text 和 Fields 二选一
	Fields []MessageElementField `json:"fields"`
	Extra  MessageElementExtra   `json:"extra"` // 可选，附加的元素展示在文本内容右侧

	// 交互模块 Tag = "action"

	// TODO：还有图片模块、交互模块没写
}

// MessageElementField 内容模块的多文本展示
type MessageElementField struct {
	IsShort bool               `json:"is_short"` // 是否并排布局
	Text    MessageContentText `json:"text"`
}

// MessageElementExtra 附加元素，图片、按钮等元素是 n选1 的关系
type MessageElementExtra struct {
	// 图片元素
	MessageContentImage

	// 按钮元素
	MessageContentButton

	// 日期选择
	MessageContentDatePicker

	// TODO：还有 selectMenu、overflow、datePicker 几种元素没写
}

// MessageContentText 文本元素
type MessageContentText struct {
	Tag     string `json:"tag"`     // 文本的内容类型，可以是 "plain_text" 或 "lark_md"
	Content string `json:"content"` // 文本内容
	Lines   int    `json:"lines"`   // 可选参数，用于设置内容显示行数
}

// MessageContentImage 图片元素
type MessageContentImage struct {
	Tag     string             `json:"tag"` // 固定: Tag = "img"
	ImgKey  string             `json:"img_key"`
	Alt     MessageContentText `json:"alt"`     // 图片 hover 说明
	Preview bool               `json:"preview"` // 可选参数，点击后是否放大图片，默认为 true
}

// MessageElementAction 交互模块交互组件，可以是 button 或 datePicker 或 overflow 或 selectMenu
type MessageElementAction struct {
	// 按钮
	MessageContentButton

	// 日期选择
	MessageContentDatePicker
}

// MessageContentButton 按钮元素，可用于内容块的 extra 字段和交互块的 actions 字段
// TODO: 目前还不能使用多端跳转链接
type MessageContentButton struct {
	Tag      string                    `json:"tag"`       // 固定: Tag = "button"
	Text     MessageContentText        `json:"text"`      // 指定按钮的文本
	Url      string                    `json:"url"`       // 可选参数，跳转链接，与 multi_url 互斥
	MultiUrl MessageContentUrl         `json:"multi_url"` // 可选参数，多端跳转链接
	Type     string                    `json:"type"`      // 可选参数，按钮样式，默认为 "default" 可以是 "default" 或 "primary" 或 "danger"
	Value    MessageContentButtonValue `json:"value"`     // 用户在点击按钮后，向 Process 中添加的键值
}

type MessageContentUrl struct {
}

// MessageContentButtonValue 点击按钮后返回一对键值
type MessageContentButtonValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// MessageContentDatePicker 时间选择组件
type MessageContentDatePicker struct {
	Tag string `json:"tag"` // 三选一: Tag = "date_picker" ｜ "picker_time" ｜ "picker_datetime" 【日期｜时间｜日期+时间】

	InitialDate     string `json:"initial_date"`     // 可选参数，日期模式下的初始值，格式："yyyy-MM-dd"
	InitialTime     string `json:"initial_time"`     // 可选参数，时间模式下的初始值，格式："HH:mm"
	InitialDatetime string `json:"initial_datetime"` // 可选参数，日期时间模式的初始值，	格式："yyyy-MM-dd HH:mm"

	Value MessageContentDatePickerValue `json:"value"` // 用户选择后，回传的参数
}

type MessageContentDatePickerValue struct {
}
