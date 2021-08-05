/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainer

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryRequest struct {
	MaintainerID types.ID `json:"maintainerID" valid:"required~required field maintainerID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

// QueryRequest godoc
// @Summary Query maintainers using maintainer id
// @Descrption Able to query the maintainers details
// @Accept json
// @Produce json
// @Tags Maintainers
// @Param maintainerID path string true "maintainer ID"
// @Success 200 {object} queryResponse "A successful query response"
// @Failure default  {object}  queryResponse "An unexpected error response."
// @Router /maintainers/maintainers/{maintainerID} [get]
func (queryRequest queryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(queryRequest)
	return Error
}

func (queryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(flags.MaintainerID)))
}

func (queryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[Query.GetName()]))
}

func (queryRequest queryRequest) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryRequest)
}

func (queryRequest queryRequest) Decode(bytes []byte) (helpers.QueryRequest, error) {
	if Error := common.Codec.UnmarshalJSON(bytes, &queryRequest); Error != nil {
		return nil, Error
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

func newQueryRequest(maintainerID types.ID) helpers.QueryRequest {
	return queryRequest{MaintainerID: maintainerID}
}
