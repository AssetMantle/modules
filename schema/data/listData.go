// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type ListData interface {
	Get() []Data

	Search(Data) (int, bool)
	Add(...Data) ListData
	Remove(...Data) ListData

	Data
}

type listData dataSchema.ListData

func (l listData) Get() []Data {
	//TODO implement me
	panic("implement me")
}

func (l listData) Search(d Data) (int, bool) {
	//TODO implement me
	panic("implement me")
}

func (l listData) Add(d ...Data) ListData {
	//TODO implement me
	panic("implement me")
}

func (l listData) Remove(d ...Data) ListData {
	//TODO implement me
	panic("implement me")
}

func (l listData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (l listData) String() string {
	//TODO implement me
	panic("implement me")
}

func (l listData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (l listData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (l listData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (l listData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (l listData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

var _ ListData = &listData{}
