// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

//type queryRequest struct {
//	ids.IdentityID `json:"identityID" valid:"required~required field identityID missing"`
//}

var _ helpers.QueryRequest = (*QueryRequest)(nil)

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
func (queryRequest *QueryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(queryRequest)
	return err
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if identityID, err := baseIDs.ReadIdentityID(cliCommand.ReadString(constants.IdentityID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(identityID), nil
	}
}
func (*QueryRequest) FromMap(vars map[string]string) (helpers.QueryRequest, error) {
	if identityID, err := baseIDs.ReadIdentityID(vars[Query.GetName()]); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(identityID), nil
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
func newQueryRequest(identityID ids.IdentityID) helpers.QueryRequest {
	return &QueryRequest{IdentityID: identityID.(*baseIDs.IdentityID)}
}
