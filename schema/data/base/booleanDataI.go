package base

import dataSchema "buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/data"

type booleanData dataSchema.BooleanData

var _ BooleanData = &booleanData{}

func (b booleanData) GetID() ids.DataID {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) String() string {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) Bytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) GetType() ids.StringID {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) ZeroValue() Data {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) GenerateHashID() ids.HashID {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) Compare(listable traits.Listable) int {
	//TODO implement me
	panic("implement me")
}

func (b booleanData) Get() bool {
	//TODO implement me
	panic("implement me")
}
