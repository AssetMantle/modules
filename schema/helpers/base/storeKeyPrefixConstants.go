// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

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
	AssetsStoreKeyPrefix          = NewStoreKeyPrefix(assets)
	ClassificationsStoreKeyPrefix = NewStoreKeyPrefix(classifications)
	IdentitiesStoreKeyPrefix      = NewStoreKeyPrefix(identities)
	MaintainersStoreKeyPrefix     = NewStoreKeyPrefix(maintainers)
	MetasStoreKeyPrefix           = NewStoreKeyPrefix(metas)
	OrdersStoreKeyPrefix          = NewStoreKeyPrefix(orders)
	SplitsStoreKeyPrefix          = NewStoreKeyPrefix(splits)
)
