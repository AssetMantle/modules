package ownable

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryServer struct {
	queryKeeper
}

var _ QueryServer = queryServer{}

func (queryServer queryServer) Enquire(ctx context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	cntx := sdkTypes.UnwrapSDKContext(ctx)
	collection := queryServer.queryKeeper.mapper.NewCollection(cntx)
	value := utilities.GetOwnableTotalSplitsValue(collection, base.NewID(queryRequestFromInterface(queryRequest).OwnableID.IdString))
	response := newQueryResponse(value, nil)
	return &response, response.GetError()

}

func NewQueryServerImpl(keeper queryKeeper) QueryServer {
	return &queryServer{keeper}
}
