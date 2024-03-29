package docs

import (
	"net/http"

	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func splitIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		transactionRequest := Prototype()
		if !rest.ReadRESTReq(responseWriter, httpRequest, context.LegacyAmino, &transactionRequest) {
			panic(errorConstants.IncorrectFormat)
		}

		if rest.CheckBadRequestError(responseWriter, transactionRequest.Validate()) {
			panic(errorConstants.IncorrectFormat)
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
