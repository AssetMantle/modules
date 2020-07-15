package burn

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types"
)

type transactionRequest struct {
	BaseReq rest.BaseReq `json:"baseReq"`
	AssetID string       `json:"assetID"`
}

var _ types.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.AssetID),
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
		types.NewID(transactionRequest.AssetID),
	)
}

func requestPrototype() types.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, assetID string) types.TransactionRequest {
	return transactionRequest{
		BaseReq: baseReq,
		AssetID: assetID,
	}
}
