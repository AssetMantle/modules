// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	"net/http"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Query asset using asset id
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Splits
// @Param ownableID path string true "ownable ID"
// @Success 200 {object} queryRequest "Message for a successful query response"
// @Failure default  {object}  queryRequest "Message for an unexpected error response."
// @Router /ownable/{ownableID} [get]
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}

func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if ownableID, err := baseIDs.ReadOwnableID(cliCommand.ReadString(constants.OwnableID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(ownableID), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if ownableID, err := baseIDs.ReadOwnableID(httpRequest.URL.Query().Get(Query.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(ownableID), nil
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
func newQueryRequest(ownableID ids.OwnableID) helpers.QueryRequest {
	return &QueryRequest{OwnableID: ownableID.ToAnyOwnableID().(*baseIDs.AnyOwnableID)}
}
