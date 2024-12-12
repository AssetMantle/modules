// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/lists"
	"github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/helpers"
)

// TODO: move to a constants file
var parametersStorePrefix = []byte{0x01}

type parameterManager struct {
	kvStoreKey                   *storeTypes.KVStoreKey
	parameterList                lists.ParameterList
	defaultValidatableParameters []helpers.ValidatableParameter
}

var _ helpers.ParameterManager = (*parameterManager)(nil)

func (parameterManager parameterManager) Get() lists.ParameterList {
	return parameterManager.parameterList
}
func (parameterManager parameterManager) GetDefaultParameterList() lists.ParameterList {
	defaultParameters := make([]parameters.Parameter, len(parameterManager.defaultValidatableParameters))

	for i, validatableParameter := range parameterManager.defaultValidatableParameters {
		defaultParameters[i] = validatableParameter.GetParameter()
	}

	return base.NewParameterList(defaultParameters...)
}
func (parameterManager parameterManager) Set(parameters ...parameters.Parameter) helpers.ParameterManager {
	if parameterManager.parameterList == nil {
		parameterManager.parameterList = parameterManager.GetDefaultParameterList()
	}

	parameterManager.parameterList = parameterManager.parameterList.Mutate(parameters...)
	return parameterManager
}
func (parameterManager parameterManager) Validate() error {
	if len(parameterManager.defaultValidatableParameters) != len(parameterManager.parameterList.Get()) {
		return fmt.Errorf("parameter count mismatch, wanted %d, got %d", len(parameterManager.defaultValidatableParameters), len(parameterManager.parameterList.Get()))
	}

	if err := parameterManager.parameterList.ValidateBasic(); err != nil {
		return err
	}

	for _, validatableParameter := range parameterManager.defaultValidatableParameters {
		parameter := parameterManager.parameterList.GetParameter(validatableParameter.GetParameter().GetMetaProperty().GetID())
		if parameter == nil {
			return fmt.Errorf("parameter with id %s not found", validatableParameter.GetParameter().GetMetaProperty().GetID())
		}

		if err := validatableParameter.Mutate(parameter.GetMetaProperty().GetData()).Validate(); err != nil {
			return err
		}
	}

	return nil
}
func (parameterManager parameterManager) Fetch(context context.Context) helpers.ParameterManager {
	parameterList := base.PrototypeParameterList()
	CodecPrototype().MustUnmarshal(sdkTypes.UnwrapSDKContext(context).KVStore(parameterManager.kvStoreKey).Get(parametersStorePrefix), parameterList)
	parameterManager.parameterList = parameterList

	return parameterManager
}
func (parameterManager parameterManager) Update(context context.Context) (helpers.ParameterManager, error) {
	if err := parameterManager.Validate(); err != nil {
		return nil, err
	}

	sdkTypes.UnwrapSDKContext(context).KVStore(parameterManager.kvStoreKey).Set(parametersStorePrefix, CodecPrototype().MustMarshal(parameterManager.parameterList))

	return parameterManager, nil
}

func (parameterManager parameterManager) Initialize(kvStoreKey *storeTypes.KVStoreKey) helpers.ParameterManager {
	parameterManager.kvStoreKey = kvStoreKey
	return parameterManager
}

func NewParameterManager(defaultValidatableParameters ...helpers.ValidatableParameter) helpers.ParameterManager {
	return parameterManager{
		defaultValidatableParameters: defaultValidatableParameters,
	}
}
