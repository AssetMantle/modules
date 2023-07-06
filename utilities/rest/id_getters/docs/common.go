package docs

import (
	"net/http"

	codecUtilities "github.com/AssetMantle/schema/go/codec/utilities"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	"github.com/AssetMantle/schema/go/properties/constants"
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

func ReadAndProcess(context client.Context, addAuth bool, addBond bool, responseWriter http.ResponseWriter, httpRequest *http.Request) (ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	_, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := Read(context, responseWriter, httpRequest)
	Immutables, Mutables := Process(immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, addAuth, addBond)
	if len(classificationID.Bytes()) != 0 {
		return classificationID, Immutables, Mutables
	}
	return baseIDs.NewClassificationID(Immutables, Mutables), Immutables, Mutables
}

func Read(context client.Context, responseWriter http.ResponseWriter, httpRequest *http.Request) (request, ids.ClassificationID, lists.PropertyList, lists.PropertyList, lists.PropertyList, lists.PropertyList) {
	transactionRequest := Prototype()
	if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
		return request{}, nil, nil, nil, nil, nil
	}

	if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
		return request{}, nil, nil, nil, nil, nil
	}

	req := transactionRequest.(request)

	immutableMetaProperties, _ := baseLists.PrototypePropertyList().FromMetaPropertiesString(req.ImmutableMetaProperties)

	immutableProperties, _ := baseLists.PrototypePropertyList().FromMetaPropertiesString(req.ImmutableProperties)

	immutableProperties = immutableProperties.ScrubData()

	mutableMetaProperties, _ := baseLists.PrototypePropertyList().FromMetaPropertiesString(req.MutableMetaProperties)

	mutableProperties, _ := baseLists.PrototypePropertyList().FromMetaPropertiesString(req.MutableProperties)

	mutableProperties = mutableProperties.ScrubData()

	classificationID, _ := baseIDs.PrototypeClassificationID().FromString(req.ClassificationID)
	return req, classificationID.(ids.ClassificationID), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties
}

func Process(immutableMetaPropertyList, immutablePropertyList, mutableMetaPropertyList, mutablePropertyList lists.PropertyList, addAuth bool, addBond bool) (qualified.Immutables, qualified.Mutables) {
	immutables := base.NewImmutables(immutableMetaPropertyList.Add(baseLists.AnyPropertiesToProperties(immutablePropertyList.Get()...)...))
	var Mutables qualified.Mutables
	if addAuth {
		Mutables = base.NewMutables(mutableMetaPropertyList.Add(baseLists.AnyPropertiesToProperties(mutablePropertyList.Add(constants.AuthenticationProperty.ToAnyProperty()).Get()...)...))
	} else {
		Mutables = base.NewMutables(mutableMetaPropertyList.Add(baseLists.AnyPropertiesToProperties(mutablePropertyList.Get()...)...))
	}
	var Immutables qualified.Immutables
	if addBond {
		Immutables = base.NewImmutables(immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(GetTotalWeight(immutables, Mutables).Mul(baseData.NewNumberData(sdkTypes.OneInt()).Get())))))
	} else {
		Immutables = immutables
	}
	return Immutables, Mutables
}

func RegisterRESTRoutes(context client.Context, router *mux.Router) {
	router.HandleFunc("/get/classification/identity", identityClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/asset", assetClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/order", orderClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/assetID", assetIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/identityID", identityIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/orderID", orderIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/nubID", nubIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/splitID", splitIDHandler(context)).Methods("POST")
}
