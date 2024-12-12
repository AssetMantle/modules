// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package assets

import (
	"net/http"
	"strconv"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AssetMantle/modules/helpers"
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

	if queryRequest.Limit > constants.PaginationLimit {
		return constants.InvalidRequest.Wrapf("limit cannot be greater than %d", constants.PaginationLimit)
	}

	if queryRequest.Limit < 0 {
		return constants.InvalidRequest.Wrapf("limit cannot be less than 0")
	}

	return nil
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	assetID, err := baseIDs.PrototypeAssetID().FromString(cliCommand.ReadString(constants.AssetID))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit := cliCommand.ReadInt(constants.Limit)

	return newQueryRequest(assetID.(ids.AssetID), int32(limit)), nil
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	assetID, err := baseIDs.PrototypeAssetID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName()))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit, err := strconv.Atoi(httpRequest.URL.Query().Get(constants.Limit.GetName()))
	if err != nil {
		limit = query.DefaultLimit
	}

	return newQueryRequest(assetID.(ids.AssetID), int32(limit)), nil
}

func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(assetID ids.AssetID, limit int32) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(assetID).(*key.Key), Limit: limit}
}
