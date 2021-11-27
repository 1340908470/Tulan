package def

// Handle 处理函数
type Handle struct {
	Index             int      `json:"index"`
	Handler           string   `json:"handler"`
	HandlerParas      []string `json:"handler_paras"`
	ValueName         string   `json:"value_name"`
	SuccessGuideIndex int      `json:"success_guide_index"`
	FailedGuideIndex  int      `json:"failed_guide_index"`
}
