package main

import (
	"log"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	"github.com/persistenceOne/persistenceSDK/schema/applications/base"
	"github.com/persistenceOne/persistenceSDK/swagger/configurations"

	_ "github.com/persistenceOne/persistenceSDK/swagger/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

var Prototype = base.NewApplication(
	configurations.Name,
	configurations.ModuleBasicManager,
	configurations.EnabledWasmProposalTypeList,
	configurations.ModuleAccountPermissions,
	configurations.TokenReceiveAllowedModules,
)

// @title Persistence Swagger Documentation
// @version 0.1.0
// @description API Documentation of Persistence custom modules
// @host localhost:1317
// @BasePath /xrpt

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	//r.Use(mux.CORSMethodMiddleware(r))
	//c := cors.New(cors.Options{
	//	AllowedMethods: []string{"GET","POST", "OPTIONS"},
	//	AllowedOrigins: []string{"*"},
	//	AllowCredentials: true,
	//	AllowedHeaders: []string{"Content-Type","Bearer","Bearer ","content-type","Origin","Accept"},
	//	OptionsPassthrough: true,
	//})
	//http.Handle("/", r)
	ctx := context.NewCLIContext()
	Prototype.GetModuleBasicManager().RegisterRESTRoutes(ctx, r)
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//r.Headers("Access-Control-Allow-Origin", "*")
	log.Println("listen on :1318")
	log.Fatal(http.ListenAndServe(":1318", r))
}

//func fooHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")
//	w.Header().Set("Sec-Fetch-Mode","no-cors")
//	if r.Method == http.MethodOptions {
//		return
//	}
//
//	w.Write([]byte("foo"))
//}
