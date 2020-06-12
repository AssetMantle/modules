package constants

import "github.com/persistenceOne/persistenceSDK/types"

var AssetID = types.NewCLIFlag("assetID", "", "AssetID")
var ChainID = types.NewCLIFlag("chainID", "", "ChainID")
var MaintainersID = types.NewCLIFlag("maintainersID", "", "MaintainersID")
var ClassificationID = types.NewCLIFlag("classificationID", "", "ClassificationID")
var Properties = types.NewCLIFlag("properties", "", "Properties")
var Lock = types.NewCLIFlag("lock", -1, "Lock")
var Burn = types.NewCLIFlag("burn", -1, "Burn")
