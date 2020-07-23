package make

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
	"strings"
)

type transactionRequest struct {
	BaseReq             rest.BaseReq `json:"baseReq"`
	ClassificationID    string       `json:"classificationID"`
	MaintainersID       string       `json:"maintainersID"`
	Properties          string       `json:"properties"`
	Lock                int64        `json:"lock"`
	Burn                int64        `json:"burn"`
	TakerAddress        string       `json:"takerAddress"`
	SenderAddress       string       `json:"senderAddress"`
	FeeRecipientAddress string       `json:"feeRecipientAddress"`
	MakerAssetAmount    int64        `json:"makerAssetAmount"`
	MakerAssetData      string       `json:"makerAssetData"`
	MakerFee            int64        `json:"makerFee"`
	MakerFeeAssetData   string       `json:"makerFeeAssetData"`
	TakerAssetAmount    int64        `json:"takerAssetAmount"`
	TakerAssetData      string       `json:"takerAssetData"`
	TakerFee            int64        `json:"takerFee"`
	TakerFeeAssetData   string       `json:"takerFeeAssetData"`
	ExpirationTime      int64        `json:"expirationTime"`
	Salt                int64        `json:"salt"`
}

var _ utilities.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utilities.CLICommand, cliContext context.CLIContext) utilities.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Properties),
		cliCommand.ReadInt64(constants.Lock),
		cliCommand.ReadInt64(constants.Burn),
		cliCommand.ReadString(constants.TakerAddress),
		cliCommand.ReadInt64(constants.MakerAssetAmount),
		cliCommand.ReadString(constants.MakerAssetData),
		cliCommand.ReadInt64(constants.TakerAssetAmount),
		cliCommand.ReadString(constants.TakerAssetData),
		cliCommand.ReadInt64(constants.Salt),
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

	takerAddress, Error := sdkTypes.AccAddressFromBech32(transactionRequest.TakerAddress)
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
		base.NewID(transactionRequest.MaintainersID),
		base.NewID(transactionRequest.ClassificationID),
		base.NewProperties(propertyList),
		base.NewHeight(transactionRequest.Lock),
		base.NewHeight(transactionRequest.Burn),
		takerAddress,
		sdkTypes.NewDec(transactionRequest.MakerAssetAmount),
		base.NewID(transactionRequest.MakerAssetData),
		sdkTypes.NewDec(transactionRequest.TakerAssetAmount),
		base.NewID(transactionRequest.TakerAssetData),
		base.NewHeight(transactionRequest.Salt),
	)
}

func requestPrototype() utilities.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, classificationID string, maintainersID string, properties string, lock int64, burn int64,
	takerAddress string, makerAssetAmount int64, makerAssetData string, takerAssetAmount int64, takerAssetData string, salt int64,
) utilities.TransactionRequest {
	return transactionRequest{
		BaseReq:          baseReq,
		ClassificationID: classificationID,
		MaintainersID:    maintainersID,
		Properties:       properties,
		Lock:             lock,
		Burn:             burn,
		TakerAddress:     takerAddress,
		MakerAssetAmount: makerAssetAmount,
		MakerAssetData:   makerAssetData,
		TakerAssetAmount: takerAssetAmount,
		TakerAssetData:   takerAssetData,
		Salt:             salt,
	}
}
