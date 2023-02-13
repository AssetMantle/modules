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

type parameterList struct {
	moduleName            string
	validatableParameters []helpers.ValidatableParameter
	paramsSubspace        paramsTypes.Subspace
}

var _ helpers.ParameterList = (*parameterList)(nil)

func (parameterList parameterList) Get() []helpers.Parameter {
	parameters := make([]helpers.Parameter, len(parameterList.validatableParameters))
	for i, validatableParameter := range parameterList.validatableParameters {
		parameters[i] = validatableParameter.GetParameter()
	}
	return parameters
}
func (parameterList parameterList) GetParameter(propertyID ids.PropertyID) helpers.Parameter {
	for _, validatableParameter := range parameterList.validatableParameters {
		if validatableParameter.GetParameter().GetMetaProperty().GetID().Compare(propertyID) == 0 {
			return validatableParameter.GetParameter()
		}
	}
	return nil
}
func (parameterList parameterList) Fetch(context context.Context) helpers.ParameterList {
	for _, validatableParameter := range parameterList.validatableParameters {
		parameterList.paramsSubspace.Get(sdkTypes.UnwrapSDKContext(context), validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData().Get())
	}

	return parameterList
}
func (parameterList parameterList) Set(context context.Context, parameters ...helpers.Parameter) {
	for _, parameter := range parameters {
		parameterList.paramsSubspace.Set(sdkTypes.UnwrapSDKContext(context), parameter.GetMetaProperty().GetID().Bytes(), parameter.GetMetaProperty().GetData().Get())
	}
}
func (parameterList parameterList) ParamSetPairs() paramsTypes.ParamSetPairs {
	paramSetPairList := make([]paramsTypes.ParamSetPair, len(parameterList.validatableParameters))

	for i, validatableParameter := range parameterList.validatableParameters {
		paramSetPairList[i] = paramsTypes.NewParamSetPair(validatableParameter.GetParameter().GetMetaProperty().GetID().Bytes(), validatableParameter.GetParameter().GetMetaProperty().GetData().Get(), validatableParameter.GetValidator())
	}

	return paramSetPairList
}
func (parameterList parameterList) GetKeyTable() paramsTypes.KeyTable {
	return paramsTypes.NewKeyTable().RegisterParamSet(parameterList)
}
func (parameterList parameterList) RESTQueryHandler(clientContext client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		clientContext, ok := rest.ParseQueryHeightOrReturnBadRequest(responseWriter, clientContext, request)
		if !ok {
			return
		}

		responseBytes, height, err := clientContext.QueryWithData(fmt.Sprintf("custom/%s/parameters", parameterList.moduleName), nil)
		if rest.CheckInternalServerError(responseWriter, err) {
			return
		}

		clientContext = clientContext.WithHeight(height)
		rest.PostProcessResponse(responseWriter, clientContext, responseBytes)
	}
}
func (parameterList parameterList) Initialize(subspace paramsTypes.Subspace) helpers.ParameterList {
	parameterList.paramsSubspace = subspace
	return parameterList
}

func NewParameterList(moduleName string, validatableParameters ...helpers.ValidatableParameter) helpers.ParameterList {
	return parameterList{
		moduleName:            moduleName,
		validatableParameters: validatableParameters,
	}
}
