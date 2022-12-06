// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

type ListData interface {
	Get() []Data

	Search(Data) (int, bool)
	Add(...Data) ListData
	Remove(...Data) ListData

	Data
}
