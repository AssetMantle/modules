package constants

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

//Note: Arranged alphabetically
var (
	AssetID          = base.NewCLIFlag("assetID", "", "OwnableID")
	Burn             = base.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = base.NewCLIFlag("classificationID", "", "ClassificationID")
	FromID           = base.NewCLIFlag("fromID", "", "FromID")
	IdentityID       = base.NewCLIFlag("identityID", "", "OwnerID")
	Lock             = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = base.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID        = base.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = base.NewCLIFlag("properties", "", "GetProperties")
	Split            = base.NewCLIFlag("split", "", "Split")
	SplitID          = base.NewCLIFlag("splitID", "", "SplitID")
	Traits           = base.NewCLIFlag("traits", "", "Traits")
	To               = base.NewCLIFlag("to", "", "To")
	ToID             = base.NewCLIFlag("toID", "", "ToID")
	OrderID          = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress     = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	MakerAssetAmount = base.NewCLIFlag("makerAssetAmount", int64(0), "MakerAssetAmount")
	MakerAssetData   = base.NewCLIFlag("makerAssetData", "persistence", "MakerAssetData")
	TakerAssetAmount = base.NewCLIFlag("takerAssetAmount", int64(0), "TakerAssetAmount")
	TakerAssetData   = base.NewCLIFlag("takerAssetData", "atom", "TakerAssetData")
	Salt             = base.NewCLIFlag("salt", int64(0), "Salt")
)
