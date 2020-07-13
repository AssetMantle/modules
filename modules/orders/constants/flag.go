package constants

import "github.com/persistenceOne/persistenceSDK/types"

var OrderID = types.NewCLIFlag("orderID", "", "AssetID")

var BuyCoinDenom = types.NewCLIFlag("buyDemon", "commit", "buycoindemon")
var BuyCoinAmount = types.NewCLIFlag("buyAmount", int64(1), "buycoinadmoun")
var SellCoinDenom = types.NewCLIFlag("sellDemon", "atom", "sellcoindenom")
var SellCoinAmount = types.NewCLIFlag("sellAmount", int64(2), "buycoindenom")
