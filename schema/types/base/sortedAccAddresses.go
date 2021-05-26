/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"bytes"
	"sort"

	"github.com/persistenceOne/persistenceSDK/constants/errors"

	"github.com/persistenceOne/persistenceSDK/schema/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type sortedAccAddresses []sdkTypes.AccAddress

var _ types.SortedList = (*sortedAccAddresses)(nil)

func (accAddresses sortedAccAddresses) Len() int {
	return len(accAddresses)
}
func (accAddresses sortedAccAddresses) Less(i, j int) bool {
	return bytes.Compare(accAddresses[i], accAddresses[j]) < 0
}
func (accAddresses sortedAccAddresses) Swap(i, j int) {
	accAddresses[i], accAddresses[j] = accAddresses[j], accAddresses[i]
}
func (accAddresses sortedAccAddresses) Sort() types.SortedList {
	sort.Sort(accAddresses)
	return accAddresses
}
func (accAddresses sortedAccAddresses) Insert(i interface{}) types.SortedList {
	accAddress := accAddressFromInterface(i)
	if accAddresses.Search(accAddress) != accAddresses.Len() {
		return accAddresses
	}

	index := sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Compare(accAddresses[i].Bytes(), accAddress.Bytes()) < 0
		},
	)

	newAccAddresses := append(accAddresses, sdkTypes.AccAddress{})
	copy(newAccAddresses[index+1:], newAccAddresses[index:])
	newAccAddresses[index] = accAddress

	return newAccAddresses
}
func (accAddresses sortedAccAddresses) Delete(i interface{}) types.SortedList {
	accAddress := accAddressFromInterface(i)
	index := accAddresses.Search(accAddress)

	if index == accAddresses.Len() {
		return accAddresses
	}

	return append(accAddresses[:index], accAddresses[index+1:]...)
}
func (accAddresses sortedAccAddresses) Search(i interface{}) int {
	accAddress := accAddressFromInterface(i)

	return sort.Search(
		accAddresses.Len(),
		func(i int) bool {
			return bytes.Equal(accAddresses[i].Bytes(), accAddress.Bytes())
		},
	)
}
func accAddressFromInterface(i interface{}) sdkTypes.AccAddress {
	switch value := i.(type) {
	case sdkTypes.AccAddress:
		return value
	default:
		panic(errors.IncorrectFormat)
	}
}
