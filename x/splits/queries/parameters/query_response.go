package parameters

import (
	"github.com/AssetMantle/modules/helpers"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(parameters []parameters.Parameter) *QueryResponse {
	return &QueryResponse{
		ParameterList: baseLists.NewParameterList(parameters...).(*baseLists.ParameterList),
	}
}
