package genesis

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis Genesis) Default() helpers.Genesis {
	return genesis.Initialize(genesis.DefaultMappableList, genesis.DefaultParameterList)
}

func (genesis Genesis) Validate() error {
	if len(genesis.ParameterList) != len(genesis.DefaultParameterList) {
		return errors.InvalidParameter
	}

	for _, parameter := range genesis.ParameterList {
		var isPresent bool
		for _, defaultParameter := range genesis.DefaultParameterList {
			isPresent = false
			if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return errors.InvalidParameter
		}

		if Error := parameter.Validate(); Error != nil {
			return Error
		}
	}

	_, Error := govalidator.ValidateStruct(genesis)

	return Error
}

func (Genesis Genesis) Import(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) {
	for _, mappable := range Genesis.MappableList {
		mapper.Create(context, mappable)
	}

	for _, parameter := range Genesis.ParameterList {
		parameters.Mutate(context, parameter)
	}
}

func (Genesis Genesis) Export(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) helpers.Genesis {
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, key.Prototype(), appendMappableList)

	for _, defaultParameter := range Genesis.DefaultParameterList {
		parameters = parameters.Fetch(context, defaultParameter.GetID())
	}

	return Genesis.Initialize(mappableList, parameters.GetList())
}

func (Genesis Genesis) LegacyAminoEncode() []byte {
	legacyAminoCodec := codec.NewLegacyAmino()
	bytes, Error := legacyAminoCodec.MarshalJSON(Genesis)
	if Error != nil {
		panic(Error)
	}

	return bytes
}
func (Genesis Genesis) LegacyAminoDecode(byte []byte) helpers.Genesis {
	newGenesis := Genesis
	legacyAminoCodec := codec.NewLegacyAmino()
	if Error := legacyAminoCodec.UnmarshalJSON(byte, &newGenesis); Error != nil {
		panic(Error)
	}

	return NewGenesis(Genesis.DefaultMappableList, Genesis.DefaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}

func (Genesis Genesis) Encode(cdc codec.JSONMarshaler) []byte {

	bytes, Error := cdc.MarshalJSON(&Genesis)
	if Error != nil {
		panic(Error)
	}

	return bytes
}

func (Genesis Genesis) Decode(cdc codec.JSONMarshaler, byte []byte) helpers.Genesis {
	newGenesis := Genesis
	if Error := cdc.UnmarshalJSON(byte, &newGenesis); Error != nil {
		panic(Error)
	}

	return NewGenesis(Genesis.DefaultMappableList, Genesis.DefaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}

func (Genesis Genesis) Initialize(mappableList []helpers.Mappable, parameterList []types.Parameter) helpers.Genesis {
	if len(mappableList) == 0 {
		Genesis.MappableList = Genesis.DefaultMappableList
	} else {
		Genesis.MappableList = mappableList
	}

	if len(parameterList) == 0 {
		Genesis.ParameterList = Genesis.DefaultParameterList
	} else {
		for _, defaultParameter := range Genesis.DefaultParameterList {
			for i, parameter := range parameterList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					parameterList[i] = defaultParameter.Mutate(parameter.GetData())
				}
			}
		}
		Genesis.ParameterList = parameterList
	}

	if Error := Genesis.Validate(); Error != nil {
		panic(Error)
	}

	return &Genesis
}

func (Genesis Genesis) GetParameterList() []types.Parameter {
	return Genesis.ParameterList
}
func (Genesis Genesis) GetMappableList() []helpers.Mappable {
	return Genesis.MappableList
}

func NewGenesis(defaultMappableList []helpers.Mappable, defaultParameterList []types.Parameter) helpers.Genesis {
	return &Genesis{
		DefaultMappableList:  defaultMappableList,
		DefaultParameterList: defaultParameterList,
		MappableList:         []helpers.Mappable{},
		ParameterList:        []types.Parameter{},
	}
}
