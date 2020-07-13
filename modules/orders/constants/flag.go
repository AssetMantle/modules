package constants

import "github.com/persistenceOne/persistenceSDK/types"

var BuyCoinDenom = types.NewCLIFlag("buyDenom", "commit", "buycoindemon")
var BuyCoinAmount = types.NewCLIFlag("buyAmount", int64(1), "buycoinadmoun")
var SellCoinDenom = types.NewCLIFlag("sellDenom", "atom", "sellcoindenom")
var SellCoinAmount = types.NewCLIFlag("sellAmount", int64(2), "sellcoinamount")
var Lock = types.NewCLIFlag("lock", int64(-1), "Lock")
var Burn = types.NewCLIFlag("burn", int64(-1), "Burn")
var OrderID = types.NewCLIFlag("orderID", "", "OrderID")
var Properties = types.NewCLIFlag("properties", "", "--properties=test1:test2")
