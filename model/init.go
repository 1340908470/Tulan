package model

import (
	"errors"
	"gorm.io/gorm"
)

func InitModel(db *gorm.DB) error {
	err := db.AutoMigrate(&Activity{})
	if err != nil {
		return errors.New("数据库迁移失败: " + err.Error())
	}
	return err
}
