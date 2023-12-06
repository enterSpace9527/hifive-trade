package main

import (
	"flag"
	"fmt"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/config"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/gvar"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/handler"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"os"
)

var configFile = flag.String("c", "etc/tradeapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	var cfg logx.LogConf
	_ = conf.FillDefault(&cfg)
	cfg.Mode = "file"
	logc.MustSetup(cfg)
	defer logc.Close()

	err := gvar.InitGVar(&c)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
