// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
)

// Note: Arranged alphabetically
// TODO define usages
var (
	AssetID                 = baseHelpers.NewCLIFlag("assetID", "", "AssetID")
	CanAddMaintainer        = baseHelpers.NewCLIFlag("canAddMaintainer", false, "CanAddMaintainer")
	CanBurnAsset            = baseHelpers.NewCLIFlag("canBurnAsset", false, "CanBurnAsset")
	CanCancelOrder          = baseHelpers.NewCLIFlag("canCancelOrder", false, "CanCancelOrder")
	CanIssueIdentity        = baseHelpers.NewCLIFlag("canIssueIdentity", false, "CanIssueIdentity")
	CanMakeOrder            = baseHelpers.NewCLIFlag("canMakeOrder", false, "CanMakeOrder")
	CanMintAsset            = baseHelpers.NewCLIFlag("canMintAsset", false, "CanMintAsset")
	CanModifyOrder          = baseHelpers.NewCLIFlag("canModifyOrder", false, "CanModifyOrder")
	CanMutateIdentity       = baseHelpers.NewCLIFlag("canMutateIdentity", false, "CanMutateIdentity")
	CanMutateMaintainer     = baseHelpers.NewCLIFlag("canMutateMaintainer", false, "CanMutateMaintainer")
	CanRemoveMaintainer     = baseHelpers.NewCLIFlag("canRemoveMaintainer", false, "CanRemoveMaintainer")
	CanRenumerateAsset      = baseHelpers.NewCLIFlag("canRenumerateAsset", false, "CanRenumerateAsset")
	CanQuashIdentity        = baseHelpers.NewCLIFlag("canQuashIdentity", false, "CanQuashIdentity")
	ClassificationID        = baseHelpers.NewCLIFlag("classificationID", "", "ClassificationID")
	Coins                   = baseHelpers.NewCLIFlag("coins", "", "Coins")
	Data                    = baseHelpers.NewCLIFlag("data", "", "Data")
	DataID                  = baseHelpers.NewCLIFlag("dataID", "", "DataID")
	ExpiresIn               = baseHelpers.NewCLIFlag("expiresIn", int64(-1), "ExpiresIn")
	FromIdentityID          = baseHelpers.NewCLIFlag("fromIdentityID", "", "FromIdentityID")
	IdentityID              = baseHelpers.NewCLIFlag("identityID", "", "IdentityID")
	ImmutableMetaProperties = baseHelpers.NewCLIFlag("immutableMetaProperties", "", "immutableMetaProperties")
	ImmutableProperties     = baseHelpers.NewCLIFlag("immutableProperties", "", "immutableProperties")
	KafkaNodes              = baseHelpers.NewCLIFlag("kafkaNodes", "localhost:9092", "Space separated addresses in quotes of the kafka listening node: example: --kafkaPort \"addr1 addr2\" ")
	MaintainerID            = baseHelpers.NewCLIFlag("maintainerID", "", "MaintainerID")
	MaintainedProperties    = baseHelpers.NewCLIFlag("maintainedProperties", "", "MaintainedProperties")
	MakerOwnableID          = baseHelpers.NewCLIFlag("makerOwnableID", "", "MakerOwnableID")
	MakerOwnableSplit       = baseHelpers.NewCLIFlag("makerOwnableSplit", "", "MakerOwnableSplit")
	MutableMetaProperties   = baseHelpers.NewCLIFlag("mutableMetaProperties", "", "mutableMetaProperties")
	MutableProperties       = baseHelpers.NewCLIFlag("mutableProperties", "", "mutableProperties")
	NubID                   = baseHelpers.NewCLIFlag("nubID", "", "NubID")
	OrderID                 = baseHelpers.NewCLIFlag("orderID", "", "OrderID")
	OwnableID               = baseHelpers.NewCLIFlag("ownableID", "", "MakerOwnableID")
	Queuing                 = baseHelpers.NewCLIFlag("queuing", false, "Enable kafka queuing and squashing of transactions")
	SplitID                 = baseHelpers.NewCLIFlag("splitID", "", "SplitID")
	To                      = baseHelpers.NewCLIFlag("to", "", "To")
	ToIdentityID            = baseHelpers.NewCLIFlag("toIdentityID", "", "ToIdentityID")
	TakerID                 = baseHelpers.NewCLIFlag("takerID", "", "TakerID")
	TakerOwnableID          = baseHelpers.NewCLIFlag("takerOwnableID", "", "TakerOwnableID")
	TakerOwnableSplit       = baseHelpers.NewCLIFlag("takerOwnableSplit", "0", "TakerOwnableSplit")
	Value                   = baseHelpers.NewCLIFlag("value", "0", "Value")
	Offset                  = baseHelpers.NewCLIFlag("offset", 0, "offset")
	Limit                   = baseHelpers.NewCLIFlag("limit", 100, "limit")
)
