/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"net/http"
)

type Transaction interface {
	GetName() string
	Command(*codec.Codec) *cobra.Command
	HandleMessage(sdkTypes.Context, sdkTypes.Msg) (*sdkTypes.Result, error)
	RESTRequestHandler(context.CLIContext) http.HandlerFunc
	RegisterCodec(*codec.Codec)
	DecodeTransactionRequest(json.RawMessage) (sdkTypes.Msg, error)
	InitializeKeeper(Mapper, Parameters, ...interface{}) Transaction
}
