package queuing

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliContext client.Context, router *mux.Router) {
	router.HandleFunc("/response/{TicketID}", queryDB(cliContext.Codec)).Methods("GET")
}
