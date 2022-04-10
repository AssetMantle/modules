// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

type MetaProperty interface {
	GetData() Data
	RemoveData() Property
	Property
}
