// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"

	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

type queryRequest struct {
	IdentityID ids.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

// Validate godoc
// @Summary Search for an identity by identity ID
// @Description Able to query the asset
// @Accept json
// @Produce json
// @Tags Identities
// @Param identityID path string true "Query identity using identityID"
// @Success 200 {object} queryResponse "Message for a successful response."
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /identities/identities/{identityID} [get]
func (queryRequest queryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (queryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(baseIDs.NewID(cliCommand.ReadString(constants.IdentityID)))
}
func (queryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(baseIDs.NewID(vars[Query.GetName()]))
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
func newQueryRequest(identityID ids.ID) helpers.QueryRequest {
	return queryRequest{IdentityID: identityID}
}
