// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	"github.com/AssetMantle/modules/schema/types"
)

type HeightData interface {
	Data
	Get() types.Height
}

type heightData dataSchema.HeightData

func (h heightData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (h heightData) String() string {
	//TODO implement me
	panic("implement me")
}

func (h heightData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (h heightData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (h heightData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (h heightData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (h heightData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (h heightData) Get() types.Height {
	//TODO implement me
	panic("implement me")
}

var _ HeightData = &heightData{}
