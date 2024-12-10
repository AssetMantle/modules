package genesis

import (
	"context"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
	"github.com/AssetMantle/modules/x/maintainers/record"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis *Genesis) Default() helpers.Genesis {
	return Prototype()
}
func (genesis *Genesis) GetRecords() []helpers.Record {
	return helpers.RecordsFromImplementations(genesis.Records)
}
func (genesis *Genesis) GetParameterList() lists.ParameterList {
	return genesis.ParameterList
}
func (genesis *Genesis) SetRecords(records []helpers.Record) helpers.Genesis {
	genesis.Records = helpers.RecordsToImplementations(&record.Record{}, records)
	return genesis
}
func (genesis *Genesis) SetParameters(parameterList lists.ParameterList) helpers.Genesis {
	genesis.ParameterList = parameterList.(*base.ParameterList)
	return genesis
}
func (genesis *Genesis) ValidateBasic(parameterManager helpers.ParameterManager) error {
	return helpers.ValidateGenesis(genesis, parameterManager)
}
func (genesis *Genesis) Import(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) {
	helpers.ImportGenesis(genesis, context, mapper, parameterManager)
}
func (genesis *Genesis) Export(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) helpers.Genesis {
	return helpers.ExportGenesis(genesis, context, mapper, parameterManager)
}
func (genesis *Genesis) Encode(jsonCodec sdkCodec.JSONCodec) []byte {
	return helpers.EncodeGenesis(genesis, jsonCodec)
}
func (genesis *Genesis) Decode(jsonCodec sdkCodec.JSONCodec, byte []byte) helpers.Genesis {
	return helpers.DecodeGenesis(genesis, jsonCodec, byte)
}
func (genesis *Genesis) Initialize(records []helpers.Record, parameterList lists.ParameterList) helpers.Genesis {
	return helpers.InitializeGenesis(genesis, records, parameterList)
}

func Prototype() helpers.Genesis {
	return &Genesis{
		Records:       []*record.Record{},
		ParameterList: parameters.Prototype().GetDefaultParameterList().(*base.ParameterList),
	}
}
