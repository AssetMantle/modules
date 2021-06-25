package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
)

func RegisterRoutes(cliContext context.CLIContext, router *mux.Router) {
	if flags.Queuing.GetValue().(bool) {
		router.HandleFunc("/response/{TicketID}", queryDB(cliContext.Codec)).Methods("GET")
	}
}
