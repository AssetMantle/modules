// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"net/http"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/key"
)

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
	if err := queryRequest.Key.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf(err.Error())
	}

	return nil
}
func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if identityID, err := baseIDs.PrototypeIdentityID().FromString(cliCommand.ReadString(constants.IdentityID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(identityID.(ids.IdentityID)), nil
	}
}
func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if identityID, err := baseIDs.PrototypeIdentityID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(identityID.(ids.IdentityID)), nil
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

func newQueryRequest(identityID ids.IdentityID) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(identityID).(*key.Key)}
}
