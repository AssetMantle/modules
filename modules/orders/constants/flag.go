package constants

import "github.com/persistenceOne/persistenceSDK/types"

var AssetID = types.NewCLIFlag("assetID", "", "AssetID")
var ChainID = types.NewCLIFlag("chainID", "", "GetChainID")
var MaintainersID = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
var ClassificationID = types.NewCLIFlag("classificationID", "", "GetClassificationID")
var Properties = types.NewCLIFlag("properties", "", "GetProperties")
var Lock = types.NewCLIFlag("lock", int64(-1), "Lock")
var Burn = types.NewCLIFlag("burn", int64(-1), "Burn")

var BuyCoinDenom = types.NewCLIFlag("buyD", "commit", "")
var BuyCoinAmount = types.NewCLIFlag("buyA", int64(1), "")
var SellCoinDenom = types.NewCLIFlag("sellD", "atom", "")
var SellCoinAmount = types.NewCLIFlag("sellA", int64(2), "")
