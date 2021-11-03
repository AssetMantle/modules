/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import "github.com/gogo/protobuf/proto"

type MetaProperty interface {
	proto.Message

	GetID() ID
	GetMetaFact() MetaFact
	ToProperty() Property
}
