/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package queuing

import (
	"fmt"
	"github.com/gorilla/mux"
	dbm "github.com/tendermint/tm-db"
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"
)

// setTicketIDtoDB : initiates TicketID in Database
func setTicketIDtoDB(ticket TicketID, kafkaDB *dbm.GoLevelDB, cdc *codec.Codec, msg []byte) {
	ticketID, Error := cdc.MarshalJSON(ticket)
	if Error != nil {
		panic(Error)
	}

	if Error := kafkaDB.Set(ticketID, msg); Error != nil {
		panic(Error)
	}
}

// addResponseToDB : Updates response to DB
func addResponseToDB(ticket TicketID, response []byte, kafkaDB *dbm.GoLevelDB, cdc *codec.Codec) {
	ticketID, err := cdc.MarshalJSON(ticket)
	if err != nil {
		panic(err)
	}

	err = kafkaDB.SetSync(ticketID, response)
	if err != nil {
		panic(err)
	}
}

// getResponseFromDB : gives the response from DB
func getResponseFromDB(ticket TicketID, kafkaDB *dbm.GoLevelDB, cdc *codec.Codec) []byte {
	ticketID, err := cdc.MarshalJSON(ticket)
	if err != nil {
		panic(err)
	}

	val, _ := kafkaDB.Get(ticketID)

	return val
}

// queryDB : REST outputs info from DB
func queryDB(cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		iDByte, err := cdc.MarshalJSON(vars["TicketID"])
		if err != nil {
			panic(err)
		}

		var response []byte

		check, _ := KafkaState.KafkaDB.Has(iDByte)
		if check {
			response = getResponseFromDB(TicketID(vars["TicketID"]), KafkaState.KafkaDB, cdc)
		} else {
			output, err := cdc.MarshalJSON("The ticket ID does not exist, it must have been deleted, Query the chain to know")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte(fmt.Sprintf("ticket ID does not exist. Error: %s", err.Error())))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(output)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write(response)
	}
}
