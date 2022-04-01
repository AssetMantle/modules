// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

// TODO do sortable interface
type Property interface {
	GetID() ID
	GetDataID() ID
	GetKeyID() ID
	GetTypeID() ID
	GetHashID() ID
}
