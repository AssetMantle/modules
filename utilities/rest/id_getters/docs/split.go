package docs

import (
	"fmt"
	"github.com/AssetMantle/modules/utilities/rest"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"net/http"
)

func splitIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := Prototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			panic(fmt.Errorf("failed to read request"))
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			panic(fmt.Errorf("failed to validate request"))
		}

		req := transactionRequest.(request)

		fromID, _ := baseIDs.PrototypeIdentityID().FromString(req.FromID)

		coins, _ := sdkTypes.ParseCoinsNormalized(req.Coins)

		var coinID ids.AssetID
		for _, coin := range coins {
			coinID = baseDocuments.NewCoinAsset(coin.Denom).GetCoinAssetID()
		}

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewSplitID(coinID, fromID.(ids.IdentityID)).AsString(), nil))
	}
}
