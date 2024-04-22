// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

import (
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/helpers"
)

type request struct {
	BaseReq                 rest.BaseReq `json:"baseReq"`
	FromID                  string       `json:"fromID" valid:"optional"`
	ImmutableMetaProperties string       `json:"immutableMetaProperties" valid:"optional"`
	ImmutableProperties     string       `json:"immutableProperties" valid:"optional"`
	MutableMetaProperties   string       `json:"mutableMetaProperties" valid:"optional"`
	MutableProperties       string       `json:"mutableProperties" valid:"optional"`
	ClassificationID        string       `json:"classificationID" valid:"optional"`
	MakerAssetID            string       `json:"makerAssetID" valid:"optional"`
	TakerAssetID            string       `json:"takerAssetID" valid:"optional"`
	MakerSplit              string       `json:"makerSplit" valid:"optional"`
	TakerSplit              string       `json:"takerSplit" valid:"optional"`
	ExpiresIn               string       `json:"expiresIn" valid:"optional"`
	Height                  string       `json:"height" valid:"optional"`
	TakerID                 string       `json:"takerID" valid:"optional"`
	Name                    string       `json:"name" valid:"optional"`
	Coins                   string       `json:"coins" valid:"optional"`
}

var _ helpers.Request = &request{}

func (request request) Validate() error {
	return nil
}

func (request request) GetBaseReq() rest.BaseReq {
	return request.BaseReq
}

func Prototype() helpers.Request {
	return request{}
}
