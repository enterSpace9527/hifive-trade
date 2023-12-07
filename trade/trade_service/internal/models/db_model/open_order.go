package db_model

import (
	"github.com/shopspring/decimal"
	"time"
)

type OpenOrder struct {
	ID           uint64
	Symbol       string
	CID          string          `gorm:"column:cid"`
	FreezeAmount decimal.Decimal `gorm:"column:freeze_amount"`

	UID          int64
	Type         string
	Side         string
	Price        decimal.Decimal
	Quantity     decimal.Decimal
	DonePrice    decimal.Decimal `gorm:"column:done_price"`
	DoneQuantity decimal.Decimal `gorm:"column:done_quantity"`
	CreateTime   time.Time       `gorm:"column:create_time"`
}

// TableName 设置模型对应的数据库表名
func (OpenOrder) TableName() string {
	return "open_orders"
}
