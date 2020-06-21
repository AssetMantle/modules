package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

type Request interface {
	ReadFromCLI(CLICommand, context.CLIContext) Request
	GetBaseReq() rest.BaseReq
	MakeMsg() sdkTypes.Msg
}
