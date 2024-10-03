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
func (genesis *Genesis) ValidateBasic(parameterManager helpers.ParameterManager) error {
	if err := parameterManager.ValidateGenesisParameters(genesis.ParameterList.Get()); err != nil {
		return err
	}

	for _, record := range genesis.Records {
		if err := record.GetMappable().ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (genesis *Genesis) Import(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) {
	for _, record := range genesis.Records {
		mapper.NewCollection(context).Add(record)
	}

	parameterManager.Set(context, genesis.ParameterList)
}
func (genesis *Genesis) Export(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) helpers.Genesis {
	return genesis.Initialize(mapper.NewCollection(context).FetchAll().Get(), parameterManager.Fetch(context).Get())
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
func (genesis *Genesis) Initialize(records []helpers.Record, parameterList lists.ParameterList) helpers.Genesis {
	if len(records) == 0 {
		genesis.Records = genesis.Default().(*Genesis).Records
	} else {
		genesis.Records = record.RecordsFromInterface(records)
	}

	if len(parameterList.Get()) == 0 {
		genesis.ParameterList = genesis.Default().(*Genesis).ParameterList
	} else {
		parameters := parameterList.Get()
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
		Records:       []*record.Record{},
		ParameterList: parameters.Prototype().Get().(*base.ParameterList),
	}
}
