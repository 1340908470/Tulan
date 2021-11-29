package engine

import (
	"fmt"
	"regexp"
)

func ParseJson(json *[]byte, userId string) {
	str := string(*json)
	r, _ := regexp.Compile("@@[\\s\\S]*?@@")
	indexes := r.FindAllIndex(*json, -1)
	sessionCtx, _ := GetSessionCtx(userId)
	for _, index := range indexes {
		for key, value := range sessionCtx.Params {
			if index[1]-index[0] > 3 && str[index[0]+2:index[1]-2] == key {
				str = fmt.Sprintf("%v%v%v", str[:index[0]], value, str[index[1]:])
			}
		}
	}
	*json = []byte(str)
}
