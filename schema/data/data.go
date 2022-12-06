// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	data2 "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"

	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO URI and PropertyID data type
type Data interface {
	GetID() ids.DataID

	String() string
	Bytes() []byte

	GetType() ids.StringID
	ZeroValue() Data
	// GenerateHash returns the hash of the Data as an PropertyID
	// * Returns PropertyID of empty bytes when the value of Data is that Data type's zero value
	GenerateHashID() ids.HashID

	traits.Listable
}

type data data2.Data

var _ Data = (*data)(nil)

func (x *data) String() string {
	return x.Impl.(Data).String()
}
func (x *data) GetID() ids.DataID {
	return x.Impl.(Data).GetID()
}

func (x *data) Bytes() []byte {
	return x.Impl.(Data).Bytes()
}

func (x *data) GetType() ids.StringID {
	return x.Impl.(Data).GetType()
}

func (x *data) ZeroValue() Data {
	return x.Impl.(Data).ZeroValue()
}

func (x *data) GenerateHashID() ids.HashID {
	return x.Impl.(Data).GenerateHashID()
}

func (x *data) Compare(listable traits.Listable) int {
	return x.Impl.(Data).Compare(listable)
}
