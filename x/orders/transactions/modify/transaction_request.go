// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package modify

import (
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/client"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"net/http"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	helpers.CommonTransactionRequest `json:"commonTransactionRequest"`
	FromID                           string `json:"fromID"`
	OrderID                          string `json:"orderID"`
	TakerSplit                       string `json:"takerSplit"`
	MakerSplit                       string `json:"makerSplit"`
	ExpiresIn                        int64  `json:"expiresIn"`
	MutableMetaProperties            string `json:"mutableMetaProperties"`
	MutableProperties                string `json:"mutableProperties"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Modify order transaction
// @Description Modify order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body transactionRequest true "Request body to modify order transaction"
// @Success 200 {object} transactionResponse "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/modify [post]
func (transactionRequest transactionRequest) Validate() error {
	return helpers.ValidateTransactionRequest(transactionRequest)
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.FromIdentityID),
		cliCommand.ReadString(constants.OrderID),
		cliCommand.ReadString(constants.TakerSplit),
		cliCommand.ReadString(constants.MakerSplit),
		cliCommand.ReadInt64(constants.ExpiresIn),
		cliCommand.ReadString(constants.MutableMetaProperties),
		cliCommand.ReadString(constants.MutableProperties),
	), nil
}
func (transactionRequest transactionRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.TransactionRequest, error) {
	return helpers.TransactionRequestFromHTTPRequest(httpRequest, &transactionRequest)
}
func (transactionRequest transactionRequest) GetCommonTransactionRequest() helpers.CommonTransactionRequest {
	return transactionRequest.CommonTransactionRequest
}

func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetCommonTransactionRequest().GetFrom())
	if err != nil {
		return nil, err
	}

	makerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.MakerSplit)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("maker split %s is not a valid integer", transactionRequest.MakerSplit)
	}

	takerSplit, ok := sdkTypes.NewIntFromString(transactionRequest.TakerSplit)
	if !ok {
		return nil, constants.IncorrectFormat.Wrapf("taker split %s is not a valid integer", transactionRequest.TakerSplit)
	}

	mutableMetaProperties, err := base.NewPropertyList().FromMetaPropertiesString(transactionRequest.MutableMetaProperties)
	if err != nil {
		return nil, err
	}

	mutableProperties, err := base.NewPropertyList().FromMetaPropertiesString(transactionRequest.MutableProperties)
	if err != nil {
		return nil, err
	}
	mutableProperties = mutableProperties.ScrubData()

	fromID, err := baseIDs.PrototypeIdentityID().FromString(transactionRequest.FromID)
	if err != nil {
		return nil, err
	}

	orderID, err := baseIDs.PrototypeOrderID().FromString(transactionRequest.OrderID)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		fromID.(ids.IdentityID),
		orderID.(ids.OrderID),
		takerSplit,
		makerSplit,
		baseTypes.NewHeight(transactionRequest.ExpiresIn),
		mutableMetaProperties,
		mutableProperties,
	), nil
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, fromID string, orderID string, takerSplit string, makerSplit string, expiresIn int64, mutableMetaProperties string, mutableProperties string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		FromID:                   fromID,
		OrderID:                  orderID,
		TakerSplit:               takerSplit,
		MakerSplit:               makerSplit,
		ExpiresIn:                expiresIn,
		MutableMetaProperties:    mutableMetaProperties,
		MutableProperties:        mutableProperties,
	}
}
