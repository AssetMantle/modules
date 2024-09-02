// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"net/http"
	"strconv"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/key"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Search for an identity by identity ID
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Identities
// @Param identityID path string true "Query identity using identityID"
// @Success 200 {object} queryResponse "Message for a successful response."
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /identities/identities/{identityID} [get]
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
	identityID, err := baseIDs.PrototypeIdentityID().FromString(cliCommand.ReadString(constants.IdentityID))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit := cliCommand.ReadInt(constants.Limit)

	return newQueryRequest(identityID.(ids.IdentityID), int32(limit)), nil
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	identityID, err := baseIDs.PrototypeIdentityID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName()))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit, err := strconv.Atoi(httpRequest.URL.Query().Get(constants.Limit.GetName()))
	if err != nil {
		limit = query.DefaultLimit
	}
	return newQueryRequest(identityID.(ids.IdentityID), int32(limit)), nil
}
func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(identityID ids.IdentityID, limit int32) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(identityID).(*key.Key), Limit: limit}
}
