// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	"net/http"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/classifications/key"
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
	if err := queryRequest.Key.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf(err.Error())
	}

	return nil
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if classificationID, err := baseIDs.PrototypeClassificationID().FromString(cliCommand.ReadString(constants.ClassificationID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(classificationID.(ids.ClassificationID)), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if classificationID, err := baseIDs.PrototypeClassificationID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(classificationID.(ids.ClassificationID)), nil
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
func newQueryRequest(classificationID ids.ClassificationID) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(classificationID).(*key.Key)}
}
