// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

type Transactions interface {
	Get(string) Transaction
	GetList() []Transaction
}
