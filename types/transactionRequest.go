package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type TransactionRequest interface {
	Request
	ReadFromCLI(CLICommand, context.CLIContext) TransactionRequest
	GetBaseReq() rest.BaseReq
	MakeMsg() sdkTypes.Msg
}
