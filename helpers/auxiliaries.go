// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

type Auxiliaries interface {
	Get() []Auxiliary
	GetAuxiliary(string) Auxiliary
}
