package genesis

import (
	"context"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/lists/base"
	parametersSchema "github.com/AssetMantle/schema/parameters"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/parameters"
	"github.com/AssetMantle/modules/x/classifications/record"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis *Genesis) Default() helpers.Genesis {
	return Prototype()
}
func (genesis *Genesis) GetRecords() []helpers.Record {
	return helpers.RecordsFromImplementations(genesis.Records)
}
func (genesis *Genesis) GetParameters() []parametersSchema.Parameter {
	return genesis.ParameterList.Get()
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
	bytes, err := jsonCodec.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return bytes
}
func (genesis *Genesis) Decode(jsonCodec sdkCodec.JSONCodec, byte []byte) helpers.Genesis {
	if err := jsonCodec.UnmarshalJSON(byte, genesis); err != nil {
		panic(err)
	}

	return genesis
}
func (genesis *Genesis) Initialize(records []helpers.Record, parameters []parametersSchema.Parameter) helpers.Genesis {
	if len(records) == 0 {
		genesis.Records = genesis.Default().(*Genesis).Records
	} else {
		genesis.Records = record.RecordsFromInterface(records)
	}

	if len(parameters) == 0 {
		genesis.ParameterList = genesis.Default().(*Genesis).ParameterList
	} else {
		for _, defaultParameter := range genesis.Default().(*Genesis).ParameterList.Get() {
			for i, parameter := range parameters {
				if defaultParameter.GetMetaProperty().GetID().Compare(parameter.GetMetaProperty().GetID()) == 0 {
					parameters[i] = defaultParameter.Mutate(parameter.GetMetaProperty().GetData())
				}
			}
		}
		genesis.ParameterList = base.NewParameterList(parameters...).(*base.ParameterList)
	}

	return genesis
}

func Prototype() helpers.Genesis {
	return &Genesis{
		Records: []*record.Record{
			record.NewRecord(baseDocuments.NewClassificationFromDocument(baseDocuments.PrototypeCoinAsset())).(*record.Record),
			record.NewRecord(baseDocuments.NewClassificationFromDocument(baseDocuments.PrototypeNameIdentity())).(*record.Record),
			record.NewRecord(baseDocuments.NewClassificationFromDocument(baseDocuments.PrototypeModuleIdentity())).(*record.Record),
			record.NewRecord(baseDocuments.NewClassificationFromDocument(baseDocuments.PrototypeMaintainer())).(*record.Record),
			record.NewRecord(baseDocuments.NewClassificationFromDocument(baseDocuments.PrototypePutOrder())).(*record.Record),
		},
		ParameterList: base.NewParameterList(parameters.Prototype().Get()...).(*base.ParameterList),
	}
}
