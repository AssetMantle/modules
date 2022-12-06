// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/data/utilities"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

func (transactionRequest *TransactionRequest) GetRequest() helpers.TransactionRequest {
	// TODO implement me
	panic("implement me")
}

func (transactionRequest *TransactionRequest) ValidateBasic(responseWriter http.ResponseWriter) bool {
	// TODO implement me
	panic("implement me")
}
func (transactionRequest *TransactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest *TransactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context),
		cliCommand.ReadString(constants.Data),
	), nil
}
func (transactionRequest *TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}

func (transactionRequest *TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from := transactionRequest.From

	data, err := utilities.ReadData(transactionRequest.Data)
	if err != nil {
		return nil, err
	}

	return newMessage(
		[]byte(from),
		data,
	), nil
}
func (*TransactionRequest) RegisterCodec(codec *codec.LegacyAmino) {
	schema.RegisterModuleConcrete(codec, TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return &TransactionRequest{}
}
func newTransactionRequest(baseReq rest.BaseReq, data string) helpers.TransactionRequest {
	return &TransactionRequest{
		From: baseReq.From,
		Data: data,
	}
}
