// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"github.com/AssetMantle/modules/modules/metas/module/common"
	"github.com/AssetMantle/modules/modules/metas/module/mappable"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) IsSuccessful() bool {
	return queryResponse.Success
}
func (queryResponse *QueryResponse) Encode() ([]byte, error) {
	return common.Codec.MarshalJSON(queryResponse)
}
func (queryResponse *QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := common.Codec.UnmarshalJSON(bytes, &queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection, error error) helpers.QueryResponse {
	var list []*baseHelpers.Mappable

	for _, i := range collection.GetList() {
		list = append(list, mappable.NewMappable(i.(*baseHelpers.Mappable).Impl.(*baseHelpers.Mappable_DataMappable).DataMappable.Data).(*baseHelpers.Mappable))
	}

	if error != nil {
		return &QueryResponse{
			Success: false,
			Error:   error.Error(),
			List:    list,
		}
	}
	return &QueryResponse{
		Success: true,
		Error:   "",
		List:    list,
	}
}
