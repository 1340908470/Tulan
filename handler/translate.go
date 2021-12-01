package handler

import (
	"Tulan/web"
	json2 "encoding/json"
)

// TranslateFromZhToEn 将中文翻译为英文，输入参数为要翻译的源中文文字。
// 翻译完成后，会将翻译结果保存到上下文中，在 success_guide_index 指向的 guide 中，
// 之后，可以由用户自定义消息卡片显示翻译结果，或者在之后的 handler 或者 guide 中使用此次的翻译结果
func (h Handler) TranslateFromZhToEn(text string) string {
	resJson := web.Request(web.ApiTranslate, web.ApiTranslateReq{
		SourceLanguage: "zh",
		Text:           text,
		TargetLanguage: "en",
	}, nil)
	var res web.ApiTranslateRes
	err := json2.Unmarshal(resJson, &res)
	if err != nil || res.Msg != "Success" {
		return ""
	} else {
		return res.Data.Text
	}
}
