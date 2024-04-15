package docs

import (
	"net/http"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	"github.com/AssetMantle/schema/go/qualified"
	"github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, request{})
}

func GetTotalWeight(immutables qualified.Immutables, mutables qualified.Mutables) sdkTypes.Int {
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
	transactionRequest := Prototype()
	if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
		return request{}, nil, nil, nil, nil, nil
	}

	if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
		return request{}, nil, nil, nil, nil, nil
	}

	req := transactionRequest.(request)

	immutableMetaProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(req.ImmutableMetaProperties)

	immutableProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(req.ImmutableProperties)

	immutableProperties = immutableProperties.ScrubData()

	mutableMetaProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(req.MutableMetaProperties)

	mutableProperties, _ := baseLists.NewPropertyList().FromMetaPropertiesString(req.MutableProperties)

	mutableProperties = mutableProperties.ScrubData()

	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(req.ClassificationID)
	return req, classificationID.(ids.ClassificationID), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties
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
