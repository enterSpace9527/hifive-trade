package logic

import (
	"context"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenOrdersLogic {
	return &OpenOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OpenOrdersLogic) OpenOrders(req *types.OrderRequest) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line

	return
}
