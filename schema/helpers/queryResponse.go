/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import "github.com/persistenceOne/persistenceSDK/schema/types"

type QueryResponse interface {
	Response
	LegacyAminoEncode() ([]byte, error)
	LegacyAminoDecode([]byte) (QueryResponse, error)
	Encode() ([]byte, error)
	Decode([]byte) (QueryResponse, error)
	types.Proto
	Reset()
	String() string
	ProtoMessage()
}

