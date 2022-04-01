// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"github.com/persistenceOne/persistenceSDK/schema/traits"
)

type List interface {
	GetList() []traits.Listable

	Search(func()) int

	Apply(func()) List
	Add(...interface{}) List
	Remove(...interface{}) List
	Mutate(...interface{}) List
}
