package logic

import (
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/gvar"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/db_model"
	"github.com/shopspring/decimal"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logc"
	"sync"
	"time"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

//todo_wr:币安客户端并发请求问题,暂时记录一下

type OrderLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	flake      *sonyflake.Sonyflake
	orderMutex sync.Mutex
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		flake:  sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

func (l *OrderLogic) Order(req *types.OrderRequest) (resp *types.OrderResponse, err error) {

	commResp := types.CommResp{
		Code:    0,
		Message: "success",
	}

	resp = &types.OrderResponse{
		CommResp: commResp,
	}

	err = l.CheckOrderParams(req)
	if err != nil {
		commResp.Code = 100
		commResp.Message = err.Error()
		return resp, err
	}

	orderInfo, err := l.BianOrder(req)
	if err != nil {
		commResp.Code = 100
		commResp.Message = err.Error()
		return resp, nil
	}

	l.UpdateResp(resp, orderInfo)
	return resp, err
}

func (l *OrderLogic) CheckOrderParams(req *types.OrderRequest) error {
	tradeType := gvar.TradeType(req.Type)
	tradeSide := gvar.TradeSide(req.Side)

	if (tradeSide == gvar.TradeSideBuy || tradeSide == gvar.TradeSideSell) != true {
		return fmt.Errorf("order params incorrect")
	}

	if tradeType == gvar.TradeTypeMarket {
		if req.Quantity != "" || req.QuoteOrderQty != "" {
			return nil
		}
	} else if tradeType == gvar.TradeTypeLimit {
		if req.Quantity != "" && req.Price != "" {
			return nil
		}
	}

	return fmt.Errorf("order params incorrect")
}

func (l *OrderLogic) BianOrder(req *types.OrderRequest) (*binance.CreateOrderResponse, error) {

	l.orderMutex.Lock()
	defer l.orderMutex.Unlock()

	err := l.CheckOrderParams(req)
	if err != nil {
		return nil, err
	}

	err = l.reqAdjustment(req)
	if err != nil {
		return nil, err
	}

	cli := gvar.BinanceClient
	createOrder := cli.NewCreateOrderService()

	id, err := l.flake.NextID()
	if err != nil {
		return nil, err
	}

	createOrder.NewClientOrderID(fmt.Sprintf("C%v", id))
	createOrder.Symbol(req.Symbol)

	tradeType := gvar.TradeType(req.Type)
	if tradeType == gvar.TradeTypeLimit {
		createOrder.Price(req.Price)
		createOrder.Quantity(req.Quantity)
		createOrder.TimeInForce("GTC")
	} else {
		if req.QuoteOrderQty != "" {
			createOrder.QuoteOrderQty(req.QuoteOrderQty)
		} else {
			createOrder.Quantity(req.Quantity)
		}
	}

	createOrder.Type(binance.OrderType(req.Type))
	createOrder.Side(binance.SideType(req.Side))

	ctx := context.Background()
	orderInfo, err := createOrder.Do(ctx)
	return orderInfo, err
}

func (l *OrderLogic) reqAdjustment(req *types.OrderRequest) error {
	tradeType := gvar.TradeType(req.Type)
	if tradeType == gvar.TradeTypeMarket {
		return l.marketTradeReqAdjustment(req)
	} else {
		return l.limitTradeReqAdjustment(req)
	}
}

func (l *OrderLogic) marketTradeReqAdjustment(req *types.OrderRequest) error {
	marketSymbol, err := gvar.GetMarketSymbol(req.Symbol)
	if err != nil {
		return err
	}

	userFunds := db_model.UserFunds{}
	tradeType := gvar.TradeType(req.Type)
	tradeSide := gvar.TradeSide(req.Side)
	if tradeType == gvar.TradeTypeMarket && tradeSide == gvar.TradeSideBuy {
		err = l.GetUserFunds(req.Uid, marketSymbol.BaseSymbol, &userFunds)
		//市价按量单处理买
		if req.Quantity != "" {
			bookTicker, err := gvar.GetBookTicker(req.Symbol)
			if err != nil {
				return err
			}

			currentTime := time.Now()
			timeDifference := currentTime.Sub(bookTicker.UpdateTime)
			if timeDifference > 5*time.Minute {
				return fmt.Errorf("bookTicker update time timeout")
			}

			quantityDecimal, err := decimal.NewFromString(req.Quantity)
			if err != nil {
				return err
			}

			bestAskPriceDecimal, err := decimal.NewFromString(bookTicker.BestAskPrice)
			if err != nil {
				return err
			}

			totalDecimal := quantityDecimal.Mul(bestAskPriceDecimal)
			b := l.isMarketPriceOrder(totalDecimal, userFunds.Amount)
			if b {
				req.Quantity = ""
				req.QuoteOrderQty = userFunds.Amount.String()
			}

			//市价按价单处理买
		} else {
			priceDecimal, err := decimal.NewFromString(req.QuoteOrderQty)
			if err != nil {
				return err
			}

			if priceDecimal.GreaterThanOrEqual(userFunds.Amount) {
				req.QuoteOrderQty = userFunds.Amount.String()
			}
		}
	} else if tradeType == gvar.TradeTypeMarket && tradeSide == gvar.TradeSideSell {
		err = l.GetUserFunds(req.Uid, marketSymbol.QuoteSymbol, &userFunds)
		//市价按量单处理卖
		if req.QuoteOrderQty != "" {
			priceDecimal, err := decimal.NewFromString(req.QuoteOrderQty)
			if err != nil {
				return err
			}

			if priceDecimal.GreaterThanOrEqual(userFunds.Amount) {
				req.QuoteOrderQty = userFunds.Amount.String()
			}
		} else {

			bookTicker, err := gvar.GetBookTicker(req.Symbol)
			if err != nil {
				return err
			}

			currentTime := time.Now()
			timeDifference := currentTime.Sub(bookTicker.UpdateTime)
			if timeDifference > 5*time.Minute {
				return fmt.Errorf("bookTicker update time timeout")
			}

			priceDecimal, err := decimal.NewFromString(req.Price)
			if err != nil {
				return err
			}

			bestBidPriceDecimal, err := decimal.NewFromString(bookTicker.BestBidPrice)
			if err != nil {
				return err
			}

			totalDecimal := priceDecimal.Div(bestBidPriceDecimal)
			b := l.isMarketPriceOrder(totalDecimal, userFunds.Amount)
			if b {
				req.Quantity = ""
				req.QuoteOrderQty = userFunds.Amount.String()
			}
		}
	}

	return nil
}

func (l *OrderLogic) limitTradeReqAdjustment(req *types.OrderRequest) error {
	marketSymbol, err := gvar.GetMarketSymbol(req.Symbol)
	if err != nil {
		return err
	}
	userFunds := db_model.UserFunds{}
	tradeSide := gvar.TradeSide(req.Side)
	if tradeSide == gvar.TradeSideBuy {
		err = l.GetUserFunds(req.Uid, marketSymbol.BaseSymbol, &userFunds)
		if err != nil {
			return err
		}
	} else {
		err = l.GetUserFunds(req.Uid, marketSymbol.QuoteSymbol, &userFunds)
		if err != nil {
			return err
		}
	}

	quantityDecimal, err := decimal.NewFromString(req.Quantity)
	if err != nil {
		return err
	}

	_, err = decimal.NewFromString(req.Price)
	if err != nil {
		return err
	}

	//如果
	if quantityDecimal.GreaterThan(userFunds.Amount) {
		req.Quantity = userFunds.Amount.String()
	}
	return nil
}

func (l *OrderLogic) GetUserFunds(uid string, symbol string, userFunds *db_model.UserFunds) error {
	sqlResult := gvar.PostgresClient.Select("amount").
		Where("uid = ? AND coin_name = ?", uid, symbol).
		Find(userFunds)

	if sqlResult.Error != nil {
		logc.Infof(context.Background(), "result.Error: %v", sqlResult.Error.Error())
		return sqlResult.Error
	}
	return nil
}

func (l *OrderLogic) isMarketPriceOrder(totalDecimal decimal.Decimal, userBalanceDecimal decimal.Decimal) bool {
	if l.svcCtx.Config.OrderMarketPriceConvertRate < 0.95 {
		l.svcCtx.Config.OrderMarketPriceConvertRate = 0.95
	}

	if l.svcCtx.Config.OrderMarketPriceConvertRate > 0.99 {
		l.svcCtx.Config.OrderMarketPriceConvertRate = 0.99
	}

	orderMarketRateDecimal := decimal.NewFromFloat(l.svcCtx.Config.OrderMarketPriceConvertRate)
	orderMarketDecimal := userBalanceDecimal.Mul(orderMarketRateDecimal)
	if totalDecimal.GreaterThan(orderMarketDecimal) {
		return true
	}
	return false
}

func (l *OrderLogic) CalculateCost(totalDecimal decimal.Decimal) (feeDecimal decimal.Decimal, freezeAmount decimal.Decimal) {
	if l.svcCtx.Config.FeeRate < 0 {
		l.svcCtx.Config.FeeRate = 0
	}
	feeRate := decimal.NewFromFloat(l.svcCtx.Config.FeeRate)
	feeDecimal = totalDecimal.Mul(feeRate)
	freezeAmount = feeDecimal.Add(totalDecimal)
	return
}

func (l *OrderLogic) UpdateOrderDB(req *types.OrderRequest, orderInfo *binance.CreateOrderResponse) error {
	return nil
}

func (l *OrderLogic) UpdateResp(resp *types.OrderResponse, orderInfo *binance.CreateOrderResponse) {
	resp.OrderId = orderInfo.ClientOrderID
	resp.Symbol = orderInfo.Symbol
	resp.Price = orderInfo.Price
	resp.OrigQty = orderInfo.OrigQuantity
	resp.ExecutedQty = orderInfo.ExecutedQuantity
	resp.Side = string(orderInfo.Side)
	resp.Type = string(orderInfo.Type)
	resp.Status = string(orderInfo.Status)
}
