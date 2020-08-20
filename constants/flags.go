/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package constants

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

//Note: Arranged alphabetically
var (
	AddMaintainer           = base.NewCLIFlag("addMaintainer", false, "AddMaintainer")
	AssetID                 = base.NewCLIFlag("assetID", "", "AssetID")
	ClassificationID        = base.NewCLIFlag("classificationID", "", "ClassificationID")
	ExpiresIn               = base.NewCLIFlag("expiresIn", int64(-1), "ExpiresIn")
	ExchangeRate            = base.NewCLIFlag("exchangeRate", "1", "ExchangeRateProperty")
	FromID                  = base.NewCLIFlag("fromID", "", "FromID")
	IdentityID              = base.NewCLIFlag("identityID", "", "IdentityID")
	MaintainersID           = base.NewCLIFlag("maintainersID", "", "MaintainersID")
	OwnableID               = base.NewCLIFlag("ownableID", "", "MakerOwnableID")
	Split                   = base.NewCLIFlag("split", "0", "Split")
	SplitID                 = base.NewCLIFlag("splitID", "", "SplitID")
	Coins                   = base.NewCLIFlag("coins", "", "Coins")
	ImmutableMetaProperties = base.NewCLIFlag("immutableMetaProperties", "", "immutableMetaProperties")
	ImmutableMetaTraits     = base.NewCLIFlag("immutableMetaTraits", "", "immutableMetaTraits")
	ImmutableProperties     = base.NewCLIFlag("immutableProperties", "", "immutableProperties")
	ImmutableTraits         = base.NewCLIFlag("immutableTraits", "", "immutableTraits")
	MetaFact                = base.NewCLIFlag("metaFact", "", "MetaFact")
	MaintainedTraits        = base.NewCLIFlag("maintainedTraits", "", "MaintainedTraits")
	MakerOwnableID          = base.NewCLIFlag("makerOwnableID", "", "MakerOwnableID")
	MakerOwnableSplit       = base.NewCLIFlag("makerOwnableSplit", "", "MakerOwnableSplit")
	MutableMetaProperties   = base.NewCLIFlag("mutableMetaProperties", "", "mutableMetaProperties")
	MutableMetaTraits       = base.NewCLIFlag("mutableMetaTraits", "", "mutableMetaTraits")
	MutableProperties       = base.NewCLIFlag("mutableProperties", "", "mutableProperties")
	MutableTraits           = base.NewCLIFlag("mutableTraits", "", "mutableTraits")
	To                      = base.NewCLIFlag("to", "", "To")
	ToID                    = base.NewCLIFlag("toID", "", "ToID")
	OrderID                 = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress            = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	TakerID                 = base.NewCLIFlag("takerID", "", "TakerID")
	TakerOwnableID          = base.NewCLIFlag("takerOwnableID", "", "TakerOwnableID")
	MaintainerID            = base.NewCLIFlag("makerSplit", int64(0), "MaintainerID")
	MakerSplit              = base.NewCLIFlag("makerSplit", int64(0), "MakerSplitProperty")
	MakerSplitID            = base.NewCLIFlag("makerSplitID", "", "MakerSplitIDProperty")
	MetaID                  = base.NewCLIFlag("metaID", "", "MetaID")
	MutateMaintainer        = base.NewCLIFlag("mutateMaintainer", false, "MutateMaintainer")
	RemoveMaintainer        = base.NewCLIFlag("removeMaintainer", false, "RemoveMaintainer")
	TakerOwnableSplit       = base.NewCLIFlag("takerOwnableSplit", int64(0), "TakerOwnableSplit")
	TakerSplitID            = base.NewCLIFlag("takerSplitID", "", "TakerSplitIDProperty")
)
