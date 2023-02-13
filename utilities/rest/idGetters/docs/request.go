// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/helpers"
)

type request struct {
	BaseReq                 rest.BaseReq `json:"baseReq"`
	FromID                  string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"required~required field immutableMetaProperties missing, matches(^.*$)~invalid field immutableMetaProperties"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"required~required field immutableProperties missing, matches(^.*$)~invalid field immutableProperties"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"required~required field mutableMetaProperties missing, matches(^.*$)~invalid field mutableMetaProperties"`
	MutableProperties       string       `json:"mutableProperties" valid:"required~required field mutableProperties missing, matches(^.*$)~invalid field mutableProperties"`
}

var _ helpers.TransactionRequest = &request{}

func (request request) FromCLI(command helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (request) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	//TODO implement me
	panic("implement me")
}

func (request request) FromJSON(message json.RawMessage) (helpers.TransactionRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (request request) MakeMsg() (sdkTypes.Msg, error) {
	//TODO implement me
	panic("implement me")
}

func (request request) Validate() error {
	_, err := govalidator.ValidateStruct(request)
	return err
}

func (request request) GetBaseReq() rest.BaseReq {
	return request.BaseReq
}

func Prototype() helpers.TransactionRequest {
	return request{}
}
