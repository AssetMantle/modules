package base

import (
	"context"

	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseParameters "github.com/AssetMantle/modules/schema/parameters/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis *Genesis) Default() helpers.Genesis {
	return Prototype()
}
func (genesis *Genesis) ValidateBasic() error {
	if len(genesis.Parameters) != len(genesis.Default().(*Genesis).Parameters) {
		return constants.InvalidParameter
	}

	for _, parameter := range genesis.Parameters {
		var isPresent bool
		for _, defaultParameter := range genesis.Default().(*Genesis).Parameters {
			isPresent = false
			if defaultParameter.GetMetaProperty().Compare(parameter.GetMetaProperty()) == 0 {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return constants.InvalidParameter
		}

		if err := parameter.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (genesis *Genesis) Import(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) {
	for _, Mappable := range genesis.Mappables {
		mapper.Create(context, Mappable)
	}

	parameterManager.Set(context, baseParameters.ParametersToInterface(genesis.Parameters)...)
}
func (genesis *Genesis) Export(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) helpers.Genesis {
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.IterateAll(context, appendMappableList)

	return genesis.Initialize(mappableList, parameterManager.Fetch(context).Get())
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
func (genesis *Genesis) Initialize(mappables []helpers.Mappable, parameters []helpers.Parameter) helpers.Genesis {
	if len(mappables) == 0 {
		genesis.Mappables = genesis.Default().(*Genesis).Mappables
	} else {
		genesis.Mappables = MappablesFromInterface(mappables)
	}

	if len(parameters) == 0 {
		genesis.Parameters = genesis.Default().(*Genesis).Parameters
	} else {
		for _, defaultParameter := range genesis.Default().(*Genesis).Parameters {
			for i, parameter := range parameters {
				if defaultParameter.GetMetaProperty().GetID().Compare(parameter.GetMetaProperty().GetID()) == 0 {
					parameters[i] = defaultParameter.Mutate(parameter.GetMetaProperty().GetData())
				}
			}
		}
		genesis.Parameters = baseParameters.ParametersFromInterface(parameters)
	}

	return genesis
}

func Prototype() helpers.Genesis {
	return &Genesis{
		Mappables:  []*TestMappable{},
		Parameters: []*baseParameters.Parameter{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("testParameter"), baseData.NewStringData("testData")))},
	}
}
