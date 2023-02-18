package docs

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	propertiesUtilities "github.com/AssetMantle/modules/schema/properties/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, request{})
}

func assetIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		classificationID, immutables, _ := ReadAndProcess(context, false, responseWriter, httpRequest)

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewAssetID(classificationID, immutables).AsString(), "", nil))
	}
}

func orderIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		req, classificationID, ImmutableMetaProperties, ImmutableProperties, _, _ := Read(context, responseWriter, httpRequest)
		makerOwnableSplit, _ := sdkTypes.NewDecFromStr(req.MakerOwnableSplit)

		takerOwnableSplit, _ := sdkTypes.NewDecFromStr(req.TakerOwnableSplit)

		fromID, _ := baseIDs.ReadIdentityID(req.FromID)

		takerID, _ := baseIDs.ReadIdentityID(req.TakerID)

		makerOwnableID, _ := baseIDs.ReadOwnableID(req.MakerOwnableID)
		height, _ := strconv.Atoi(req.Height)
		takerOwnableID, _ := baseIDs.ReadOwnableID(req.TakerOwnableID)

		immutableMetaProperties := ImmutableMetaProperties.
			Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(takerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(makerOwnableSplit)))).
			Add(baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(int64(height))))).
			Add(baseProperties.NewMetaProperty(constants.MakerOwnableIDProperty.GetKey(), baseData.NewIDData(makerOwnableID))).
			Add(baseProperties.NewMetaProperty(constants.TakerOwnableIDProperty.GetKey(), baseData.NewIDData(takerOwnableID))).
			Add(baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(fromID))).
			Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(takerID))).
			Add(constants.BondAmountProperty)
		Immutables := base.NewImmutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(immutableMetaProperties.GetList(), ImmutableProperties.GetList()...)...)...))

		//Mutables := base.NewMutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(mutableMetaProperties.GetList(), mutableProperties.GetList()...)...)...))

		//Immutables := base.NewImmutables(immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewDecData(GetTotalWeight(immutables, Mutables).Mul(sdkTypes.NewDec(1))))))
		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewOrderID(classificationID, Immutables).AsString(), "", nil))
	}
}

func GetTotalWeight(immutables qualified.Immutables, mutables qualified.Mutables) int64 {
	totalWeight := int64(0)
	for _, property := range append(immutables.GetImmutablePropertyList().GetList(), mutables.GetMutablePropertyList().GetList()...) {
		totalWeight += property.Get().GetBondWeight()
	}
	return totalWeight
}

func identityIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		classificationID, immutables, _ := ReadAndProcess(context, false, responseWriter, httpRequest)

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewIdentityID(classificationID, immutables).AsString(), "", nil))
	}
}

func identityClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		id, immutables, _ := ReadAndProcess(context, true, responseWriter, httpRequest)
		rest.PostProcessResponse(responseWriter, context, newResponse(id.AsString(), immutables.GetProperty(constants.BondAmountProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).AsString(), nil))
	}
}

func assetClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		id, immutables, _ := ReadAndProcess(context, false, responseWriter, httpRequest)
		rest.PostProcessResponse(responseWriter, context, newResponse(id.AsString(), immutables.GetProperty(constants.BondAmountProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).AsString(), nil))
	}
}

func orderClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		_, _, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := Read(context, responseWriter, httpRequest)
		immutables := base.NewImmutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(append(immutableMetaProperties.GetList(), immutableProperties.GetList()...), constants.ExchangeRateProperty.ToAnyProperty(), constants.CreationHeightProperty.ToAnyProperty(), constants.MakerOwnableIDProperty.ToAnyProperty(), constants.TakerOwnableIDProperty.ToAnyProperty(), constants.MakerIDProperty.ToAnyProperty(), constants.TakerIDProperty.ToAnyProperty())...)...))
		mutables := base.NewMutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(append(mutableMetaProperties.GetList(), mutableProperties.GetList()...), constants.ExpiryHeightProperty.ToAnyProperty(), constants.MakerOwnableSplitProperty.ToAnyProperty())...)...))
		Immutables := base.NewImmutables(immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(GetTotalWeight(immutables, mutables)*baseData.NewNumberData(1).Get()))))
		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewClassificationID(Immutables, mutables).AsString(), Immutables.GetProperty(constants.BondAmountProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).AsString(), nil))
	}
}

func ReadAndProcess(context client.Context, addAuth bool, responseWriter http.ResponseWriter, httpRequest *http.Request) (ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	_, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := Read(context, responseWriter, httpRequest)
	Immutables, Mutables := Process(immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties, addAuth)
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

	immutableMetaProperties, _ := utilities.ReadMetaPropertyList(req.ImmutableMetaProperties)

	immutableProperties, _ := utilities.ReadMetaPropertyList(req.ImmutableProperties)

	immutableProperties = immutableProperties.ScrubData()

	mutableMetaProperties, _ := utilities.ReadMetaPropertyList(req.MutableMetaProperties)

	mutableProperties, _ := utilities.ReadMetaPropertyList(req.MutableProperties)

	mutableProperties = mutableProperties.ScrubData()

	classificationID, _ := baseIDs.ReadClassificationID(req.ClassificationID)
	return req, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties
}

func Process(immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties lists.PropertyList, addAuth bool) (qualified.Immutables, qualified.Mutables) {
	immutables := base.NewImmutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(immutableMetaProperties.GetList(), immutableProperties.GetList()...)...)...))
	var Mutables qualified.Mutables
	if addAuth {
		Mutables = base.NewMutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(append(mutableMetaProperties.GetList(), mutableProperties.GetList()...), constants.AuthenticationProperty.ToAnyProperty())...)...))
	} else {
		Mutables = base.NewMutables(baseLists.NewPropertyList(propertiesUtilities.AnyPropertyListToPropertyList(append(mutableMetaProperties.GetList(), mutableProperties.GetList()...)...)...))
	}

	Immutables := base.NewImmutables(immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(GetTotalWeight(immutables, Mutables)*baseData.NewNumberData(1).Get()))))

	return Immutables, Mutables
}

func RegisterRESTRoutes(context client.Context, router *mux.Router) {
	router.HandleFunc("/get/classification/identity", identityClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/asset", assetClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/classification/order", orderClassificationHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/assetID", assetIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/identityID", identityIDHandler(context)).Methods("POST")
	router.HandleFunc("/get/document/orderID", orderIDHandler(context)).Methods("POST")
}
