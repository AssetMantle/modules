package docs

import (
	"encoding/json"
	"github.com/AssetMantle/modules/utilities/rest"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"io"
	"net/http"
)

func splitIDHandler(context client.Context) http.HandlerFunc {
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

		fromID, _ := baseIDs.PrototypeIdentityID().FromString(request.FromID)

		coins, _ := sdkTypes.ParseCoinsNormalized(request.Coins)

		var coinID ids.AssetID
		for _, coin := range coins {
			coinID = baseDocuments.NewCoinAsset(coin.Denom).GetCoinAssetID()
		}

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewSplitID(coinID, fromID.(ids.IdentityID)).AsString(), nil))
	}
}
