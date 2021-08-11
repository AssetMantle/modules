/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

type Transaction interface {
	GetName() string
	Command() *cobra.Command
	HandleMessage(sdkTypes.Context, sdkTypes.Msg) (*sdkTypes.Result, error)
	RESTRequestHandler(client.Context) http.HandlerFunc
	RegisterCodec(*codec.LegacyAmino)
	DecodeTransactionRequest(json.RawMessage) (sdkTypes.Msg, error)
	InitializeKeeper(Mapper, Parameters, ...interface{}) Transaction
}
