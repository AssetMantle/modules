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
	Burn                    = base.NewCLIFlag("burn", int64(-1), "BurnProperty")
	ClassificationID        = base.NewCLIFlag("classificationID", "", "ClassificationID")
	ExchangeRate            = base.NewCLIFlag("exchangeRate", "1", "ExchangeRateProperty")
	FromID                  = base.NewCLIFlag("fromID", "", "FromID")
	IdentityID              = base.NewCLIFlag("identityID", "", "IdentityID")
	Lock                    = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID           = base.NewCLIFlag("maintainersID", "", "MaintainersID")
	OwnableID               = base.NewCLIFlag("ownableID", "", "OwnableID")
	Split                   = base.NewCLIFlag("split", "0", "Split")
	SplitID                 = base.NewCLIFlag("splitID", "", "SplitID")
	Coins                   = base.NewCLIFlag("coins", "", "Coins")
	ImmutableMetaProperties = base.NewCLIFlag("immutableMetaProperties", "", "immutableMetaProperties")
	ImmutableMetaTraits     = base.NewCLIFlag("immutableMetaTraits", "", "immutableMetaTraits")
	ImmutableProperties     = base.NewCLIFlag("immutableProperties", "", "immutableProperties")
	ImmutableTraits         = base.NewCLIFlag("immutableTraits", "", "immutableTraits")
	MetaFact                = base.NewCLIFlag("metaFact", "", "MetaFact")
	MaintainedTraits        = base.NewCLIFlag("maintainedTraits", "", "MaintainedTraits")
	MutableMetaProperties   = base.NewCLIFlag("mutableMetaProperties", "", "mutableMetaProperties")
	MutableMetaTraits       = base.NewCLIFlag("mutableMetaTraits", "", "mutableMetaTraits")
	MutableProperties       = base.NewCLIFlag("mutableProperties", "", "mutableProperties")
	MutableTraits           = base.NewCLIFlag("mutableTraits", "", "mutableTraits")
	To                      = base.NewCLIFlag("to", "", "To")
	ToID                    = base.NewCLIFlag("toID", "", "ToID")
	OrderID                 = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress            = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	MaintainerID            = base.NewCLIFlag("makerSplit", int64(0), "MaintainerID")
	MakerSplit              = base.NewCLIFlag("makerSplit", int64(0), "MakerSplitProperty")
	MakerSplitID            = base.NewCLIFlag("makerSplitID", "", "MakerSplitIDProperty")
	MetaID                  = base.NewCLIFlag("metaID", "", "MetaID")
	MutateMaintainer        = base.NewCLIFlag("mutateMaintainer", false, "MutateMaintainer")
	RemoveMaintainer        = base.NewCLIFlag("removeMaintainer", false, "RemoveMaintainer")
	TakerSplit              = base.NewCLIFlag("takerSplit", int64(0), "TakerSplit")
	TakerSplitID            = base.NewCLIFlag("takerSplitID", "", "TakerSplitIDProperty")
)
