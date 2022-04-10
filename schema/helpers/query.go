// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

type Query interface {
	GetName() string
	Command(*codec.Codec) *cobra.Command
	HandleMessage(sdkTypes.Context, abciTypes.RequestQuery) ([]byte, error)
	RESTQueryHandler(context.CLIContext) http.HandlerFunc
	Initialize(Mapper, Parameters, ...interface{}) Query
}
