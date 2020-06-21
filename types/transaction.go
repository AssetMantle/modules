package types

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"net/http"
)

type Transaction interface {
	GetTransactionCommand(*codec.Codec) *cobra.Command
	HandleMessage(sdkTypes.Context, TransactionKeeper, sdkTypes.Msg) (*sdkTypes.Result, error)
	GetRESTRequestHandler(context.CLIContext) http.HandlerFunc
}

type transaction struct {
	Module             string
	TransactionCommand func(*codec.Codec) *cobra.Command
	RESTRequestHandler func(context.CLIContext) http.HandlerFunc
}

var _ Transaction = (*transaction)(nil)

func (transaction transaction) GetTransactionCommand(codec *codec.Codec) *cobra.Command {
	return transaction.TransactionCommand(codec)
}

func (transaction transaction) HandleMessage(context sdkTypes.Context, transactionKeeper TransactionKeeper, message sdkTypes.Msg) (*sdkTypes.Result, error) {

	if Error := transactionKeeper.Transact(context, message); Error != nil {
		return nil, Error
	}

	context.EventManager().EmitEvent(
		sdkTypes.NewEvent(
			sdkTypes.EventTypeMessage,
			sdkTypes.NewAttribute(sdkTypes.AttributeKeyModule, transaction.Module),
		),
	)

	return &sdkTypes.Result{Events: context.EventManager().ABCIEvents()}, nil
}

func (transaction transaction) GetRESTRequestHandler(cliContext context.CLIContext) http.HandlerFunc {
	return transaction.RESTRequestHandler(cliContext)
}

func NewTransaction(module string, use string, short string, long string, requestPrototype func() Request, flagList []CLIFlag) Transaction {
	return &transaction{
		Module:             module,
		TransactionCommand: NewCLICommand(use, short, long, flagList).TransactionCommand(requestPrototype),
		RESTRequestHandler: NewRESTRequest().RequestHandler(requestPrototype),
	}
}
