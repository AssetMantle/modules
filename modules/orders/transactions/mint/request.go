package mint

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/modules/orders/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

type transactionRequest struct {
	BaseReq        rest.BaseReq `json:"baseReq"`
	BuyCoinDenom   string       `json:"buyCoinDenom"`
	SellCoinDenom  string       `json:"sellCoinDenom"`
	Properties     string       `json:"properties"`
	BuyCoinAmount  int64        `json:"buyCoinAmount"`
	SellCoinAmount int64        `json:"sellCoinAmount"`
}

var _ types.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand types.CLICommand, cliContext context.CLIContext) types.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.Properties),
		cliCommand.ReadString(constants.BuyCoinDenom),
		cliCommand.ReadString(constants.SellCoinDenom),
		cliCommand.ReadInt64(constants.BuyCoinAmount),
		cliCommand.ReadInt64(constants.SellCoinAmount),
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
	properties := strings.Split(transactionRequest.Properties, constants.PropertiesSeparator)
	if len(properties) > constants.MaxTraitCount {
		panic(errors.New(fmt.Sprintf("")))
	}

	var propertyList []types.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, types.NewProperty(types.NewID(traitIDAndProperty[0]), types.NewFact(traitIDAndProperty[1], types.NewSignatures(nil))))
		}
	}
	return newMessage(from, types.NewProperties(propertyList), transactionRequest.SellCoinDenom, sdkTypes.NewInt(transactionRequest.SellCoinAmount), transactionRequest.BuyCoinDenom, sdkTypes.NewInt(transactionRequest.BuyCoinAmount))

}

func requestPrototype() types.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, properties string, buyCoinDenom string, sellCoinDenom string, buyCoinAmount int64, sellCoinAmount int64) types.TransactionRequest {
	return transactionRequest{
		BaseReq:        baseReq,
		BuyCoinDenom:   buyCoinDenom,
		SellCoinDenom:  sellCoinDenom,
		Properties:     properties,
		BuyCoinAmount:  buyCoinAmount,
		SellCoinAmount: sellCoinAmount,
	}
}
