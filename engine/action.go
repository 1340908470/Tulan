package engine

type Action struct {
	OpenId        string `json:"open_id"`
	UserId        string `json:"user_id"`
	OpenMessageId string `json:"open_message_id"`
	TenantKey     string `json:"tenant_key"`
	Token         string `json:"token"`
	Key           string `json:"key"`
	Value         string `json:"value"`
}

func HandleAction(action Action) {
	// 首先判断是否是 trigger_action，并依此判断是否进入事务
	if action.Key == "trigger_action" {
		// 用户确认则进入第一个guide
		if action.Value == "yes" {

		}
		// 用户取消则重置sessionCtx
		if action.Value == "no" {
			ResetSessionCtx(action.UserId)
		}
	}
}
