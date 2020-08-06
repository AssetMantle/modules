package make

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type transactionRequest struct {
	BaseReq       rest.BaseReq `json:"baseReq"`
	FromID        string       `json:"fromID"`
	ToID          string       `json:"toID"`
	MaintainersID string       `json:"maintainersID"`
	MakerSplit    int64        `json:"makerSplit"`
	MakerSplitID  string       `json:"makerSplitID"`
	ExchangeRate  string       `json:"exchangeRate"`
	TakerSplitID  string       `json:"takerSplitID"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadInt64(constants.MakerSplit),
		cliCommand.ReadString(constants.MakerSplitID),
		cliCommand.ReadString(constants.ExchangeRate),
		cliCommand.ReadString(constants.TakerSplitID),
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

	exchangeRate, Error := sdkTypes.NewDecFromStr(transactionRequest.ExchangeRate)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.MaintainersID),
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ToID),
		sdkTypes.NewDec(transactionRequest.MakerSplit),
		base.NewID(transactionRequest.MakerSplitID),
		exchangeRate,
		base.NewID(transactionRequest.TakerSplitID),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, maintainersID string, fromID string, toID string,
	makerSplit int64, makerSplitID string, exchangeRate string, takerSplitID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:       baseReq,
		MaintainersID: maintainersID,
		FromID:        fromID,
		ToID:          toID,
		MakerSplit:    makerSplit,
		MakerSplitID:  makerSplitID,
		ExchangeRate:  exchangeRate,
		TakerSplitID:  takerSplitID,
	}
}
