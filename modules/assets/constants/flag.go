package constants

import "github.com/persistenceOne/persistenceSDK/types"

var AssetID = types.NewCLIFlag("assetID", "", "AssetID")
var ChainID = types.NewCLIFlag("chainID", "", "GetChainID")
var MaintainersID = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
var ClassificationID = types.NewCLIFlag("classificationID", "", "GetClassificationID")
var Properties = types.NewCLIFlag("properties", "", "GetProperties")
var Lock = types.NewCLIFlag("lock", int64(-1), "Lock")
var Burn = types.NewCLIFlag("burn", int64(-1), "Burn")
