/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import codecTypes "github.com/cosmos/cosmos-sdk/codec/types"

type ListData interface {
	Data

	Search(Data) int

	GetList() []Data

	Add(...Data) ListData
	Remove(...Data) ListData
	codecTypes.UnpackInterfacesMessage
}
