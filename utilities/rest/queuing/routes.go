package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/response/{TicketID}", queryDB(cliContext.Codec)).Methods("GET")
}
