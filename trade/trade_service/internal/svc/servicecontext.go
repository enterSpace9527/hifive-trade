package svc

import (
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
