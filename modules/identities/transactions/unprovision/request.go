package unprovision

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

type transactionRequest struct {
	BaseReq    rest.BaseReq `json:"baseReq"`
	To         string       `json:"to" valid:"required~required field toaddress missing matches(^commit[a-z0-9]{39}$)~invalid field toaddress"`
	IdentityID string       `json:"identityid" valid:"required~required field identityid missing matches(^[A-Za-z]$)~invalid field identityid"`
}

var _ utilities.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utilities.CLICommand, cliContext context.CLIContext) utilities.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.To),
		cliCommand.ReadString(constants.IdentityID),
	)
}

func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}

func (transactionRequest transactionRequest) MakeMsg() sdkTypes.Msg {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	to, Error := sdkTypes.AccAddressFromBech32(transactionRequest.To)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		to,
		base.NewID(transactionRequest.IdentityID),
	)
}

func requestPrototype() utilities.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, to string, identityID string) utilities.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		To:         to,
		IdentityID: identityID,
	}
}
