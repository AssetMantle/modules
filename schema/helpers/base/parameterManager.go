// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"golang.org/x/net/context"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
)

type parameterManager struct {
	moduleName            string
	validatableParameters []helpers.ValidatableParameter
	paramsSubspace        paramsTypes.Subspace
}

var _ helpers.ParameterManager = (*parameterManager)(nil)

func (parameterManager parameterManager) Get() []helpers.Parameter {
	parameters := make([]helpers.Parameter, len(parameterManager.validatableParameters))
	for i, validatableParameter := range parameterManager.validatableParameters {
		parameters[i] = validatableParameter.GetParameter()
	}
	return parameters
}
func (parameterManager parameterManager) GetParameter(propertyID ids.PropertyID) helpers.Parameter {
	for _, validatableParameter := range parameterManager.validatableParameters {
		if validatableParameter.GetParameter().GetMetaProperty().GetID().Compare(propertyID) == 0 {
			return validatableParameter.GetParameter()
		}
	}
	return nil
}
func (parameterManager parameterManager) Fetch(context context.Context) helpers.ParameterManager {
	for _, validatableParameter := range parameterManager.validatableParameters {
		parameterManager.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData().Get())
	}

	return parameterManager
}
func (parameterManager parameterManager) Set(context context.Context, parameters ...helpers.Parameter) {
	for _, parameter := range parameters {
		parameterManager.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData().Get())
	}
}
func (parameterManager parameterManager) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameterManager.validatableParameters))

	for i, validatableParameter := range parameterManager.validatableParameters {
		paramSetPairList[i] = paramsTypes.NewParamSetPair(validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData().Get(), validatableParameter.GetValidator())
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
