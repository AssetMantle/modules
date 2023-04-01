// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

import (
	"github.com/AssetMantle/modules/schema/ids"
)

type IDList interface {
	GetList() []ids.AnyID
	Search(ids.ID) (index int, found bool)
	Add(...ids.ID) IDList
	Remove(...ids.ID) IDList
	ValidateBasic() error
}
