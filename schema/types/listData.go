// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

type ListData interface {
	Data

	Search(Data) int

	GetList() []Data

	Add(...Data) ListData
	Remove(...Data) ListData
}
