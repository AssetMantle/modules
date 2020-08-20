/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package constants

import "github.com/persistenceOne/persistenceSDK/schema/types/base"

//Note: Arranged alphabetically
var (
	BurnProperty              = base.NewID("burn")
	CreationProperty          = base.NewID("creationProperty")
	ExpiryProperty            = base.NewID("expiryProperty")
	LockProperty              = base.NewID("lock")
	MakerIDProperty           = base.NewID("makerID")
	MakerSplitIDProperty      = base.NewID("makerSplitID")
	MakerSplitProperty        = base.NewID("makerSplit")
	MakerOwnableSplitProperty = base.NewID("makerOwnableSplit")
	TakerIDProperty           = base.NewID("takerID")
	ExchangeRateProperty      = base.NewID("exchangeRate")
	TakerSplitIDProperty      = base.NewID("takerSplitID")
)
