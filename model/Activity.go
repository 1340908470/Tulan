package model

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	gorm.Model
	Name     string    // 活动名称
	Describe string    // 活动介绍
	Time     time.Time // 活动时间
	Location string    // 活动地点
}

// CreateActivity 创建活动
func CreateActivity(db *gorm.DB, activity *Activity) error {
	err := db.Create(activity).Error
	if err != nil {
		return errors.New("创建活动时出现错误：" + err.Error())
	}
	return nil
}
