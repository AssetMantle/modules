package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(parameterList lists.ParameterList) *QueryResponse {
	return &QueryResponse{
		ParameterList: parameterList.(*baseLists.ParameterList),
	}
}
