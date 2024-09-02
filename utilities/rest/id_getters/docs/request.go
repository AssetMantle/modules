// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package docs

type request struct {
	FromID                  string `json:"fromID" valid:"optional"`
	ImmutableMetaProperties string `json:"immutableMetaProperties" valid:"optional"`
	ImmutableProperties     string `json:"immutableProperties" valid:"optional"`
	MutableMetaProperties   string `json:"mutableMetaProperties" valid:"optional"`
	MutableProperties       string `json:"mutableProperties" valid:"optional"`
	ClassificationID        string `json:"classificationID" valid:"optional"`
	MakerAssetID            string `json:"makerAssetID" valid:"optional"`
	TakerAssetID            string `json:"takerAssetID" valid:"optional"`
	MakerSplit              string `json:"makerSplit" valid:"optional"`
	TakerSplit              string `json:"takerSplit" valid:"optional"`
	ExpiresIn               string `json:"expiresIn" valid:"optional"`
	Height                  string `json:"height" valid:"optional"`
	TakerID                 string `json:"takerID" valid:"optional"`
	Name                    string `json:"name" valid:"optional"`
	Coins                   string `json:"coins" valid:"optional"`
}
