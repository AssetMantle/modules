// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package split

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
//	ids.SplitID `json:"splitID" valid:"required~required field splitID missing"`
//}

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
	if splitID, err := baseIDs.ReadSplitID(cliCommand.ReadString(constants.SplitID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(splitID), nil
	}
}
func (*QueryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if splitID, err := baseIDs.ReadSplitID(vars[Query.GetName()]); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(splitID), nil
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
func newQueryRequest(splitID ids.SplitID) helpers.QueryRequest {
	return &QueryRequest{SplitId: splitID.(*baseIDs.SplitID)}
}
