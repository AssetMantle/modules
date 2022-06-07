// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

// Note: Arranged alphabetically
var (
	AddMaintainer           = baseHelpers.NewCLIFlag("addMaintainer", false, "AddMaintainer")
	AssetID                 = baseHelpers.NewCLIFlag("assetID", "", "AssetID")
	ClassificationID        = baseHelpers.NewCLIFlag("classificationID", "", "ClassificationID")
	Coins                   = baseHelpers.NewCLIFlag("coins", "", "Coins")
	ExpiresIn               = baseHelpers.NewCLIFlag("expiresIn", int64(-1), "ExpiresIn")
	FromID                  = baseHelpers.NewCLIFlag("fromID", "", "FromID")
	IdentityID              = baseHelpers.NewCLIFlag("identityID", "", "IdentityID")
	ImmutableMetaProperties = baseHelpers.NewCLIFlag("immutableMetaProperties", "", "immutableMetaProperties")
	ImmutableProperties     = baseHelpers.NewCLIFlag("immutableProperties", "", "immutableProperties")
	KafkaNodes              = baseHelpers.NewCLIFlag("kafkaNodes", "localhost:9092", "Space separated addresses in quotes of the kafka listening node: example: --kafkaPort \"addr1 addr2\" ")
	MetaFact                = baseHelpers.NewCLIFlag("metaFact", "", "MetaFact")
	MaintainerID            = baseHelpers.NewCLIFlag("maintainerID", "", "MaintainerID")
	MaintainedProperties    = baseHelpers.NewCLIFlag("maintainedProperties", "", "MaintainedProperties")
	MakerOwnableID          = baseHelpers.NewCLIFlag("makerOwnableID", "", "MakerOwnableID")
	MakerOwnableSplit       = baseHelpers.NewCLIFlag("makerOwnableSplit", "", "MakerOwnableSplit")
	MutableMetaProperties   = baseHelpers.NewCLIFlag("mutableMetaProperties", "", "mutableMetaProperties")
	MutableProperties       = baseHelpers.NewCLIFlag("mutableProperties", "", "mutableProperties")
	MetaID                  = baseHelpers.NewCLIFlag("metaID", "", "MetaID")
	MutateMaintainer        = baseHelpers.NewCLIFlag("mutateMaintainer", false, "MutateMaintainer")
	NubID                   = baseHelpers.NewCLIFlag("nubID", "", "NubID")
	OrderID                 = baseHelpers.NewCLIFlag("orderID", "", "OrderID")
	OwnableID               = baseHelpers.NewCLIFlag("ownableID", "", "MakerOwnableID")
	Queuing                 = baseHelpers.NewCLIFlag("queuing", false, "Enable kafka queuing and squashing of transactions")
	RemoveMaintainer        = baseHelpers.NewCLIFlag("removeMaintainer", false, "RemoveMaintainer")
	Value                   = baseHelpers.NewCLIFlag("value", "0", "Value")
	SplitID                 = baseHelpers.NewCLIFlag("splitID", "", "SplitID")
	To                      = baseHelpers.NewCLIFlag("to", "", "To")
	ToID                    = baseHelpers.NewCLIFlag("toID", "", "ToID")
	TakerOwnableID          = baseHelpers.NewCLIFlag("takerOwnableID", "", "TakerOwnableID")
	TakerOwnableSplit       = baseHelpers.NewCLIFlag("takerOwnableSplit", "0", "TakerOwnableSplit")
)
