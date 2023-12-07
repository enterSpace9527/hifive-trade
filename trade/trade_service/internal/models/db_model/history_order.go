package db_model

import (
	"github.com/shopspring/decimal"
	"time"
)

type HistoryOrder struct {
	ID         uint64
	Symbol     string
	Side       string
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	Price      decimal.Decimal
	Quantity   decimal.Decimal
	UID        int64
	Status     string
	Type       string
	OID        int64           `gorm:"column:oid"`
	CID        string          `gorm:"column:cid"`
	SID        string          `gorm:"column:sid"`
	Fee        decimal.Decimal `gorm:"column:fee"`
	FeeSymbol  decimal.Decimal `gorm:"column:fee_symbol"`
}

// TableName 设置模型对应的数据库表名
func (HistoryOrder) TableName() string {
	return "history_orders"
}
