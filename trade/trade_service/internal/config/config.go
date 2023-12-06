package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type Redis struct {
	Host string
	Type string `json:",default=node,options=node|cluster"`
	Pass string `json:",optional"`
	Tls  bool   `json:",optional"`
}

type Config struct {
	rest.RestConf
	StartMode                   string
	FeeRate                     float64
	OrderMarketPriceConvertRate float64
	Postgres                    Postgres
	Redis                       redis.RedisConf
	HifiveAccounts              map[string]map[string]string
	WalletService               map[string]string
}
