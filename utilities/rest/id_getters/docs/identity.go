package docs

import (
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	"net/http"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
)

func nameIdentityIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := Prototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			panic(fmt.Errorf("failed to read request"))
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			panic(fmt.Errorf("failed to validate request"))
		}

		rest.PostProcessResponse(responseWriter, context, newResponse(baseDocuments.NewNameIdentity(baseIDs.NewStringID(transactionRequest.(request).Name), baseData.PrototypeListData()).GetNameIdentityID().AsString(), nil))
	}
}

func identityIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		classificationID, immutables, _ := ReadAndProcess(context, responseWriter, httpRequest)

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewIdentityID(classificationID, immutables).AsString(), nil))
	}
}

func identityClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		id, _, _ := ReadAndProcess(context, responseWriter, httpRequest)
		rest.PostProcessResponse(responseWriter, context, newResponse(id.AsString(), nil))
	}
}
