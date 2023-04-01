package docs

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

func nubIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := Prototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			panic(errorConstants.IncorrectFormat)
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			panic(errorConstants.IncorrectFormat)
		}

		req := transactionRequest.(request)

		nubID := baseIDs.NewStringID(req.NubID)
		immutables := base.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constants.NubIDProperty.GetKey(), baseData.NewIDData(nubID))))

		// TODO ***** add nub classificationID to genesis
		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewIdentityID(constansts.NubClassificationID, immutables).AsString(), "", nil))
	}
}

func identityIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		classificationID, immutables, _ := ReadAndProcess(context, false, false, responseWriter, httpRequest)

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewIdentityID(classificationID, immutables).AsString(), "", nil))
	}
}

func identityClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		id, immutables, _ := ReadAndProcess(context, true, true, responseWriter, httpRequest)
		rest.PostProcessResponse(responseWriter, context, newResponse(id.AsString(), immutables.GetProperty(constants.BondAmountProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).AsString(), nil))
	}
}
