package assetFactory

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc(fmt.Sprintf("/%v/%v", mint.Transaction.GetModuleName(), constants.MintTransaction), mint.Transaction.RESTRequestHandler(cliContext)).Methods("POST")

	router.HandleFunc(fmt.Sprintf("/%v/%v/{%v}", asset.Query.GetModuleName(), asset.Query.GetName(), constants.AssetID.GetName()), asset.Query.RESTQueryHandler(cliContext)).Methods("GET")
}
