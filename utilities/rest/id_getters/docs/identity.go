// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"
	"io"
	"net/http"

	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
)

func nameIdentityIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		body, err := io.ReadAll(httpRequest.Body)
		if err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		request := request{}
		if err := json.Unmarshal(body, &request); err != nil {
			rest.WriteErrorResponse(responseWriter, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(responseWriter, context, newResponse(baseDocuments.NewNameIdentity(baseIDs.NewStringID(request.Name), baseData.PrototypeListData()).GetNameIdentityID().AsString(), nil))
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
