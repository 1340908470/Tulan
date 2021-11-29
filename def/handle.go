package def

// Handle 处理函数
type Handle struct {
	Index             int    `json:"index"`
	Handler           string `json:"handler"`
	SuccessGuideIndex int    `json:"success_guide_index"`
	FailedGuideIndex  int    `json:"failed_guide_index"`
}
