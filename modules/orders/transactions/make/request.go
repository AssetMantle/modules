package make

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/types/schema"
	"github.com/persistenceOne/persistenceSDK/types/utility"
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

var _ utility.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand utility.CLICommand, cliContext context.CLIContext) utility.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.ClassificationID),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Properties),
		cliCommand.ReadInt64(constants.Lock),
		cliCommand.ReadInt64(constants.Burn),
		cliCommand.ReadString(constants.TakerAddress),
		cliCommand.ReadString(constants.SenderAddress),
		cliCommand.ReadString(constants.FeeRecipientAddress),
		cliCommand.ReadInt64(constants.MakerAssetAmount),
		cliCommand.ReadString(constants.MakerAssetData),
		cliCommand.ReadInt64(constants.MakerFee),
		cliCommand.ReadString(constants.MakerFeeAssetData),
		cliCommand.ReadInt64(constants.TakerAssetAmount),
		cliCommand.ReadString(constants.TakerAssetData),
		cliCommand.ReadInt64(constants.TakerFee),
		cliCommand.ReadString(constants.TakerFeeAssetData),
		cliCommand.ReadInt64(constants.ExpirationTime),
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

	senderAddress, Error := sdkTypes.AccAddressFromBech32(transactionRequest.SenderAddress)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	feeRecipientAddress, Error := sdkTypes.AccAddressFromBech32(transactionRequest.FeeRecipientAddress)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	properties := strings.Split(transactionRequest.Properties, constants.PropertiesSeparator)
	if len(properties) > constants.MaxTraitCount {
		panic(errors.New(fmt.Sprintf("")))
	}

	var propertyList []schema.Property
	for _, property := range properties {
		traitIDAndProperty := strings.Split(property, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			propertyList = append(propertyList, schema.NewProperty(schema.NewID(traitIDAndProperty[0]), schema.NewFact(traitIDAndProperty[1], schema.NewSignatures(nil))))
		}
	}

	return newMessage(
		from,
		schema.NewID(transactionRequest.MaintainersID),
		schema.NewID(transactionRequest.ClassificationID),
		schema.NewProperties(propertyList),
		schema.NewHeight(transactionRequest.Lock),
		schema.NewHeight(transactionRequest.Burn),
		takerAddress,
		senderAddress,
		feeRecipientAddress,
		sdkTypes.NewDec(transactionRequest.MakerAssetAmount),
		schema.NewID(transactionRequest.MakerAssetData),
		sdkTypes.NewDec(transactionRequest.MakerFee),
		schema.NewID(transactionRequest.MakerFeeAssetData),
		sdkTypes.NewDec(transactionRequest.TakerAssetAmount),
		schema.NewID(transactionRequest.TakerAssetData),
		sdkTypes.NewDec(transactionRequest.TakerFee),
		schema.NewID(transactionRequest.TakerFeeAssetData),
		schema.NewHeight(transactionRequest.ExpirationTime),
		schema.NewHeight(transactionRequest.Salt),
	)
}

func requestPrototype() utility.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, classificationID string, maintainersID string, properties string, lock int64, burn int64,
	takerAddress string,
	senderAddress string,
	feeRecipientAddress string,
	makerAssetAmount int64,
	makerAssetData string,
	makerFee int64,
	makerFeeAssetData string,
	takerAssetAmount int64,
	takerAssetData string,
	takerFee int64,
	takerFeeAssetData string,
	expirationTime int64,
	salt int64,
) utility.TransactionRequest {
	return transactionRequest{
		BaseReq:             baseReq,
		ClassificationID:    classificationID,
		MaintainersID:       maintainersID,
		Properties:          properties,
		Lock:                lock,
		Burn:                burn,
		TakerAddress:        takerAddress,
		SenderAddress:       senderAddress,
		FeeRecipientAddress: feeRecipientAddress,
		MakerAssetAmount:    makerAssetAmount,
		MakerAssetData:      makerAssetData,
		MakerFee:            makerFee,
		MakerFeeAssetData:   makerFeeAssetData,
		TakerAssetAmount:    takerAssetAmount,
		TakerAssetData:      takerAssetData,
		TakerFee:            takerFee,
		TakerFeeAssetData:   takerFeeAssetData,
		ExpirationTime:      expirationTime,
		Salt:                salt,
	}
}
