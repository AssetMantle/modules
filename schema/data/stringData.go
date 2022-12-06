// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package data

import dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"

type StringData interface {
	Data
	Get() string
}

type stringData dataSchema.StringData

func (s stringData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (s stringData) String() string {
	//TODO implement me
	panic("implement me")
}

func (s stringData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (s stringData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (s stringData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (s stringData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (s stringData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (s stringData) Get() string {
	//TODO implement me
	panic("implement me")
}

var _ StringData = &stringData{}
