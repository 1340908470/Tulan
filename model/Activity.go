package model

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct {
	Name     string    // 活动名称
	Describe string    // 活动介绍
	Time     time.Time // 活动时间
	Location string    // 活动地点
}

// CreateActivity 创建活动
func CreateActivity(db *gorm.DB, activity Activity) {

}
