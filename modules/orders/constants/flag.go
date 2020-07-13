package constants

import "github.com/persistenceOne/persistenceSDK/types"

var AssetID = types.NewCLIFlag("assetID", "", "AssetID")
var ChainID = types.NewCLIFlag("chainID", "", "GetChainID")
var MaintainersID = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
var ClassificationID = types.NewCLIFlag("classificationID", "", "GetClassificationID")
var Properties = types.NewCLIFlag("properties", "", "GetProperties")
var Lock = types.NewCLIFlag("lock", int64(-1), "Lock")
var Burn = types.NewCLIFlag("burn", int64(-1), "Burn")

var BuyCoinDenom = types.NewCLIFlag("buyDemon", "commit", "buycoindemon")
var BuyCoinAmount = types.NewCLIFlag("buyAmount", int64(1), "buycoinadmoun")
var SellCoinDenom = types.NewCLIFlag("sellDemon", "atom", "sellcoindenom")
var SellCoinAmount = types.NewCLIFlag("sellAmount", int64(2), "buycoindenom")
