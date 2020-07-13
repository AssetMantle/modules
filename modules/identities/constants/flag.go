package constants

import "github.com/persistenceOne/persistenceSDK/types"

var IdentityID = types.NewCLIFlag("identityID", "", "IdentityID")
var ChainID = types.NewCLIFlag("chainID", "", "GetChainID")
var MaintainersID = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
var ClassificationID = types.NewCLIFlag("classificationID", "", "GetClassificationID")
var Properties = types.NewCLIFlag("properties", "", "GetProperties")
var Lock = types.NewCLIFlag("lock", int64(-1), "Lock")
var Burn = types.NewCLIFlag("issue", int64(-1), "Burn")
