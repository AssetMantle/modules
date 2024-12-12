// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"cosmossdk.io/math"
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"
	"io"
	"net/http"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/qualified"
	"github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/gorilla/mux"
)

func GetTotalWeight(immutables qualified.Immutables, mutables qualified.Mutables) math.Int {
	totalWeight := sdkTypes.ZeroInt()
	for _, property := range append(immutables.GetImmutablePropertyList().Get(), mutables.GetMutablePropertyList().Get()...) {
		totalWeight = totalWeight.Add(property.Get().GetBondWeight())
	}
	return totalWeight
}

func ReadAndProcess(context client.Context, responseWriter http.ResponseWriter, httpRequest *http.Request) (ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	_, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := read(context, responseWriter, httpRequest)
	Immutables, Mutables := Process(immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties)
	if len(classificationID.Bytes()) != 0 {
		return classificationID, Immutables, Mutables
	}
	return baseIDs.NewClassificationID(Immutables, Mutables), Immutables, Mutables
}

func read(context client.Context, responseWriter http.ResponseWriter, httpRequest *http.Request) (request, ids.ClassificationID, lists.PropertyList, lists.PropertyList, lists.PropertyList, lists.PropertyList) {
	body, err := io.ReadAll(httpRequest.Body)
	if err != nil {
		rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
		return request{}, nil, nil, nil, nil, nil
	}

	request := request{}
	if err := json.Unmarshal(body, &request); err != nil {
		rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
		return request, nil, nil, nil, nil, nil
	}

	immutableMetaProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(request.ImmutableMetaProperties)

	immutableProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(request.ImmutableProperties)

	immutableProperties = immutableProperties.ScrubData()

	mutableMetaProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(request.MutableMetaProperties)

	mutableProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(request.MutableProperties)

	mutableProperties = mutableProperties.ScrubData()

	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(request.ClassificationID)
	return request, classificationID.(ids.ClassificationID), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties
}

func Process(immutableMetaPropertyList, immutablePropertyList, mutableMetaPropertyList, mutablePropertyList lists.PropertyList) (qualified.Immutables, qualified.Mutables) {
	immutables := base.NewImmutables(immutableMetaPropertyList.Add(baseLists.AnyPropertiesToProperties(immutablePropertyList.Get()...)...))
	mutables := base.NewMutables(mutableMetaPropertyList.Add(baseLists.AnyPropertiesToProperties(mutablePropertyList.Get()...)...))

	return immutables, mutables
}

func RegisterRESTRoutes(context client.Context, router *mux.Router) {
	router.HandleFunc("/get/classification/identity", identityClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/asset", assetClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/order", orderClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/assetID", assetIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/identityID", identityIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/orderID", orderIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/nameIdentityID", nameIdentityIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/splitID", splitIDHandler(context)).Methods("POST")
}
