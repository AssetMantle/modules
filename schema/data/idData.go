// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type IDData interface {
	Data
	Get() ids.ID
}

type idData dataSchema.IdData

func (i idData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (i idData) String() string {
	//TODO implement me
	panic("implement me")
}

func (i idData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (i idData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (i idData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (i idData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (i idData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (i idData) Get() ids.ID {
	//TODO implement me
	panic("implement me")
}

var _ IDData = &idData{}
