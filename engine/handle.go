package engine

import (
	"Tulan/def"
	handler2 "Tulan/handler"
	"fmt"
	"reflect"
)

func FindHandleByIndex(process def.Process, index int) (def.Handle, bool) {
	for _, handle := range process.Handles {
		if handle.Index == index {
			return handle, true
		}
	}
	return def.Handle{}, false
}

// DoHandle 通过传入上下文进行handle处理
func DoHandle(userId string, c SessionCtx) {
	// 首先判断当前是否处于handle状态
	if c.NowType != TYPE_HANDLE {
		return
	}

	// 然后找到对应handle
	process := def.GetProcesses()[c.ProcessIndex]
	handle, isFound := FindHandleByIndex(process, c.NowIndex)
	if !isFound {
		return
	}

	handlerName := handle.Handler

	// 然后通过反射找到 对应于 def.json 中与 handler 字段同名的 handler
	handler := handler2.GetHandler()
	handlerValue := reflect.ValueOf(handler)
	handlerMethod := handlerValue.MethodByName(handlerName)
	// 将def.json中的参数传到handler里，并调用
	var paras []reflect.Value
	for _, para := range handle.Params {
		in := ""
		// para为上下文中Params的键
		// 如果包含@@则将其替换为上下文中的变量
		if len(para) > 4 && para[:2] == "@@" && para[len(para)-2:] == "@@" {
			in = c.Params[para[2:len(para)-2]].(string)
		} else {
			in = para
		}
		paras = append(paras, reflect.ValueOf(in))
	}
	values := handlerMethod.Call(paras)

	// 将返回的结果存到上下文中
	for _, val := range values {
		key := fmt.Sprintf("handler_%v_value", c.NowIndex)
		c.Params[key] = val.String()
	}

	// next_guide_index 的优先级高于 guideIndex，判断进入下一个 handle 或 guide
	if handle.NextHandleIndex == 0 {
		// 将状态转移到下一个guide
		c.NowType = TYPE_GUIDE
		c.NowIndex = handle.SuccessGuideIndex
		// 更新上下文
		UpdateSessionCtx(userId, c)
		// 发送guide消息
		SendMessageGuide(userId, c.NowIndex)
	} else {
		c.NowIndex = handle.NextHandleIndex
		// 更新上下文
		UpdateSessionCtx(userId, c)
		// 继续调用 DoHandle
		DoHandle(userId, c)
	}
}
