// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package ownable

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/modules/splits/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

//type queryRequest struct {
//	ids.OwnableID `json:"ownableID" valid:"required~required field ownableID missing"`
//}

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
func (*QueryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if ownableID, err := baseIDs.ReadOwnableID(vars[Query.GetName()]); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(ownableID), nil
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
func newQueryRequest(ownableID ids.OwnableID) helpers.QueryRequest {
	return &QueryRequest{OwnableId: ownableID.(*baseIDs.OwnableID)}
}
