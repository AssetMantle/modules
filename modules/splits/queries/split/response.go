/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package split

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
)

type queryResponse struct {
	Splits mappers.Splits `json:"splits" valid:"required~required field splits missing"`
}

var _ helpers.QueryResponse = (*queryResponse)(nil)

func queryResponsePrototype() helpers.QueryResponse {
	return queryResponse{}
}

func newQueryResponse(splits mappers.Splits) helpers.QueryResponse {
	return queryResponse{Splits: splits}
}
