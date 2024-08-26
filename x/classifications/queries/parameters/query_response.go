package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func (queryResponse *QueryResponse) Encode() ([]byte, error) {
	return base.CodecPrototype().MarshalJSON(queryResponse)
}
func (queryResponse *QueryResponse) Decode(bytes []byte) (helpers.QueryResponse, error) {
	if err := base.CodecPrototype().UnmarshalJSON(bytes, queryResponse); err != nil {
		return nil, err
	}

	return queryResponse, nil
}
func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(parameterList lists.ParameterList) *QueryResponse {
	return &QueryResponse{
		ParameterList: parameterList.(*baseLists.ParameterList),
	}
}
