package constants

import (
	"github.com/persistenceOne/persistenceSDK/types"
)

var (
	AssetID             = types.NewCLIFlag("assetID", "", "OrderID")
	Burn                = types.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID    = types.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID          = types.NewCLIFlag("identityID", "", "IdentityID")
	Lock                = types.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID       = types.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	Properties          = types.NewCLIFlag("properties", "", "GetProperties")
	To                  = types.NewCLIFlag("to", "", "To")
	OrderID             = types.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress        = types.NewCLIFlag("takerAddress", "", "TakerAddress")
	SenderAddress       = types.NewCLIFlag("senderAddress", "", "SenderAddress")
	FeeRecipientAddress = types.NewCLIFlag("feeRecipientAddress", "", "FeeRecipientAddress")
	MakerAssetAmount    = types.NewCLIFlag("makerAssetAmount", int64(0), "MakerAssetAmount")
	MakerAssetData      = types.NewCLIFlag("makerAssetData", "persistence", "MakerAssetData")
	MakerFee            = types.NewCLIFlag("makerFee", int64(0), "MakerFee")
	MakerFeeAssetData   = types.NewCLIFlag("makerFeeAssetData", "", "MakerFeeAssetData")
	TakerAssetAmount    = types.NewCLIFlag("takerAssetAmount", int64(0), "TakerAssetAmount")
	TakerAssetData      = types.NewCLIFlag("takerAssetData", "atom", "TakerAssetData")
	TakerFee            = types.NewCLIFlag("takerFee", int64(0), "TakerFee")
	TakerFeeAssetData   = types.NewCLIFlag("takerFeeAssetData", "", "TakerFeeAssetData")
	ExpirationTime      = types.NewCLIFlag("expirationTime", int64(0), "ExpirationTime")
	Salt                = types.NewCLIFlag("salt", int64(0), "Salt")
)
