package logic

import (
	"context"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllOrdersLogic {
	return &AllOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllOrdersLogic) AllOrders(req *types.OrderRequest) (resp *types.OrderResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
