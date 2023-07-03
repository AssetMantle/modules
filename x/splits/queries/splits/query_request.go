// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package splits

import (
	"github.com/AssetMantle/modules/x/splits/key"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Query split using split id
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Splits
// @Param splitID path string true "split ID"
// @Success 200 {object} queryResponse "Message for a successful query response"
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /splits/splits/{splitID} [get]
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}

func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	splitID, err := baseIDs.ReadSplitID(cliCommand.ReadString(constants.SplitID))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit := cliCommand.ReadInt(constants.Limit)

	return newQueryRequest(splitID, int32(limit)), nil
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	splitID, err := baseIDs.ReadSplitID(httpRequest.URL.Query().Get(Query.GetName()))
	if err != nil {
		return &QueryRequest{}, err
	}

	limit, err := strconv.Atoi(httpRequest.URL.Query().Get(constants.Limit.GetName()))
	if err != nil {
		limit = query.DefaultLimit
	}

	return newQueryRequest(splitID, int32(limit)), nil
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
func newQueryRequest(splitID ids.SplitID, limit int32) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(splitID).(*key.Key), Limit: limit}
}
