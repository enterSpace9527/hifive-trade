package db_model

import (
	"github.com/shopspring/decimal"
	"time"
)

type UserFunds struct {
	ID           int64
	UID          int64
	Address      string
	CoinName     string          `gorm:"column:coin_name"`
	TotalAmount  decimal.Decimal `gorm:"column:total_amount"`
	Amount       decimal.Decimal `gorm:"column:amount"`
	FreezeAmount decimal.Decimal `gorm:"column:freeze_amount"`
	CreateTime   time.Time       `gorm:"column:create_time"`
	UpdateTime   time.Time       `gorm:"column:update_time"`
}

// TableName 设置模型对应的数据库表名
func (UserDetail) user_funds() string {
	return "user_funds"
}
