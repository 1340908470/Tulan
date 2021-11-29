// Package engine.trigger 以接受者的角度，解析消息并得到对应的process
package engine

import (
	"dsl/def"
	"strings"
)

// FindProcess 根据用户发送的消息内容，找到符合def.json中配置的process 以及 是否找到
func FindProcess(msgContent string) (def.Process, int, bool) {
	processes := def.GetProcesses()
	for index, process := range processes {
		trigger := process.Trigger
		for _, keyword := range trigger.Keywords {
			// 如果包含关键词，则将该 process 返回
			if strings.Contains(msgContent, keyword) {
				return process, index, true
			}
		}
	}
	return def.Process{}, 0, false
}
