// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

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
// @Summary Search for an identity by identity ID
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Classifications
// @Param classificationID path string true "Unique identifier of an asset classification."
// @Success 200 {object} queryResponse "Message for a successful search response."
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /classifications/classifications/{classificationID} [get]
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if classificationID, err := baseIDs.ReadClassificationID(cliCommand.ReadString(constants.ClassificationID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(classificationID), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if classificationID, err := baseIDs.ReadClassificationID(httpRequest.URL.Query().Get(Query.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(classificationID), nil
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
func newQueryRequest(classificationID ids.ClassificationID) helpers.QueryRequest {
	return &QueryRequest{ClassificationID: classificationID.(*baseIDs.ClassificationID)}
}
