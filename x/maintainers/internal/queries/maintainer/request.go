// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainer

import (
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"net/http"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Search for a maintainer by maintainer ID
// @Description Able to query the maintainers details
// @Accept json
// @Produce json
// @Tags Maintainers
// @Param maintainerID path string true "Unique identifier of a maintainer."
// @Success 200 {object} queryResponse "Message for a successful query response"
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /maintainers/maintainers/{maintainerID} [get]
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}

func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if maintainerID, err := baseIDs.ReadMaintainerID(cliCommand.ReadString(constants.MaintainerID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(maintainerID), nil
	}
}

func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if maintainerID, err := baseIDs.ReadMaintainerID(httpRequest.URL.Query().Get(Query.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(maintainerID), nil
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

func newQueryRequest(maintainerID ids.MaintainerID) helpers.QueryRequest {
	return &QueryRequest{MaintainerID: maintainerID.(*baseIDs.MaintainerID)}
}
