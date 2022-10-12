// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package constants

import (
	"github.com/AssetMantle/modules/schema/helpers/base"
)

const (
	assets int8 = iota + 8
	classifications
	identities
	maintainers
	metas
	orders
	splits
)

// TODO convert to enum
var (
	AssetsStoreKeyPrefix          = base.NewStoreKeyPrefix(assets)
	ClassificationsStoreKeyPrefix = base.NewStoreKeyPrefix(classifications)
	IdentitiesStoreKeyPrefix      = base.NewStoreKeyPrefix(identities)
	MaintainersStoreKeyPrefix     = base.NewStoreKeyPrefix(maintainers)
	MetasStoreKeyPrefix           = base.NewStoreKeyPrefix(metas)
	OrdersStoreKeyPrefix          = base.NewStoreKeyPrefix(orders)
	SplitsStoreKeyPrefix          = base.NewStoreKeyPrefix(splits)
)
