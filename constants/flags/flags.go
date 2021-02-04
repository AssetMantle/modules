/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package flags

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

//Note: Arranged alphabetically
var (
	AddMaintainer           = base.NewCLIFlag("addMaintainer", false, "AddMaintainer")
	AssetID                 = base.NewCLIFlag("assetID", "", "AssetID")
	ClassificationID        = base.NewCLIFlag("classificationID", "", "ClassificationID")
	Coins                   = base.NewCLIFlag("coins", "", "Coins")
	ExpiresIn               = base.NewCLIFlag("expiresIn", int64(-1), "ExpiresIn")
	ExchangeRate            = base.NewCLIFlag("exchangeRate", "1", "ExchangeRateProperty")
	FromID                  = base.NewCLIFlag("fromID", "", "FromID")
	IdentityID              = base.NewCLIFlag("identityID", "", "IdentityID")
	ImmutableMetaProperties = base.NewCLIFlag("immutableMetaProperties", "", "immutableMetaProperties")
	ImmutableProperties     = base.NewCLIFlag("immutableProperties", "", "immutableProperties")
	MetaFact                = base.NewCLIFlag("metaFact", "", "MetaFact")
	MaintainerID            = base.NewCLIFlag("maintainerID", "", "MaintainerID")
	MaintainedProperties    = base.NewCLIFlag("maintainedProperties", "", "MaintainedProperties")
	MakerOwnableID          = base.NewCLIFlag("makerOwnableID", "", "MakerOwnableID")
	MakerOwnableSplit       = base.NewCLIFlag("makerOwnableSplit", "", "MakerOwnableSplit")
	MutableMetaProperties   = base.NewCLIFlag("mutableMetaProperties", "", "mutableMetaProperties")
	MutableProperties       = base.NewCLIFlag("mutableProperties", "", "mutableProperties")
	MetaID                  = base.NewCLIFlag("metaID", "", "MetaID")
	MutateMaintainer        = base.NewCLIFlag("mutateMaintainer", false, "MutateMaintainer")
	NubID                   = base.NewCLIFlag("nubID", "", "NubID")
	OrderID                 = base.NewCLIFlag("orderID", "", "OrderID")
	OwnableID               = base.NewCLIFlag("ownableID", "", "MakerOwnableID")
	RemoveMaintainer        = base.NewCLIFlag("removeMaintainer", false, "RemoveMaintainer")
	Value                   = base.NewCLIFlag("value", "0", "Value")
	SplitID                 = base.NewCLIFlag("splitID", "", "SplitID")
	To                      = base.NewCLIFlag("to", "", "To")
	ToID                    = base.NewCLIFlag("toID", "", "ToID")
	TakerOwnableID          = base.NewCLIFlag("takerOwnableID", "", "TakerOwnableID")
	TakerOwnableSplit       = base.NewCLIFlag("takerOwnableSplit", "0", "TakerOwnableSplit")
)
