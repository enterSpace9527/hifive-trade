package logic

import (
	"context"

	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelOrderLogic) CancelOrder(req *types.OrderRequest) (resp *types.CommResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.CommResp{
		Code:    200,
		Message: "success",
	}
	return
}
