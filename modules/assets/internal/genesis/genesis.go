package genesis

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mappable"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters/dummy"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis Genesis) Default() helpers.Genesis {
	return genesis.Initialize(genesis.GetMappableList(), genesis.GetParameterList())
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
	for _, mappableValue := range genesis.MappableList {
		mapper.Create(context, &mappableValue)
	}

	for _, parameter := range genesis.ParameterList {
		parameters.Mutate(context, &parameter)
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

	return NewGenesis(newGenesis.DefaultMappableList, newGenesis.DefaultParameterList).Initialize(newGenesis.GetMappableList(), newGenesis.GetParameterList())
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

	return NewGenesis(newGenesis.DefaultMappableList, newGenesis.DefaultParameterList).Initialize(newGenesis.GetMappableList(), newGenesis.GetParameterList())
}

func (genesis Genesis) Initialize(mappableList []helpers.Mappable, parameterList []types.Parameter) helpers.Genesis {
	newParametersList := make([]dummy.DummyParameter, len(parameterList))
	for i, _ := range parameterList {
		newParametersList[i] = *dummy.NewParameter(parameterList[i].GetID(), parameterList[i].GetData())
	}
	newMappableList := make([]mappable.Asset, len(mappableList))
	for i, _ := range mappableList {
		newMappableList[i] = *mappableList[i].(*mappable.Asset)
	}
	newParameter := dummy.Parameter.Mutate(dummy.Parameter.GetData())
	genesis.DefaultParameterList = []dummy.DummyParameter{*dummy.NewParameter(newParameter.GetID(), newParameter.GetData())}
	if len(newMappableList) == 0 {
		genesis.MappableList = genesis.DefaultMappableList
	} else {
		genesis.MappableList = newMappableList
	}

	if len(newParametersList) == 0 {
		genesis.ParameterList = genesis.DefaultParameterList
	} else {
		for _, defaultParameter := range genesis.DefaultParameterList {
			for i, parameter := range newParametersList {
				if defaultParameter.GetID().Compare(parameter.GetID()) == 0 {
					newParametersList[i] = *dummy.NewParameter(defaultParameter.Mutate(parameter.GetData()).GetID(), defaultParameter.Mutate(parameter.GetData()).GetData())
				}
			}
		}
		genesis.ParameterList = newParametersList
	}

	if Error := genesis.Validate(); Error != nil {
		panic(Error)
	}

	return &genesis
}

func (genesis Genesis) GetParameterList() []types.Parameter {
	newParameterList := make([]types.Parameter, len(genesis.ParameterList))
	for i, _ := range genesis.ParameterList {
		newParameterList[i] = &genesis.ParameterList[i]
	}
	return newParameterList
}
func (genesis Genesis) GetMappableList() []helpers.Mappable {
	newMappableList := make([]helpers.Mappable, len(genesis.MappableList))
	for i, _ := range genesis.MappableList {
		newMappableList[i] = &genesis.MappableList[i]
	}
	return newMappableList
}

func (genesis Genesis) RegisterInterface(registry codecTypes.InterfaceRegistry) {
	registry.RegisterImplementations((*helpers.Key)(nil),
		&key.AssetID{},
	)
	registry.RegisterImplementations((*helpers.Mappable)(nil),
		&mappable.Asset{},
	)
	registry.RegisterImplementations((*types.Parameter)(nil),
		&dummy.DummyParameter{},
	)
	registry.RegisterImplementations((*helpers.Genesis)(nil),
		&Genesis{},
	)
}

func NewGenesis(defaultMappableList []mappable.Asset, defaultParameterList []dummy.DummyParameter) *Genesis {
	return &Genesis{
		DefaultMappableList:  defaultMappableList,
		DefaultParameterList: defaultParameterList,
		MappableList:         []mappable.Asset{},
		ParameterList:        []dummy.DummyParameter{},
	}
}
