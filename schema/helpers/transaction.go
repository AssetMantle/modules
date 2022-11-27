// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type Transaction interface {
	GetName() string
	Command() *cobra.Command
	HandleMessage(sdkTypes.Context, Message) (*sdkTypes.Result, error)
	RESTRequestHandler(client.Context) http.HandlerFunc
	Service() (*grpc.ServiceDesc, interface{})
	RegisterCodec(*codec.LegacyAmino)
	DecodeTransactionRequest(json.RawMessage) (sdkTypes.Msg, error)
	InitializeKeeper(Mapper, Parameters, ...interface{}) Transaction
}
