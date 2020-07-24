package send

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
	BaseReq   rest.BaseReq `json:"baseReq"`
	ToID      string       `json:"toid" valid:"required~required field toid missing matches(^[A-Za-z]$)~invalid field toid"`
	OwnableID string       `json:"ownableid" valid:"required~required field ownableid missing matches(^[A-Za-z]$)~invalid field ownableid"`
	Split     string       `json:"split" valid:"required~required field split missing matches(^[A-Za-z]$)~invalid field split"`
}

var _ utilities.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utilities.CLICommand, cliContext context.CLIContext) utilities.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.OwnableID),
		cliCommand.ReadString(constants.Split),
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
	split, Error := sdkTypes.NewDecFromStr(transactionRequest.Split)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.ToID),
		base.NewID(transactionRequest.OwnableID),
		split,
	)
}

func requestPrototype() utilities.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, toID string, ownableID string, split string) utilities.TransactionRequest {
	return transactionRequest{
		BaseReq:   baseReq,
		ToID:      toID,
		OwnableID: ownableID,
		Split:     split,
	}
}
