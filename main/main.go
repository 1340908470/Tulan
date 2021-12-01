package main

import (
	"Tulan/def"
	"Tulan/engine"
	"Tulan/handlerFunc"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化def，读入def.json
	err := def.InitDef()
	if err != nil {
		return
	}

	// 初始化dispatchCenter
	engine.InitDispatchCenter()

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	// 处理订阅等事件，如接收到用户发来的消息
	r.POST("/feishu/event", handlerFunc.EventHandlerFunc)

	// 处理交互，如用户点击消息卡片中的按钮
	r.POST("feishu/action", handlerFunc.ActionHandlerFunc)

	err = r.Run(":8081") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
