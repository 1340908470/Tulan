package engine

import (
	"strconv"
	"time"
)

// DispatchCenter 调度中心，用于调度每个人与机器人的会话
type DispatchCenter struct {
	Session map[string]SessionCtx // 以 user_id 为键，为每一个用户维护一个会话上下文
}

// SessionCtx 会话上下文
type SessionCtx struct {
	LastTime    string            // 用户上一次操作的时间，格式为时间戳，如 1638018223137，如果当前时间已过30分钟，则会重置
	ProcessName string            // 当前会话进入了哪一个 process
	Params      map[string]string // 当前会话的参数列表
	NowType     string            // 当前在 process 中的类型："guide" | "handle"
	NowIndex    int               // 当前处于 guides[NowIndex] | handles[NowIndex]
}

var dispatchCenter DispatchCenter

func InitDispatchCenter() {
	dispatchCenter.Session = make(map[string]SessionCtx)
}

// GetSessionCtx 根据用户id获取会话上下文
func GetSessionCtx(userId string) SessionCtx {
	sessionCtx, ok := dispatchCenter.Session[userId]
	// 如果还没有创建过会话 或者 已经过期了，则会根据message的内容，在session中更新ctx
	now := time.Now().UnixNano() / 1e6
	lastTime, _ := strconv.ParseInt(sessionCtx.LastTime, 10, 64)
	if !ok || now-lastTime > 30*60*1000 {
		// TODO: 会话过期等导致等创建新的process应向用户发送一段提示消息

		// TODO: 通过engine/trigger提供等函数，解析消息并得到对应的process

		dispatchCenter.Session[userId] = SessionCtx{
			LastTime:    strconv.FormatInt(now, 10),
			ProcessName: "",
			Params:      nil,
			NowType:     "",
			NowIndex:    0,
		}
	}
}
