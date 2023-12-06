package gvar

import (
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/config"
	"github.com/enterSpace9527/hifive-trade/trade/trade_service/internal/models/trade_model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"sync"
)

var BinanceClient *binance.Client
var PostgresClient *gorm.DB
var RedisClient *redis.Redis
var UserAddrMap sync.Map
var marketSymbolMap map[string]trade_model.MarketSymbol
var bookTickerMap map[string]trade_model.BookTicker

func initBinanceClient(c *config.Config) {
	if BinanceClient == nil {
		account1 := c.HifiveAccounts["Account1"]
		apiKey := account1["ApiKey"]
		secretKey := account1["SecretKey"]

		//todo_wr: 暂时使用测试网络
		binance.UseTestnet = true
		BinanceClient = binance.NewClient(apiKey, secretKey)
	}
}

func initPostgresClient(c *config.Config) error {
	if PostgresClient == nil {
		postgresC := &c.Postgres
		dns := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			postgresC.Host, postgresC.Port, postgresC.User, postgresC.Password, postgresC.Dbname)

		var loger logger.Interface = nil
		if c.StartMode == "test" {
			loger = logger.Default.LogMode(logger.Info)
		}

		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: loger,
		})
		if err != nil {
			return err
		}
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		PostgresClient = db
		return err
	}
	return nil
}

func initRedisClient(c *config.Config) {
	if RedisClient == nil {
		RedisClient = redis.MustNewRedis(c.Redis)
	}
}

func InitGVar(c *config.Config) error {
	initRedisClient(c)
	initBinanceClient(c)
	initBinanceUserDataSubscribe(c)
	err := initBinanceBookTickerSubscribe()
	if err != nil {
		return err
	}

	err = initPostgresClient(c)
	if err != nil {
		return err
	}

	return nil
}

func GetMarketSymbol(market string) (trade_model.MarketSymbol, error) {
	marketSymbol, ok := marketSymbolMap[market]
	if ok == true {
		return marketSymbol, nil
	}

	index := strings.Index(market, "USDT")
	if index == -1 || index == 0 {
		return marketSymbol, fmt.Errorf("wrong trading pair")
	}

	marketSymbol.QuoteSymbol = market[:index]
	marketSymbol.BaseSymbol = market[index:]
	marketSymbolMap[market] = marketSymbol
	return marketSymbol, nil
}

func GetBookTicker(market string) (trade_model.BookTicker, error) {
	bookTicker, ok := bookTickerMap[market]
	if ok == true {
		return bookTicker, nil
	}

	return bookTicker, fmt.Errorf("there is no such: %v", market)
}
