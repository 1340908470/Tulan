package main

import (
	"Tulan/def"
	"Tulan/engine"
	"Tulan/handler"
	"Tulan/handlerFunc"
	"Tulan/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// 初始化def，读入def.json
	err := def.InitDef()
	if err != nil {
		return
	}

	// 初始化dispatchCenter
	engine.InitDispatchCenter()

	// 初始化db
	db, err := gorm.Open(sqlite.Open("tulan.db"), &gorm.Config{})
	if err != nil {
		panic("数据库初始化失败" + err.Error())
	}

	// 初始化Model(迁移)
	err = model.InitModel(db)
	if err != nil {
		panic(err)
	}

	// 初始化handler
	handler.InitHandler(db)

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
