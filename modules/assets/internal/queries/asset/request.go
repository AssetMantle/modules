// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

//type queryRequest struct {
//	ids.AssetID `json:"assetID" valid:"required~required field assetID missing"`
//}

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
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if assetID, err := baseIDs.ReadAssetID(cliCommand.ReadString(constants.AssetID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(assetID), nil
	}
}
func (*QueryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if assetID, err := baseIDs.ReadAssetID(vars[Query.GetName()]); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(assetID), nil
	}
}
func (queryRequest *QueryRequest) Encode() ([]byte, error) {
	return common.LegacyAmino.MarshalJSON(queryRequest)
}

func (queryRequest *QueryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := common.LegacyAmino.UnmarshalJSON(bytes, &queryRequest); err != nil {
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
func newQueryRequest(assetID ids.AssetID) helpers.QueryRequest {
	return &QueryRequest{AssetId: assetID.(*baseIDs.AssetID)}
}
