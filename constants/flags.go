package constants

import "github.com/persistenceOne/persistenceSDK/types"

var (
	AssetID          = types.NewCLIFlag("assetID", "", "OrderID")
	Burn             = types.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = types.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID       = types.NewCLIFlag("identityID", "", "IdentityID")
	Lock             = types.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	Properties       = types.NewCLIFlag("properties", "", "GetProperties")
	To               = types.NewCLIFlag("to", "", "To")
	OrderID          = types.NewCLIFlag("orderID", "", "OrderID")
)
