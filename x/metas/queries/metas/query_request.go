// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package metas

import (
	"net/http"
	"strconv"

	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/metas/key"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Search for metadata by meta ID
// @Description Able to query the metadata
// @Accept json
// @Produce json
// @Tags Metas
// @Param dataID path string true "Unique identifier of metadata value."
// @Success 200 {object} queryResponse "Message for a successful query response"
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /metas/metas/{dataID} [get]
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
	dataID, err := baseIDs.PrototypeDataID().FromString(cliCommand.ReadString(constants.DataID))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit := cliCommand.ReadInt(constants.Limit)

	return newQueryRequest(dataID.(ids.DataID), int32(limit)), nil
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	dataID, err := baseIDs.PrototypeDataID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName()))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit, err := strconv.Atoi(httpRequest.URL.Query().Get(constants.Limit.GetName()))
	if err != nil {
		limit = query.DefaultLimit
	}

	return newQueryRequest(dataID.(ids.DataID), int32(limit)), nil
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
func queryRequestFromInterface(request helpers.QueryRequest) *QueryRequest {
	switch value := request.(type) {
	case *QueryRequest:
		return value
	default:
		return &QueryRequest{}
	}
}
func newQueryRequest(dataID ids.DataID, limit int32) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(dataID).(*key.Key), Limit: limit}
}
