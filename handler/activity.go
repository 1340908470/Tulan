package handler

import (
	"Tulan/model"
	time2 "time"
)

func (h Handler) CreateActivity(name string, describe string, timeStr string, location string) {
	layout := "2006-01-02 15:04:05"
	time, err := time2.Parse(layout, timeStr)
	if err != nil {
		println("err: 创建活动时，传入的日期格式错误")
	}
	activity := model.Activity{
		Name:     name,
		Describe: describe,
		Time:     time,
		Location: location,
	}
	err = model.CreateActivity(h.db, &activity)
	if err != nil {
		println(err.Error())
	}
}
