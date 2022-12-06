// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import (
	dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type DecData interface {
	Data
	Get() sdkTypes.Dec
}

type decData dataSchema.DecData

func (d decData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (d decData) String() string {
	//TODO implement me
	panic("implement me")
}

func (d decData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (d decData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (d decData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (d decData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (d decData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (d decData) Get() sdkTypes.Dec {
	//TODO implement me
	panic("implement me")
}

var _ DecData = &decData{}
