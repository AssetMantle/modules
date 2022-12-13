package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/persistenceOne/persistenceSDK/schema/applications/base"
	"github.com/persistenceOne/persistenceSDK/swagger/configurations"

	_ "github.com/persistenceOne/persistenceSDK/swagger/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var Prototype = base.NewApplication(
	configurations.Name,
	configurations.ModuleBasicManager,
	configurations.MakeEncodingConfig(),
	configurations.EnabledWasmProposalTypeList,
	configurations.ModuleAccountPermissions,
)

// @title Persistence Swagger Documentation
// @version 0.1.0
// @description API Documentation of Persistence custom modules
// @host localhost:1317

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := mux.NewRouter()
	ctx := client.Context{}
	Prototype.GetModuleBasicManager().RegisterRESTRoutes(ctx, r)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("listen on :1318")
	log.Fatal(http.ListenAndServe(":1318", r))
}
