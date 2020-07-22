package constants

import (
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var (
	AssetID             = utility.NewCLIFlag("assetID", "", "OwnableID")
	Burn                = utility.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID    = utility.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID          = utility.NewCLIFlag("identityID", "", "OwnerID")
	Lock                = utility.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID       = utility.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID           = utility.NewCLIFlag("ownableID", "", "OwnableID")
	Properties          = utility.NewCLIFlag("properties", "", "GetProperties")
	Split               = utility.NewCLIFlag("split", "", "Split")
	SplitID             = utility.NewCLIFlag("splitID", "", "SplitID")
	To                  = utility.NewCLIFlag("to", "", "To")
	ToID                = utility.NewCLIFlag("toID", "", "ToID")
	OrderID             = utility.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress        = utility.NewCLIFlag("takerAddress", "", "TakerAddress")
	SenderAddress       = utility.NewCLIFlag("senderAddress", "", "SenderAddress")
	FeeRecipientAddress = utility.NewCLIFlag("feeRecipientAddress", "", "FeeRecipientAddress")
	MakerAssetAmount    = utility.NewCLIFlag("makerAssetAmount", int64(0), "MakerAssetAmount")
	MakerAssetData      = utility.NewCLIFlag("makerAssetData", "persistence", "MakerAssetData")
	MakerFee            = utility.NewCLIFlag("makerFee", int64(0), "MakerFee")
	MakerFeeAssetData   = utility.NewCLIFlag("makerFeeAssetData", "", "MakerFeeAssetData")
	TakerAssetAmount    = utility.NewCLIFlag("takerAssetAmount", int64(0), "TakerAssetAmount")
	TakerAssetData      = utility.NewCLIFlag("takerAssetData", "atom", "TakerAssetData")
	TakerFee            = utility.NewCLIFlag("takerFee", int64(0), "TakerFee")
	TakerFeeAssetData   = utility.NewCLIFlag("takerFeeAssetData", "", "TakerFeeAssetData")
	ExpirationTime      = utility.NewCLIFlag("expirationTime", int64(0), "ExpirationTime")
	Salt                = utility.NewCLIFlag("salt", int64(0), "Salt")
)
