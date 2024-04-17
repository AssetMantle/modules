// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"net/http"

	"github.com/AssetMantle/schema/go/ids"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/parameters"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/helpers"
)

type parameterManager struct {
	moduleName            string
	validatableParameters []helpers.ValidatableParameter
	paramsSubspace        paramsTypes.Subspace
}

var _ helpers.ParameterManager = (*parameterManager)(nil)

func (parameterManager parameterManager) Get() lists.ParameterList {
	Parameters := make([]parameters.Parameter, len(parameterManager.validatableParameters))
	for i, validatableParameter := range parameterManager.validatableParameters {
		Parameters[i] = validatableParameter.GetParameter()
	}
	return baseLists.NewParameterList(Parameters...)
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
func (parameterManager parameterManager) ValidateParameter(parameter parameters.Parameter) error {
	validator := parameterManager.GetValidatableParameter(parameter.GetMetaProperty().GetID())
	if validator != nil {
		if err := parameter.ValidateBasic(); err != nil {
			return err
		}

		return validator.GetValidator()(parameter.GetMetaProperty().GetData().Get().AsString())
	}
	return errorConstants.EntityNotFound.Wrapf("parameter with id %s not found", parameter.GetMetaProperty().GetID().AsString())
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
func (parameterManager parameterManager) Set(context context.Context, parameterList lists.ParameterList) helpers.ParameterManager {
	for _, parameter := range parameterList.Get() {
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
func (parameterManager parameterManager) RESTQueryHandler(clientContext client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		clientContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, clientContext, request)
		if !ok {
			return
		}

		responseBytes, height, err := clientContext.QueryWithData(fmt.Sprintf("custom/%s/parameters", parameterManager.moduleName), nil)
		if rest.CheckInternalServerError(responseWriter, err) {
			return
		}

		clientContext = clientContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, clientContext, responseBytes)
	}
}
func (parameterManager parameterManager) Initialize(subspace paramsTypes.Subspace) helpers.ParameterManager {
	parameterManager.paramsSubspace = subspace
	return parameterManager
}

func NewParameterManager(moduleName string, validatableParameters ...helpers.ValidatableParameter) helpers.ParameterManager {
	return parameterManager{
		moduleName:            moduleName,
		validatableParameters: validatableParameters,
	}
}
