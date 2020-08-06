/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryRequest struct {
	IdentityID types.ID `json:"identityID" valid:"required~required field identityID missing"`
}

var _ helpers.QueryRequest = (*queryRequest)(nil)

func (QueryRequest queryRequest) FromCLI(cliCommand helpers.CLICommand, _ context.CLIContext) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(constants.IdentityID)))
}

func (QueryRequest queryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[constants.IdentityID.GetName()]))
}

func queryRequestPrototype() helpers.QueryRequest {
	return queryRequest{}
}

func queryRequestFromInterface(QueryRequest helpers.QueryRequest) queryRequest {
	switch value := QueryRequest.(type) {
	case queryRequest:
		return value
	default:
		return queryRequest{}
	}
}

func newQueryRequest(identityID types.ID) helpers.QueryRequest {
	return queryRequest{IdentityID: identityID}
}
