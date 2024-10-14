package helpers

import (
	"context"
	parametersSchema "github.com/AssetMantle/schema/parameters"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/gogoproto/proto"
)

type Genesis interface {
	GetRecords() []Record
	GetParameters() []parametersSchema.Parameter

	SetRecords([]Record) Genesis
	SetParameters([]parametersSchema.Parameter) Genesis

	Default() Genesis

	ValidateBasic(ParameterManager) error

	Import(context.Context, Mapper, ParameterManager)
	Export(context.Context, Mapper, ParameterManager) Genesis

	Encode(sdkCodec.JSONCodec) []byte
	Decode(sdkCodec.JSONCodec, []byte) Genesis

	Initialize([]Record, []parametersSchema.Parameter) Genesis

	proto.Message
}

func ValidateGenesis[T Genesis](genesis T, parameterManager ParameterManager) error {
	if err := parameterManager.ValidateGenesisParameters(genesis); err != nil {
		return err
	}

	for _, record := range genesis.GetRecords() {
		if err := record.GetMappable().ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}

func ImportGenesis[T Genesis](genesis T, context context.Context, mapper Mapper, parameterManager ParameterManager) {
	for _, record := range genesis.GetRecords() {
		mapper.NewCollection(context).Add(record)
	}

	parameterManager.Set(context, genesis.GetParameters())
}

func ExportGenesis[T Genesis](genesis T, context context.Context, mapper Mapper, parameterManager ParameterManager) Genesis {
	return genesis.Initialize(mapper.NewCollection(context).FetchAll().Get(), parameterManager.Fetch(context).Get())
}

func EncodeGenesis[T Genesis](genesis T, jsonCodec sdkCodec.JSONCodec) []byte {
	bytes, err := jsonCodec.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return bytes
}
func DecodeGenesis[T Genesis](genesis T, jsonCodec sdkCodec.JSONCodec, byte []byte) Genesis {
	if err := jsonCodec.UnmarshalJSON(byte, genesis); err != nil {
		panic(err)
	}

	return genesis
}

func InitializeGenesis[T Genesis](genesis T, records []Record, parameters []parametersSchema.Parameter) Genesis {
	if len(records) == 0 {
		records = genesis.Default().GetRecords()
	}

	if len(parameters) == 0 {
		parameters = genesis.Default().GetParameters()
	} else {
		providedParamsMap := make(map[string]parametersSchema.Parameter)
		for _, parameter := range parameters {
			providedParamsMap[parameter.GetMetaProperty().GetID().AsString()] = parameter
		}

		defaultParameters := genesis.Default().GetParameters()
		for i, defaultParameter := range defaultParameters {
			if providedParameter, exists := providedParamsMap[defaultParameter.GetMetaProperty().GetID().AsString()]; exists {
				defaultParameters[i] = defaultParameter.Mutate(providedParameter.GetMetaProperty().GetData())
			}
		}

		parameters = defaultParameters
	}

	return genesis.SetRecords(records).SetParameters(parameters)
}
