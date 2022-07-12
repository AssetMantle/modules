// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type queryRequest struct {
	AssetID ids.ID `json:"assetID" valid:"required~required field assetID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

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
func (queryRequest queryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (queryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(baseIDs.NewStringID(cliCommand.ReadString(constants.AssetID)))
}
func (queryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(baseIDs.NewStringID(vars[Query.GetName()]))
}
func (queryRequest queryRequest) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryRequest)
}

func (queryRequest queryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := common.Codec.UnmarshalJSON(bytes, &queryRequest); err != nil {
		return nil, err
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return queryRequest{}
}
func queryRequestFromInterface(request helpers.QueryRequest) queryRequest {
	switch value := request.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}
func newQueryRequest(assetID ids.ID) helpers.QueryRequest {
	return queryRequest{AssetID: assetID}
}
