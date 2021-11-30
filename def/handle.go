package def

// Handle 处理函数
type Handle struct {
	Index   int    `json:"index"`
	Handler string `json:"handler"`

	// 键为对应handler的参数名，值为对应参数列表中的键，如：handler中需要一个参数name，
	// 而这个值由index为1的guide下的用户回应构成，则应写为：["guide_<index>_response"]
	Params            []string `json:"params"`
	SuccessGuideIndex int      `json:"success_guide_index"`
	FailedGuideIndex  int      `json:"failed_guide_index"`

	// 下一个 handle 的索引，该字段的优先级大于 guideIndex，即当该字段非零时，会接着执行下一个handle，只有当该字段为0时，才会进入指引
	NextHandleIndex int `json:"next_guide_index"`
}
