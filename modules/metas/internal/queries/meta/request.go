// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/AssetMantle/modules/modules/metas/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type queryRequest struct {
	ids.DataID `json:"dataID" valid:"required~required field dataID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

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
func (queryRequest queryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) (helpers.QueryRequest, error) {
	if dataID, err := baseIDs.ReadDataID(cliCommand.ReadString(constants.DataID)); err != nil {
		return queryRequest{}, err
	} else {
		return newQueryRequest(dataID), nil
	}
}
func (queryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if dataID, err := baseIDs.ReadDataID(vars[Query.GetName()]); err != nil {
		return queryRequest{}, err
	} else {
		return newQueryRequest(dataID), nil
	}
}
func (queryRequest queryRequest) Encode() ([]byte, error) {
	return common.LegacyAmino.MarshalJSON(queryRequest)
}
func (queryRequest queryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if err := common.LegacyAmino.UnmarshalJSON(bytes, &queryRequest); err != nil {
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
func newQueryRequest(dataID ids.DataID) helpers.QueryRequest {
	return queryRequest{DataID: dataID}
}
