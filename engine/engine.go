package engine

import (
	"strconv"
	"time"
)

// DispatchCenter 调度中心，用于调度每个人与机器人的会话
type DispatchCenter struct {
	Session map[string]SessionCtx // 以 user_id 为键，为每一个用户维护一个会话上下文
}

var GUIDE = "guide"
var HANDLE = "handle"
var WAIT = "wait"

// SessionCtx 会话上下文
type SessionCtx struct {
	LastTime     string            // 用户上一次操作的时间，格式为时间戳，如 1638018223137，如果当前时间已过30分钟，则会重置
	ProcessName  string            // 当前会话进入了哪一个 process
	ProcessIndex int               // 当前会话的 process 的索引
	Params       map[string]string // 当前会话的参数列表
	NowType      string            // 当前在 process 中的类型："guide" | "handle" ｜ "wait"，wait为等待用户接受开始事务
	NowIndex     int               // 当前处于 guides[NowIndex] | handles[NowIndex]
}

var dispatchCenter DispatchCenter

// ResetSessionCtx 重置某个用户的上下文，只需要修改LastTime让其过期即可
func ResetSessionCtx(userId string) {
	sessionCtx, ok := dispatchCenter.Session[userId]
	if ok {
		sessionCtx = SessionCtx{
			LastTime:    "0",
			ProcessName: "",
			Params:      make(map[string]string),
			NowType:     "",
			NowIndex:    0,
		}
		dispatchCenter.Session[userId] = sessionCtx
	}
}

func InitDispatchCenter() {
	dispatchCenter.Session = make(map[string]SessionCtx)
}

// GetSessionCtx 根据用户id获取会话上下文, 返回上下文以及是否是新上下文
func GetSessionCtx(userId string) (SessionCtx, bool) {
	isNew := false
	sessionCtx, ok := dispatchCenter.Session[userId]
	// 如果还没有创建过会话 或者 已经过期了，则会根据message的内容，在session中更新ctx
	now := time.Now().UnixNano() / 1e6
	lastTime, _ := strconv.ParseInt(sessionCtx.LastTime, 10, 64)
	if !ok || now-lastTime > 30*60*1000 {
		isNew = true
		sessionCtx = SessionCtx{
			LastTime:    strconv.FormatInt(now, 10),
			ProcessName: "",
			Params:      make(map[string]string),
			NowType:     "",
			NowIndex:    0,
		}
		dispatchCenter.Session[userId] = sessionCtx
	}

	return sessionCtx, isNew
}

// UpdateSessionCtx 更新上下文，每次用户进行完操作后，都会更新一次上下文; 函数会更新入参的非空字段
func UpdateSessionCtx(userId string, sessionCtx SessionCtx) {
	processName := dispatchCenter.Session[userId].ProcessName
	processIndex := dispatchCenter.Session[userId].ProcessIndex
	params := dispatchCenter.Session[userId].Params
	nowType := dispatchCenter.Session[userId].NowType
	nowIndex := dispatchCenter.Session[userId].NowIndex

	if sessionCtx.ProcessName != "" {
		processName = sessionCtx.ProcessName
	}
	if sessionCtx.ProcessIndex != 0 {
		processIndex = sessionCtx.ProcessIndex
	}
	for key, val := range sessionCtx.Params {
		params[key] = val
	}
	if sessionCtx.NowType != "" {
		nowType = sessionCtx.NowType
	}
	if sessionCtx.NowIndex != 0 {
		nowIndex = sessionCtx.NowIndex
	}

	dispatchCenter.Session[userId] = SessionCtx{
		LastTime:     dispatchCenter.Session[userId].LastTime,
		ProcessName:  processName,
		ProcessIndex: processIndex,
		Params:       params,
		NowType:      nowType,
		NowIndex:     nowIndex,
	}
}
