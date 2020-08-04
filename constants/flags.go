package constants

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var (
	AssetID          = base.NewCLIFlag("assetID", "", "OwnableID")
	Burn             = base.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = base.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID       = base.NewCLIFlag("identityID", "", "OwnerID")
	Lock             = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = base.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID        = base.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = base.NewCLIFlag("properties", "", "GetProperties")
	Split            = base.NewCLIFlag("split", "", "Split")
	SplitID          = base.NewCLIFlag("splitID", "", "SplitID")
	To               = base.NewCLIFlag("to", "", "To")
	ToID             = base.NewCLIFlag("toID", "", "ToID")
	OrderID          = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress     = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	MakerAssetAmount = base.NewCLIFlag("makerAssetAmount", int64(0), "MakerAssetAmount")
	MakerAssetData   = base.NewCLIFlag("makerAssetData", "", "MakerAssetData")
	MakerAssetType   = base.NewCLIFlag("makerAssetType", "coin", "MakerAssetType")
	TakerAssetAmount = base.NewCLIFlag("takerAssetAmount", int64(0), "TakerAssetAmount")
	TakerAssetData   = base.NewCLIFlag("takerAssetData", "", "TakerAssetData")
	TakerAssetType   = base.NewCLIFlag("takerAssetType", "coin", "TakerAssetType")
)
