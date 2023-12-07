package bian_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/google/uuid"
	"github.com/sony/sonyflake"
	"testing"
)

const BianUrl = "https://testnet.binance.vision"

// const apiKey = "U9MkUnHNE2v7fsJkAnC0hmKlNdpYjjok7GixpLKNDCvL7T2Qb0YQGbq0inCisGCd"
// const secretKey = "7WnlRaQ9nnXFwE3qDU5Pz4qGBszLOLalLAVPlj1FRILNkRx34uBrG45a0ZBcjmYx"

const apiKey = "iqNPSGNtFQG6Y4ghUsz8iy1Wuk2fDLyRoLaWZqwtxHYmqr61bGXxP96KUkPL4D5J"
const secretKey = "qaAHsHFpKvxn05NS3xWTOwJJy3UnEyaLIeTU7RWuu0NxPgYnaEue8hB2Ua4hySc5"

//const apiKey = "Sc2PGTsW6VlCcnlKj8RmFNxVk5UCK0PeL6O1kAenYbFlSEKLyTGElyrIEl9O3i2O"
//const secretKey = "lId7tLZ6y5ANSwnZuFHKbkzNhOJqvI6ka9WPf9yx9QkbYK82fBHDb1VcSfWY4r2c"

func TestExchangeInfo(t *testing.T) {

	client := binance.NewClient(apiKey, secretKey)
	ExchangeInfoSvc := client.NewExchangeInfoService()

	ctx := context.Background()
	ExchangeInfo, err := ExchangeInfoSvc.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hello111", ExchangeInfo.Symbols)
	fmt.Println("hello222")
	//fmt.Println("hello222", ExchangeInfo.RateLimits)
	//fmt.Println("hello333", ExchangeInfo.ExchangeFilters)
}

func TestUserData(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	userAssetSvc := client.NewGetUserAsset()

	ctx := context.Background()
	ExchangeInfo, err := userAssetSvc.Do(ctx)
	fmt.Println(ExchangeInfo, err)
}

func TestUserData1(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	accountSnapshot := client.NewGetAccountSnapshotService()
	ctx := context.Background()
	snapshot, err := accountSnapshot.Do(ctx)
	fmt.Println(snapshot, err)
}

func TestUserData2(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	accountSnapshot := client.NewGetAccountSnapshotService()
	ctx := context.Background()
	snapshot, err := accountSnapshot.Do(ctx)
	fmt.Println(snapshot, err)
}

func TestUserData3(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	accountSnapshot := client.NewGetAccountSnapshotService()
	ctx := context.Background()
	snapshot, err := accountSnapshot.Do(ctx)
	fmt.Println(snapshot, err)
}

func TestUserData4(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	assetDetail := client.NewGetAssetDetailService()
	ctx := context.Background()
	detail, err := assetDetail.Do(ctx)
	fmt.Println(detail, err)
}

// 当前BNB 19.34000000
func TestUserData5(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	account := client.NewGetAccountService()
	ctx := context.Background()
	detail, err := account.Do(ctx)
	fmt.Println(detail, err)
}

func TestUserMarketOrder(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	createOrder := client.NewCreateOrderService()
	createOrder.Symbol("LTCUSDT")
	createOrder.Type("MARKET")
	createOrder.Side("BUY")
	createOrder.Quantity("10")
	//createOrder.TimeInForce("GTC")
	//createOrder.QuoteOrderQty("3000")
	//createOrder.NewOrderRespType("RESULT")
	ctx := context.Background()
	orderResp, err := createOrder.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("OrderID:%v ", orderResp.OrderID)
	fmt.Printf("Symbol:%v ", orderResp.Symbol)
	fmt.Printf("ClientOrderID:%v ", orderResp.ClientOrderID)
	fmt.Printf("Price:%v ", orderResp.Price)
	fmt.Printf("ExecutedQuantity:%v ", orderResp.ExecutedQuantity)
	fmt.Printf("OrigQuantity:%v ", orderResp.OrigQuantity)
	fmt.Printf("Type:%v ", orderResp.Type)
	fmt.Printf("Side:%v ", orderResp.Side)
	fmt.Printf("Fills:%v ", orderResp.Fills)
	fmt.Printf("Status:%v ", orderResp.Status)
	fmt.Printf("TimeInForce:%v ", orderResp.TimeInForce)

	fmt.Println("")
	for _, fill := range orderResp.Fills {
		fmt.Println("-----fill-----")
		fmt.Printf("fill.Price:%v ", fill.Price)
		fmt.Printf("fill.Quantity:%v ", fill.Quantity)
		fmt.Printf("fill.Commission:%v ", fill.Commission)
		fmt.Printf("fill.CommissionAsset:%v ", fill.CommissionAsset)
		fmt.Printf("fill.TradeID:%v\n ", fill.TradeID)
	}
	fmt.Println("")

	//orderId1Eth := "8SM81zVUHiXq8Nq7hGgFvR"
	//&{ETHUSDT 7742208 JjECs3zLXZHZGj3rEloyxk 1701396763884 0.00000000 0.47930000 0.47930000 999.90511200 false FILLED GTC MARKET BUY [0xc000100af0 0xc000100b40 0xc000100b90]  } <nil>
	//&{BTCUSDT 9113527 5SovjE1N3GXMTjQ72qcZVD 1701240030592 0.00000000 0.10000000 0.10000000 3815.07231940 false FILLED GTC MARKET BUY [0xc000100af0 0xc000100b40 0xc000100b90 0xc000100be0 0xc000100c30 0xc000100c80 0xc000100cd0 0xc000100d20 0xc000100d70 0xc000100dc0 0xc000100e10 0xc000100e60 0xc000100eb0]  } <nil>
}

func TestUserLimitOrder(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	createOrder := client.NewCreateOrderService()
	createOrder.Symbol("ETHUSDT")
	createOrder.Type("LIMIT")
	createOrder.Side("BUY")
	createOrder.Quantity("0.5")
	createOrder.Price("2248.8")
	createOrder.TimeInForce("GTC")
	ctx := context.Background()
	orderResp, err := createOrder.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("OrderID:%v ", orderResp.OrderID)
	fmt.Printf("Symbol:%v ", orderResp.Symbol)
	fmt.Printf("ClientOrderID:%v ", orderResp.ClientOrderID)
	fmt.Printf("Price:%v ", orderResp.Price)
	fmt.Printf("ExecutedQuantity:%v ", orderResp.ExecutedQuantity)
	fmt.Printf("OrigQuantity:%v ", orderResp.OrigQuantity)
	fmt.Printf("Type:%v ", orderResp.Type)
	fmt.Printf("Side:%v ", orderResp.Side)
	fmt.Printf("Fills:%v ", orderResp.Fills)
	fmt.Printf("Status:%v ", orderResp.Status)
	fmt.Printf("TimeInForce:%v ", orderResp.TimeInForce)

	fmt.Println("")
	for _, fill := range orderResp.Fills {
		fmt.Printf("fill.Price:%v ", fill.Price)
		fmt.Printf("fill.Quantity:%v ", fill.Quantity)
		fmt.Printf("fill.Commission:%v ", fill.Commission)
		fmt.Printf("fill.CommissionAsset:%v ", fill.CommissionAsset)
		fmt.Printf("fill.TradeID:%v ", fill.TradeID)
		fmt.Println("-----fill-----")
	}
	fmt.Println("")

	//orderId1Eth := "8SM81zVUHiXq8Nq7hGgFvR"
	//&{ETHUSDT 7742208 JjECs3zLXZHZGj3rEloyxk 1701396763884 0.00000000 0.47930000 0.47930000 999.90511200 false FILLED GTC MARKET BUY [0xc000100af0 0xc000100b40 0xc000100b90]  } <nil>
	//&{BTCUSDT 9113527 5SovjE1N3GXMTjQ72qcZVD 1701240030592 0.00000000 0.10000000 0.10000000 3815.07231940 false FILLED GTC MARKET BUY [0xc000100af0 0xc000100b40 0xc000100b90 0xc000100be0 0xc000100c30 0xc000100c80 0xc000100cd0 0xc000100d20 0xc000100d70 0xc000100dc0 0xc000100e10 0xc000100e60 0xc000100eb0]  } <nil>
}

func TestListOrder(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	listOrderService := client.NewListOrdersService()
	listOrderService.Symbol("ETHUSDT")

	ctx := context.Background()
	orders, err := listOrderService.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range orders {
		fmt.Printf("Symbol:%v ", d.Symbol)
		fmt.Printf("OrderID:%v ", d.OrderID)
		fmt.Printf("OrderListId:%v ", d.OrderListId)
		fmt.Printf("ClientOrderID:%v ", d.ClientOrderID)
		fmt.Printf("Price:%v ", d.Price)
		fmt.Printf("OrigQuantity:%v ", d.OrigQuantity)
		fmt.Printf("ExecutedQuantity:%v ", d.ExecutedQuantity)
		fmt.Printf("CummulativeQuoteQuantity:%v ", d.CummulativeQuoteQuantity)
		fmt.Printf("Status:%v ", d.Status)
		fmt.Printf("TimeInForce:%v ", d.TimeInForce)
		fmt.Printf("Type:%v ", d.Type)
		fmt.Printf("Side:%v ", d.Side)
		fmt.Printf("StopPrice:%v ", d.StopPrice)
		fmt.Printf("IcebergQuantity:%v ", d.IcebergQuantity)
		fmt.Printf("Time:%v ", d.Time)
		fmt.Printf("UpdateTime:%v ", d.UpdateTime)
		fmt.Printf("IsWorking:%v ", d.IsWorking)
		fmt.Printf("IsIsolated:%v ", d.IsIsolated)
		fmt.Printf("OrigQuoteOrderQuantity:%v ", d.OrigQuoteOrderQuantity)

		fmt.Println("---------")
	}

	//&{BTCUSDT 9113527 5SovjE1N3GXMTjQ72qcZVD 1701240030592 0.00000000 0.10000000 0.10000000 3815.07231940 false FILLED GTC MARKET BUY [0xc000100af0 0xc000100b40 0xc000100b90 0xc000100be0 0xc000100c30 0xc000100c80 0xc000100cd0 0xc000100d20 0xc000100d70 0xc000100dc0 0xc000100e10 0xc000100e60 0xc000100eb0]  } <nil>
}

func TestGetOrder(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	getOrder := client.NewGetOrderService()
	getOrder.Symbol("ETHUSDT")
	getOrder.OrigClientOrderID("18kdgE38xblHd0qncemScf")
	ctx := context.Background()
	detail, err := getOrder.Do(ctx)
	fmt.Println(detail, err)
}

func TestOpenOrder(t *testing.T) {
	client := binance.NewClient(apiKey, secretKey)
	getOrder := client.NewListOpenOrdersService()
	getOrder.Symbol("ETHUSDT")
	ctx := context.Background()
	detail, err := getOrder.Do(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, d := range detail {
		fmt.Printf("Symbol:%v ", d.Symbol)
		fmt.Printf("OrderID:%v ", d.OrderID)
		fmt.Printf("OrderListId:%v ", d.OrderListId)
		fmt.Printf("ClientOrderID:%v ", d.ClientOrderID)
		fmt.Printf("PriceClientOrderID:%v ", d.Price)
		fmt.Printf("OrigQuantity:%v ", d.OrigQuantity)
		fmt.Printf("ExecutedQuantity:%v ", d.ExecutedQuantity)
		fmt.Printf("CummulativeQuoteQuantity:%v ", d.CummulativeQuoteQuantity)
		fmt.Printf("Status:%v ", d.Status)
		fmt.Printf("TimeInForce:%v ", d.TimeInForce)
		fmt.Printf("Type:%v ", d.Type)
		fmt.Printf("Side:%v ", d.Side)
		fmt.Printf("StopPrice:%v ", d.StopPrice)
		fmt.Printf("IcebergQuantity:%v ", d.IcebergQuantity)
		fmt.Printf("Time:%v ", d.Time)
		fmt.Printf("UpdateTime:%v ", d.UpdateTime)
		fmt.Printf("IsWorking:%v ", d.IsWorking)
		fmt.Printf("IsIsolated:%v ", d.IsIsolated)
		fmt.Printf("OrigQuoteOrderQuantity:%v ", d.OrigQuoteOrderQuantity)

		fmt.Println("")
	}
}

//func TestUserData7(t *testing.T) {
//	client := binance.NewClient(apiKey, secretKey)
//	createOrder := client.
//		createOrder.Symbol("BTCUSDT")
//	createOrder.Type("MARKET")
//	createOrder.Side("BUY")
//	createOrder.Quantity("0.1")
//	ctx := context.Background()
//	err := createOrder.Test(ctx)
//	fmt.Println(err)
//}

func TestUsertt(t *testing.T) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("id:", id)

	uuidObj, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return
	}
	uuidBytes := uuidObj[:]

	// 计算 SHA-256 散列值
	hash := sha256.Sum256(uuidBytes)

	// 将散列值转为字符串
	hashString := hex.EncodeToString(hash[:])
	fmt.Println("uuidObj:", uuidObj.ID())
	fmt.Println("hashString:", hashString)
}

func init() {
	binance.UseTestnet = true
}
