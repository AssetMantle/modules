// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type TransactionRequest struct {
	BaseReq rest.BaseReq
	ID      string
}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

func (transactionRequest TransactionRequest) Validate() error {
	return nil
}
func (transactionRequest TransactionRequest) FromCLI(_ helpers.CLICommand, _ client.Context) (helpers.TransactionRequest, error) {
	return transactionRequest, nil
}
func (transactionRequest TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest TransactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}
func (transactionRequest TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	return NewTestMessage(sdkTypes.AccAddress(transactionRequest.BaseReq.From), transactionRequest.ID), nil
}
func (TransactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, TransactionRequest{})
}
func TestTransactionRequestPrototype() helpers.TransactionRequest {
	return TransactionRequest{}
}

type transactionResponse struct {
}

func (t transactionResponse) GetResult() *sdkTypes.Result {
	// TODO implement me
	panic("implement me")
}

var _ helpers.TransactionResponse = (*transactionResponse)(nil)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (t transactionKeeper) Transact(_ context.Context, _ helpers.Message) (helpers.TransactionResponse, error) {
	return transactionResponse{}, nil
}

func (t transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return transactionKeeper{mapper: mapper}
}

func TestTransactionKeeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
