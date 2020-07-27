package burn

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
	BaseReq rest.BaseReq `json:"baseReq"`
	AssetID string       `json:"assetID" valid:"required~required field assetID missing matches(^[a-z]$)~invalid field assetID "`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
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
		base.NewID(transactionRequest.AssetID),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, assetID string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq: baseReq,
		AssetID: assetID,
	}
}
