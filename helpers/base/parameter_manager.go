// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"github.com/AssetMantle/schema/ids"
	"github.com/AssetMantle/schema/parameters"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/helpers"
)

type parameterManager struct {
	validatableParameters []helpers.ValidatableParameter
	paramsSubspace        paramsTypes.Subspace
}

var _ helpers.ParameterManager = (*parameterManager)(nil)

func (parameterManager parameterManager) Get() []parameters.Parameter {
	Parameters := make([]parameters.Parameter, len(parameterManager.validatableParameters))
	for i, validatableParameter := range parameterManager.validatableParameters {
		Parameters[i] = validatableParameter.GetParameter()
	}
	return Parameters
}
func (parameterManager parameterManager) GetParameter(propertyID ids.PropertyID) parameters.Parameter {
	if validatableParameter := parameterManager.GetValidatableParameter(propertyID); validatableParameter != nil {
		return validatableParameter.GetParameter()
	}
	return nil
}
func (parameterManager parameterManager) GetValidatableParameter(propertyID ids.PropertyID) helpers.ValidatableParameter {
	for _, validatableParameter := range parameterManager.validatableParameters {
		if validatableParameter.GetParameter().GetMetaProperty().GetID().Compare(propertyID) == 0 {
			return validatableParameter
		}
	}
	return nil
}
func (parameterManager parameterManager) ValidateGenesisParameters(genesis helpers.Genesis) error {
	if len(genesis.GetParameters()) != len(parameterManager.validatableParameters) {
		return fmt.Errorf("genesis parameters length mismatch")
	}

	for _, parameter := range genesis.GetParameters() {
		if err := parameter.ValidateBasic(); err != nil {
			return fmt.Errorf("invalid parameter in genesis %s : %s", parameter.GetMetaProperty().GetID().AsString(), err.Error())
		}

		validator := parameterManager.GetValidatableParameter(parameter.GetMetaProperty().GetID())
		if validator == nil {
			return fmt.Errorf("invalid parameter in genesis %s : not found", parameter.GetMetaProperty().GetID().AsString())
		}

		if err := validator.GetValidator()(parameter.GetMetaProperty().GetData().Get().AsString()); err != nil {
			return fmt.Errorf("invalid parameter in genesis %s : %s", parameter.GetMetaProperty().GetID().AsString(), err.Error())
		}
	}

	return nil
}
func (parameterManager parameterManager) Fetch(context context.Context) helpers.ParameterManager {
	for _, validatableParameter := range parameterManager.validatableParameters {
		var value string
		parameterManager.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), &value)
		if data, err := validatableParameter.GetParameter().GetMetaProperty().GetData().Get().FromString(value); err != nil {
			panic(err)
		} else {
			validatableParameter = validatableParameter.Mutate(data)
		}
	}

	return parameterManager
}
func (parameterManager parameterManager) Set(context context.Context, parameters []parameters.Parameter) helpers.ParameterManager {
	for _, parameter := range parameters {
		parameterManager.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData().Get().AsString())
	}

	return parameterManager
}
func (parameterManager parameterManager) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameterManager.validatableParameters))

	for i, validatableParameter := range parameterManager.validatableParameters {
		paramSetPairList[i] = paramsTypes.NewParamSetPair(validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData().Get().AsString(), validatableParameter.GetValidator())
	}

	return paramSetPairList
}
func (parameterManager parameterManager) GetKeyTable() paramsTypes.KeyTable {
	return paramsTypes.NewKeyTable().RegisterParamSet(parameterManager)
}
func (parameterManager parameterManager) Initialize(subspace paramsTypes.Subspace) helpers.ParameterManager {
	parameterManager.paramsSubspace = subspace
	return parameterManager
}

func NewParameterManager(validatableParameters ...helpers.ValidatableParameter) helpers.ParameterManager {
	return parameterManager{
		validatableParameters: validatableParameters,
	}
}
