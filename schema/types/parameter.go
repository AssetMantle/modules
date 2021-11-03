/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

type Parameter interface {
	proto.Message

	String() string

	Equal(Parameter) bool

	GetID() ID
	GetData() Data

	SetData(Data) error

	codecTypes.UnpackInterfacesMessage
}
