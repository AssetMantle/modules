/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package meta

import (
	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/metas/internal/common"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// QueryRequest godoc
// @Summary Query meta using meta id
// @Descrption Able to query the meta data
// @Accept json
// @Produce json
// @Tags Metas
// @Param metaID path string true "meta ID"
// @Success 200 {object} queryResponse "A successful query response"
// @Failure default  {object}  queryResponse "An unexpected error response."
// @Router /metas/metas/{metaID} [get]
func (queryRequest QueryRequest) Validate() error {
	_, Error := govalidator.ValidateStruct(queryRequest)
	return Error
}

func (queryRequest QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) helpers.QueryRequest {
	return newQueryRequest(base.NewID(cliCommand.ReadString(flags.MetaID)))
}

func (queryRequest QueryRequest) FromMap(vars map[string]string) helpers.QueryRequest {
	return newQueryRequest(base.NewID(vars[Query.GetName()]))
}

func (queryRequest QueryRequest) Encode(cdc codec.JSONMarshaler) ([]byte,error) {
	return cdc.MarshalJSON(&queryRequest)
}

func (queryRequest QueryRequest) Decode(cdc codec.JSONMarshaler,bytes []byte) (helpers.QueryRequest, error) {
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
	case *QueryRequest:
		return *value
	default:
		return QueryRequest{}
	}
}

func newQueryRequest(metaID types.ID) helpers.QueryRequest {
	return QueryRequest{MetaID: *base.NewID(metaID.String())}
}
