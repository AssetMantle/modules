// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package helpers

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type TransactionRequest interface {
	FromCLI(CLICommand, client.Context) (TransactionRequest, error)
	FromJSON(json.RawMessage) (TransactionRequest, error)
	MakeMsg() (sdkTypes.Msg, error)
	RegisterCodec(codec *codec.LegacyAmino)
	Request
}
