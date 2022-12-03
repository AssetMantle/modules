// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

type ListData interface {
	Get() []DataI

	Search(DataI) (int, bool)
	Add(...DataI) ListData
	Remove(...DataI) ListData

	DataI
}
