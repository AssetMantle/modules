package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"net/http"
)

func RESTQueryHandler(cliContext context.CLIContext) http.HandlerFunc {
	makeQueryBytes := func(vars map[string]string) []byte {
		return packageCodec.MustMarshalJSON(request{ID: types.NewID(vars["id"])})
	}
	return types.NewRESTQuery(constants.QuerierRoute, constants.AssetQuery).CreateQuery(cliContext, makeQueryBytes)
}
