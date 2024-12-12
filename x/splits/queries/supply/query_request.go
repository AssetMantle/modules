// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package supply

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Query asset using asset id
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Splits
// @Param assetID path string true "identity ID"
// @Success 200 {object} queryRequest "Message for a successful query response"
// @Failure default  {object}  queryRequest "Message for an unexpected error response."
// @Router /balances/{identityID} [get]
func (queryRequest *QueryRequest) Validate() error {
	if err := queryRequest.AssetID.ValidateBasic(); err != nil {
		return err
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
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(assetID ids.AssetID) helpers.QueryRequest {
	return &QueryRequest{AssetID: assetID.(*baseIDs.AssetID)}
}
