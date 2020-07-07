package assets

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/modules/assets/constants"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc(fmt.Sprintf("/%v/%v/{%v}", asset.Query.GetModuleName(), asset.Query.GetName(), constants.AssetID.GetName()), asset.Query.RESTQueryHandler(cliContext)).Methods("GET")

	router.HandleFunc(fmt.Sprintf("/%v/%v", burn.Transaction.GetModuleName(), burn.Transaction.GetName()), burn.Transaction.RESTRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("/%v/%v", mint.Transaction.GetModuleName(), mint.Transaction.GetName()), mint.Transaction.RESTRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(fmt.Sprintf("/%v/%v", mutate.Transaction.GetModuleName(), mutate.Transaction.GetName()), mutate.Transaction.RESTRequestHandler(cliContext)).Methods("POST")

}
