// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	"github.com/AssetMantle/modules/schema/applications/base"
	"github.com/AssetMantle/modules/swagger/configurations"

	_ "github.com/AssetMantle/modules/swagger/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var Prototype = base.NewApplication(
	configurations.Name,
	configurations.ModuleBasicManager,
	configurations.EnabledWasmProposalTypeList,
	configurations.ModuleAccountPermissions,
	configurations.TokenReceiveAllowedModules,
)

// @title AssetMantle Modules Swagger Documentation
// @version 0.1.0
// @description API Documentation of AssetMantle custom modules
// @host localhost:1317

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := mux.NewRouter()
	ctx := context.NewCLIContext()
	Prototype.GetModuleBasicManager().RegisterRESTRoutes(ctx, r)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("listen on :1318")
	log.Fatal(http.ListenAndServe(":1318", r))
}
