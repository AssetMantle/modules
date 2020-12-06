package base

import (
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestQuery(t *testing.T) {
	context, storeKey := base.SetupTest(t)
	codec := base.MakeCodec()
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Query := NewQuery("test", "t", "testQuery", "test", base.TestQueryRequestPrototype,
		base.TestQueryResponsePrototype, base.TestQueryKeeperPrototype).Initialize(Mapper, parametersPrototype())

	// GetName
	Query.GetName()

	//
	_, Error := Query.HandleMessage(context, abciTypes.RequestQuery{})
	require.Nil(t, Error)

	Query.Command(codec)

	// RESTQueryHandler
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")
	Query.RESTQueryHandler(cliContext)

}
