package issue

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/modules/identities/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionRequest struct {
	BaseReq    rest.BaseReq `json:"baseReq"`
	IdentityID string       `json:"identityID"`
}

var _ types.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
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
	return newMessage(
		from,
		types.NewID(transactionRequest.IdentityID),
	)
}

func requestPrototype() types.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, identityID string) types.TransactionRequest {
	return transactionRequest{
		BaseReq:    baseReq,
		IdentityID: identityID,
	}
}
