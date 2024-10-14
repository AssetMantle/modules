package parameters

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
)

type queryKeeper struct {
	parameterManager helpers.ParameterManager
}

var _ helpers.QueryKeeper = (*queryKeeper)(nil)

func (queryKeeper queryKeeper) Enquire(context context.Context, queryRequest helpers.QueryRequest) (helpers.QueryResponse, error) {
	return queryKeeper.Handle(context, queryRequest.(*QueryRequest))
}
func (queryKeeper queryKeeper) Handle(context context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	if err := queryRequest.Validate(); err != nil {
		return nil, err
	}
	return newQueryResponse(queryKeeper.parameterManager.Fetch(context).Get()), nil
}
func (queryKeeper queryKeeper) Initialize(_ helpers.Mapper, parameterManager helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	queryKeeper.parameterManager = parameterManager
	return queryKeeper
}

func keeperPrototype() helpers.QueryKeeper {
	return queryKeeper{}
}
