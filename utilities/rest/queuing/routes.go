// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc("/response/{TicketID}", queryDB(cliContext.Codec)).Methods("GET")
}
