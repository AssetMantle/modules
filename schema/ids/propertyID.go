// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ids

type PropertyID interface {
	GetKey() ID
	GetType() ID
	ID
	IsPropertyID()
	PropertyIDString() string
}
