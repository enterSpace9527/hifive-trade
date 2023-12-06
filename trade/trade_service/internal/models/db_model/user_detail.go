package db_model

import (
	"time"
)

type UserDetail struct {
	ID         int64
	Address    string
	IsKol      uint16 `gorm:"column:is_kol"`
	Role       uint16
	DeviceId   string `gorm:"column:device_id"`
	DeviceNo   string `gorm:"column:device_no"`
	DeviceIp   string `gorm:"column:device_ip"`
	Status     uint16
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

// TableName 设置模型对应的数据库表名
func (UserDetail) TableName() string {
	return "user_detail"
}
