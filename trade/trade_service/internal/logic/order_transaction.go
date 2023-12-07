package logic

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/gvar"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/db_model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)



func orderFilledTransaction(orderResponse *binance.CreateOrderResponse) {
	gvar.PostgresClient.Transaction(func(tx *gorm.DB) error {

		//插入总单
		hstoryOrder := db_model.HistoryOrder{}
		hstoryOrder.Status = string(orderResponse.Status)
		hstoryOrder.CreateTime = time.Now().UTC()
		hstoryOrder.UpdateTime = hstoryOrder.CreateTime
		hstoryOrder.Symbol = orderResponse.Symbol
		hstoryOrder.Type = string(orderResponse.Type)
		hstoryOrder.Side = string(orderResponse.Side)
		hstoryOrder.Quantity = orderResponse.ExecutedQuantity
		hstoryOrder.Price = orderResponse.Price //可能不正确
		hstoryOrder.OID = orderResponse.OrderID
		hstoryOrder.CID = orderResponse.ClientOrderID
		hstoryOrder.Fee =
		hstoryOrder.FeeSymbol =

		//插入用户单
		hstoryOrder.SID = fmt.Sprintf("%v-1", orderResponse.ClientOrderID)

		hstoryOrder.Price =
		hstoryOrder.Price, err := decimal.NewFromString()
		if err := tx.Create(&hstoryOrder).Error; err != nil {
			return err
		}

		if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	}
}
