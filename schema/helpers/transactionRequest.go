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
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type TransactionRequest interface {
	GetBaseReq() rest.BaseReq

	FromCLI(CLICommand, context.CLIContext) (TransactionRequest, error)
	FromJSON(json.RawMessage) (TransactionRequest, error)
	MakeMsg() (sdkTypes.Msg, error)
	RegisterCodec(*codec.Codec)
	Request
}
