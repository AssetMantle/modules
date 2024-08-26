// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"net/http"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/key"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Search for an asset by Asset ID
// @Description Unique identifier of an asset.
// @Accept text/plain
// @Produce json
// @Tags Assets
// @Param assetID path string true "Asset ID"
// @Success 200 {object} queryResponse "Message for a successful search."
// @Failure default  {object}  queryResponse "Message for an unexpected error."
// @Router /assets/assets/{assetID} [get]
func (queryRequest *QueryRequest) Validate() error {
	if err := queryRequest.Key.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf(err.Error())
	}

	return nil
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if assetID, err := baseIDs.PrototypeAssetID().FromString(cliCommand.ReadString(constants.AssetID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(assetID.(ids.AssetID)), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if assetID, err := baseIDs.PrototypeAssetID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(assetID.(ids.AssetID)), nil
	}
}
func (queryRequest *QueryRequest) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryRequest)
}

func (queryRequest *QueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryRequest); err != nil {
		return nil, err
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(assetID ids.AssetID) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(assetID).(*key.Key)}
}
