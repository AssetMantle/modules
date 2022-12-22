// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

type ListData interface {
	Data
	Get() []AnyData
	Search(AnyData) (int, bool)
	Add(...AnyData) ListData
	Remove(...AnyData) ListData
}
