/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	metaUtilities "github.com/persistenceOne/persistenceSDK/utilities/meta"
)

// TODO implement Json, URI, interface{} metadata
type text struct {
	ID   types.ID `json:"id" valid:"required field id missing"`
	Data string   `json:"data" valid:"required field data missing"`
}

var _ mappables.Meta = (*text)(nil)

func (text text) GetID() types.ID { return text.ID }
func (text text) Get() string     { return text.Data }
func (text text) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(text)
}
func (text text) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &text)
	return text
}
func metaPrototype() traits.Mappable {
	return text{}
}
func NewMeta(data string) mappables.Meta {
	return text{
		ID:   base.NewID(metaUtilities.Hash(data)),
		Data: data,
	}
}
