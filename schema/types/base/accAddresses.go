/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"sort"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type accAddresses []sdkTypes.AccAddress

var _ sort.Interface = (*accAddresses)(nil)

func (accAddresses accAddresses) Len() int {
	return len(accAddresses)
}

func (accAddresses accAddresses) Less(i, j int) bool {
	return bytes.Compare(accAddresses[i], accAddresses[j]) < 0
}

func (accAddresses accAddresses) Swap(i, j int) {
	accAddresses[i], accAddresses[j] = accAddresses[j], accAddresses[i]
}
