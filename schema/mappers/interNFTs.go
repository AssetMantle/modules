/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappers

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type InterNFTs interface {
	GetID() types.ID

	Get(types.ID) mappables.InterNFT
	GetList() []mappables.InterNFT

	Fetch(types.ID) InterNFTs
	Add(mappables.InterNFT) InterNFTs
	Remove(mappables.InterNFT) InterNFTs
	Mutate(mappables.InterNFT) InterNFTs
}
