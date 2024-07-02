package docs

import (
	"github.com/AssetMantle/modules/utilities/rest"
	"net/http"
	"strconv"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

func orderIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		req, classificationID, ImmutableMetaProperties, ImmutableProperties, _, _ := read(context, responseWriter, httpRequest)
		makerSplit, _ := sdkTypes.NewDecFromStr(req.MakerSplit)

		takerSplit, _ := sdkTypes.NewDecFromStr(req.TakerSplit)

		fromID, _ := baseIDs.PrototypeIdentityID().FromString(req.FromID)

		takerID, _ := baseIDs.PrototypeIdentityID().FromString(req.TakerID)

		makerAssetID, _ := baseIDs.PrototypeAssetID().FromString(req.MakerAssetID)
		height, _ := strconv.Atoi(req.Height)
		takerAssetID, _ := baseIDs.PrototypeAssetID().FromString(req.TakerAssetID)

		immutableMetaProperties := ImmutableMetaProperties.
			Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(takerSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(makerSplit)))).
			Add(baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(int64(height))))).
			Add(baseProperties.NewMetaProperty(constants.MakerAssetIDProperty.GetKey(), baseData.NewIDData(makerAssetID))).
			Add(baseProperties.NewMetaProperty(constants.TakerAssetIDProperty.GetKey(), baseData.NewIDData(takerAssetID))).
			Add(baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(fromID))).
			Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(takerID)))

		Immutables := base.NewImmutables(immutableMetaProperties.Add(baseLists.AnyPropertiesToProperties(ImmutableProperties.Get()...)...))

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewOrderID(classificationID, Immutables).AsString(), nil))
	}
}
func orderClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		_, _, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties := read(context, responseWriter, httpRequest)
		immutables := base.NewImmutables(
			immutableMetaProperties.Add(
				baseLists.AnyPropertiesToProperties(
					immutableProperties.Add(
						constants.ExchangeRateProperty.ToAnyProperty(),
						constants.CreationHeightProperty.ToAnyProperty(),
						constants.MakerAssetIDProperty.ToAnyProperty(),
						constants.TakerAssetIDProperty.ToAnyProperty(),
						constants.MakerIDProperty.ToAnyProperty(),
						constants.TakerIDProperty.ToAnyProperty(),
					).Get()...,
				)...,
			),
		)
		mutables := base.NewMutables(
			mutableMetaProperties.Add(
				baseLists.AnyPropertiesToProperties(
					mutableProperties.Add(
						constants.ExpiryHeightProperty.ToAnyProperty(),
						constants.MakerSplitProperty.ToAnyProperty(),
					).Get()...,
				)...,
			),
		)
		Immutables := base.NewImmutables(immutables.GetImmutablePropertyList().Add(baseProperties.NewMetaProperty(constants.BondAmountProperty.GetKey(), baseData.NewNumberData(GetTotalWeight(immutables, mutables).Mul(baseData.NewNumberData(sdkTypes.OneInt()).Get())))))
		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewClassificationID(Immutables, mutables).AsString(), nil))
	}
}
