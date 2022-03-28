/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package ownable

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

// QueryRequest godoc
// @Summary Query asset using asset id
// @Descrption Able to query the asset
// @Accept json
// @Produce json
// @Tags Splits
// @Param ownableID path string true "ownable ID"
// @Success 200 {object} queryResponse "A succesful query response"
// @Failure default  {object}  queryResponse "An unexpected error response."
// @Router /ownable/{ownableID} [get]
var _ helpers.QueryRequest = (*QueryRequest)(nil)

func (queryRequest QueryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(queryRequest)
	return Error
}

func (queryRequest QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(flags.OwnableID)))
}

func (queryRequest QueryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[Query.GetName()]))
}
func (queryRequest QueryRequest) Encode(cdc codec.JSONMarshaler) ([]byte, error) {
	return cdc.MarshalJSON(&queryRequest)
}

func (queryRequest QueryRequest) Decode(cdc codec.JSONMarshaler, bytes []byte) (helpers.QueryRequest, error) {
	if Error := cdc.UnmarshalJSON(bytes, &queryRequest); Error != nil {
		return nil, Error
	}

	return queryRequest, nil
}

func (queryRequest QueryRequest) LegacyAminoEncode() ([]byte, error) {
	return common.LegacyAminoCodec.MarshalJSON(queryRequest)
}

func (queryRequest QueryRequest) LegacyAminoDecode(bytes []byte) (helpers.QueryRequest, error) {
	if Error := common.LegacyAminoCodec.UnmarshalJSON(bytes, &queryRequest); Error != nil {
		return nil, Error
	}

	return queryRequest, nil
}
func requestPrototype() helpers.QueryRequest {
	return QueryRequest{}
}

func queryRequestFromInterface(request helpers.QueryRequest) QueryRequest {
	switch value := request.(type) {
	case QueryRequest:
		return value
	default:
		return QueryRequest{}
	}
}

func newQueryRequest(ownableID types.ID) helpers.QueryRequest {
	return QueryRequest{OwnableID: *base.NewID(ownableID.String())}
}
