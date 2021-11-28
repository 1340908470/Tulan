package main

import (
	"dsl/def"
	"dsl/engine"
	"dsl/handlerFunc"
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

	r.POST("/feishu/event", handlerFunc.EventHandlerFunc)

	err = r.Run(":8081") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}
