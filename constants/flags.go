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
	AssetID          = base.NewCLIFlag("assetID", "", "AssetID")
	Burn             = base.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = base.NewCLIFlag("classificationID", "", "ClassificationID")
	FromID           = base.NewCLIFlag("fromID", "", "FromID")
	IdentityID       = base.NewCLIFlag("identityID", "", "IdentityID")
	Lock             = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = base.NewCLIFlag("maintainersID", "", "MaintainersID")
	OwnableID        = base.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = base.NewCLIFlag("properties", "", "Properties")
	Split            = base.NewCLIFlag("split", "0", "Split")
	SplitID          = base.NewCLIFlag("splitID", "", "SplitID")
	Coins            = base.NewCLIFlag("coins", "", "Coins")
	Traits           = base.NewCLIFlag("traits", "", "Traits")
	To               = base.NewCLIFlag("to", "", "To")
	ToID             = base.NewCLIFlag("toID", "", "ToID")
	OrderID          = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress     = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	MaintainerID     = base.NewCLIFlag("makerSplit", int64(0), "MaintainerID")
	MakerSplit       = base.NewCLIFlag("makerSplit", int64(0), "MakerSplit")
	MakerSplitID     = base.NewCLIFlag("makerSplitID", "", "MakerSplitID")
	Meta             = base.NewCLIFlag("meta", "", "Meta")
	ExchangeRate     = base.NewCLIFlag("exchangeRate", "1", "ExchangeRate")
	TakerSplit       = base.NewCLIFlag("takerSplit", int64(0), "TakerSplit")
	TakerSplitID     = base.NewCLIFlag("takerSplitID", "", "TakerSplitID")
)
