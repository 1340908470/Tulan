package engine

import (
	"Tulan/def"
)

// FindGuideByIndex 通过索引找到guide, 以及是否找到
func FindGuideByIndex(process def.Process, index int) (def.Guide, bool) {
	for _, guide := range process.Guides {
		if guide.Index == index {
			return guide, true
		}
	}
	return def.Guide{}, false
}
