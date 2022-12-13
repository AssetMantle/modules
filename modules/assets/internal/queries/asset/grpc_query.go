package asset

import (
	"context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type queryServer struct {
	queryKeeper
}

var _ QueryServer = queryServer{}

func (queryServer queryServer) Enquire(ctx context.Context, queryRequest *QueryRequest) (*QueryResponse, error) {
	cntx := sdkTypes.UnwrapSDKContext(ctx)

	keyr := key.FromID(base.NewID(queryRequest.AssetID.String()))

	collection := queryServer.queryKeeper.mapper.NewCollection(cntx)
	response := newQueryResponse(collection.Fetch(keyr), nil)
	return &response, response.GetError()

}

func NewQueryServerImpl(keeper queryKeeper) QueryServer {
	return &queryServer{keeper}
}
