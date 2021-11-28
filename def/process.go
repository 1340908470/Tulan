package def

// Process 一个客服机器人是由若干个 Process 构成的
//
// 一个典型的例子：【通过客服机器人预定会议】
//
// 首先，通过 Trigger 定义触发器，其包括触发关键词组：["预定会议"， "预约会议"]。
// 然后，在首次触发后，机器人会进入 Guide - Handle - Guide - Handle ... 的步骤直至结束
// 	- Guide 是以消息卡片的形式进行用户指引，在引导下，用户会执行一些操作， Guide 可以包含富文本、图片、用户交互的按钮、超链接等
//	如回复文字、点击按钮等，此阶段用户执行的操作会以键值的形式存储下来，可作为链中后面的 Guide 或 Handle 的入参
//  - Handle 则是在后端进行一些逻辑操作，其总体上可分为两类，一是对于已有 model 的基本操作，二是针对机器人行为的调用(api)
// 那么，在这个例子中，可以是 G(指引用户选择日期) - G(指引用户选择人员) - H(创建会议) - G(询问是否需要提醒) - H(批量发送消息)
type Process struct {
	Name    string   `json:"name"`
	Trigger Trigger  `json:"trigger"`
	Guides  []Guide  `json:"guides"`
	Handles []Handle `json:"handles"`
}
