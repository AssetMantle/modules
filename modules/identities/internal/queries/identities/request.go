// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identities

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"
	"net/http"
	"strconv"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

// type queryRequest struct {
//	ids.IdentityID `json:"identityID" valid:"required~required field identityID missing"`
// }

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
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if offset, err := strconv.Atoi(cliCommand.ReadString(constants.Offset)); err != nil {
		return &QueryRequest{}, err
	} else if limit, err := strconv.Atoi(cliCommand.ReadString(constants.PageSize)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(&query.PageRequest{Offset: uint64(offset), Limit: uint64(limit)}), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if offset, err := strconv.Atoi(httpRequest.URL.Query().Get("offset")); err != nil {
		return &QueryRequest{}, err
	} else if limit, err := strconv.Atoi(httpRequest.URL.Query().Get("pageSize")); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(&query.PageRequest{Offset: uint64(offset), Limit: uint64(limit)}), nil
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
func queryRequestFromInterface(request helpers.QueryRequest) *QueryRequest {
	switch value := request.(type) {
	case *QueryRequest:
		return value
	default:
		return &QueryRequest{}
	}
}
func newQueryRequest(pagination *query.PageRequest) helpers.QueryRequest {
	return &QueryRequest{Pagination: pagination}
}
