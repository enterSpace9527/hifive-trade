package logic

import (
	"context"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/gvar"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/db_model"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update() (resp *types.CommResp, err error) {

	/* update order test */
	//order := db_model.Order{}
	//gvar.PostgresClient.First(&order, 1)
	//order.CreateTime = time.Now().UTC()
	//order.UpdateTime = time.Now().UTC()
	//fmt.Println("order.Price:", order.Price)
	//order.Price = order.Price.Add(decimal.NewFromFloat(2.2))
	//result := gvar.PostgresClient.Save(&order)
	//
	//if result.Error != nil {
	//	logc.Infof(context.Background(), "result.Error: %v", result.Error.Error())
	//	return
	//}
	//logc.Infof(context.Background(), "update")

	/* update UserFunds test */
	UserFunds := db_model.UserFunds{}
	UserFunds.UID = 8527
	UserFunds.Address = "0xc7e1dca818ec9d400aa78a376803a55ad03f4422"
	UserFunds.CoinName = "USDT"
	UserFunds.Amount = decimal.NewFromFloat(1.1)
	UserFunds.TotalAmount = decimal.NewFromFloat(1.1)
	UserFunds.FreezeAmount = decimal.NewFromFloat(0)
	UserFunds.CreateTime = time.Now().UTC()
	UserFunds.UpdateTime = UserFunds.CreateTime

	result := gvar.PostgresClient.Create(&UserFunds)

	if result.Error != nil {
		logc.Infof(context.Background(), "result.Error: %v", result.Error.Error())
		return
	}
	logc.Infof(context.Background(), "Create")
	return
}
