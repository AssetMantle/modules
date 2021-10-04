package genesis

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/parameters/dummy"
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

func (genesis Genesis) Import(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) {
	for _, mappable := range genesis.MappableList {
		mapper.Create(context, mappable)
	}

	for _, parameter := range genesis.ParameterList {
		parameters.Mutate(context, parameter)
	}
}

func (genesis Genesis) Export(context sdkTypes.Context, mapper helpers.Mapper, parameters helpers.Parameters) helpers.Genesis {
	var mappableList []helpers.Mappable

	appendMappableList := func(mappable helpers.Mappable) bool {
		mappableList = append(mappableList, mappable)
		return false
	}
	mapper.Iterate(context, key.Prototype(), appendMappableList)

	for _, defaultParameter := range genesis.DefaultParameterList {
		parameters = parameters.Fetch(context, defaultParameter.GetID())
	}

	return genesis.Initialize(mappableList, parameters.GetList())
}

func (genesis Genesis) LegacyAminoEncode() []byte {
	legacyAminoCodec := codec.NewLegacyAmino()
	bytes, Error := legacyAminoCodec.MarshalJSON(genesis)
	if Error != nil {
		panic(Error)
	}

	return bytes
}
func (genesis Genesis) LegacyAminoDecode(byte []byte) helpers.Genesis {
	var newGenesis Genesis
	legacyAminoCodec := codec.NewLegacyAmino()
	if Error := legacyAminoCodec.UnmarshalJSON(byte, &newGenesis); Error != nil {
		panic(Error)
	}

	return NewGenesis(genesis.DefaultMappableList, genesis.DefaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}

func (genesis Genesis) Encode(cdc codec.JSONMarshaler) []byte {

	bytes, Error := cdc.MarshalJSON(&genesis)
	if Error != nil {
		panic(Error)
	}

	return bytes
}

func (genesis Genesis) Decode(cdc codec.JSONMarshaler, byte []byte) helpers.Genesis {
	var newGenesis Genesis
	if Error := cdc.UnmarshalJSON(byte, &newGenesis); Error != nil {
		panic(Error)
	}

	return NewGenesis(genesis.DefaultMappableList, genesis.DefaultParameterList).Initialize(newGenesis.MappableList, newGenesis.ParameterList)
}

func (genesis Genesis) Initialize(mappableList []helpers.Mappable, parameterList []types.Parameter) helpers.Genesis {
	genesis.DefaultParameterList = []types.Parameter{dummy.Parameter.Mutate(dummy.Parameter.GetData())}
	if len(mappableList) == 0 {
		genesis.MappableList = genesis.DefaultMappableList
	} else {
		genesis.MappableList = mappableList
	}

	if len(parameterList) == 0 {
		genesis.ParameterList = genesis.DefaultParameterList
	} else {
		for _, defaultParameter := range genesis.DefaultParameterList {
			for i, parameter := range parameterList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					parameterList[i] = defaultParameter.Mutate(parameter.GetData())
				}
			}
		}
		genesis.ParameterList = parameterList
	}

	if Error := genesis.Validate(); Error != nil {
		panic(Error)
	}

	return &genesis
}

func (genesis Genesis) GetParameterList() []types.Parameter {
	return genesis.ParameterList
}
func (genesis Genesis) GetMappableList() []helpers.Mappable {
	return genesis.MappableList
}

func (genesis Genesis) RegisterInterface(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations((*helpers.Key)(nil),
		&key.SplitID{},
	)
	registry.RegisterImplementations((*helpers.Mappable)(nil),
		&mappable.Split{},
	)
	registry.RegisterImplementations((*types.Parameter)(nil),
		&dummy.DummyParameter{},
	)
	registry.RegisterImplementations((*helpers.Genesis)(nil),
		&Genesis{},
	)
}

func NewGenesis(defaultMappableList []helpers.Mappable, defaultParameterList []types.Parameter) helpers.Genesis {
	return &Genesis{
		DefaultMappableList:  defaultMappableList,
		DefaultParameterList: defaultParameterList,
		MappableList:         []helpers.Mappable{},
		ParameterList:        []types.Parameter{},
	}
}
