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
	AssetID          = base.NewCLIFlag("assetID", "", "OwnableID")
	Burn             = base.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = base.NewCLIFlag("classificationID", "", "ClassificationID")
	FromID           = base.NewCLIFlag("fromID", "", "MakerID")
	IdentityID       = base.NewCLIFlag("identityID", "", "OwnerID")
	Lock             = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = base.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID        = base.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = base.NewCLIFlag("properties", "", "GetProperties")
	Split            = base.NewCLIFlag("split", "", "Split")
	SplitID          = base.NewCLIFlag("splitID", "", "SplitID")
	Traits           = base.NewCLIFlag("traits", "", "Traits")
	To               = base.NewCLIFlag("to", "", "To")
	ToID             = base.NewCLIFlag("toID", "", "TakerID")
	OrderID          = base.NewCLIFlag("orderID", "", "OrderID")
	TakerAddress     = base.NewCLIFlag("takerAddress", "", "TakerAddress")
	MakerSplit       = base.NewCLIFlag("makerSplit", int64(0), "MakerSplit")
	MakerSplitID     = base.NewCLIFlag("makerSplitID", "", "MakerSplitID")
	ExchangeRate     = base.NewCLIFlag("exchangeRate", "1", "ExchangeRate")
	TakerSplit       = base.NewCLIFlag("takerSplit", int64(0), "TakerSplit")
	TakerSplitID     = base.NewCLIFlag("takerSplitID", "", "TakerSplitID")
)
