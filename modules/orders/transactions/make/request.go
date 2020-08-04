package make

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type transactionRequest struct {
	BaseReq          rest.BaseReq `json:"baseReq"`
	Properties       string       `json:"properties"`
	Lock             int64        `json:"lock"`
	Burn             int64        `json:"burn"`
	FromID           string       `json:"fromID"`
	ToID             string       `json:"toID"`
	MakerAssetAmount int64        `json:"makerAssetAmount"`
	MakerAssetData   string       `json:"makerAssetData"`
	TakerAssetAmount int64        `json:"takerAssetAmount"`
	TakerAssetData   string       `json:"takerAssetData"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.Properties),
		cliCommand.ReadInt64(constants.Lock),
		cliCommand.ReadInt64(constants.Burn),
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadInt64(constants.MakerAssetAmount),
		cliCommand.ReadString(constants.MakerAssetData),
		cliCommand.ReadInt64(constants.TakerAssetAmount),
		cliCommand.ReadString(constants.TakerAssetData),
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
			propertyList = append(propertyList, base.NewProperty(base.NewID(traitIDAndProperty[0]), base.NewFact(traitIDAndProperty[1], base.NewSignatures(nil))))
		}
	}

	return newMessage(
		from,
		base.NewProperties(propertyList),
		base.NewHeight(transactionRequest.Lock),
		base.NewHeight(transactionRequest.Burn),
		base.NewID(transactionRequest.FromID),
		base.NewID(transactionRequest.ToID),
		sdkTypes.NewDec(transactionRequest.MakerAssetAmount),
		base.NewID(transactionRequest.MakerAssetData),
		sdkTypes.NewDec(transactionRequest.TakerAssetAmount),
		base.NewID(transactionRequest.TakerAssetData),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, properties string, lock int64, burn int64, fromID string, toID string,
	makerAssetAmount int64, makerAssetData string, takerAssetAmount int64, takerAssetData string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
		FromID:           fromID,
		ToID:             toID,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
	}
}
