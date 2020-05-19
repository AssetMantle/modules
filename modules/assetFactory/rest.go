package assetFactory

import (
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/lock"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/send"
	"strings"

	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.MintTransaction}, "/"), burn.RestRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.MintTransaction}, "/"), lock.RestRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.MintTransaction}, "/"), mint.RestRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.SendTransaction}, "/"), send.RestRequestHandler(cliContext)).Methods("POST")

	router.HandleFunc(strings.Join([]string{"", QuerierRoute, constants.AssetQuery, "{address}"}, "/"), asset.RestQueryHandler(cliContext)).Methods("GET")
}
