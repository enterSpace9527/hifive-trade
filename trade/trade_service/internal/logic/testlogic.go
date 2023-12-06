package logic

import (
	"context"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/gvar"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/db_model"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test() (resp *types.CommResp, err error) {

	order := db_model.Order{}
	gvar.PostgresClient.First(&order, 1)
	logc.Infof(context.Background(), "order.Price=%v", order.Price.String())
	return
}
