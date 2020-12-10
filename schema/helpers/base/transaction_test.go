package base

import (
	"encoding/json"
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransaction(t *testing.T) {
	codec := base.MakeCodec()
	context, storeKey, _ := base.SetupTest(t)
	Mapper := NewMapper(base.KeyPrototype, base.MappablePrototype).Initialize(storeKey)
	Transaction := NewTransaction("test", "", "", base.TestTransactionRequestPrototype, base.TestMessagePrototype,
		base.TestTransactionKeeperPrototype).InitializeKeeper(Mapper, parametersPrototype()).(transaction)

	// GetName
	require.Equal(t, "test", Transaction.GetName())

	// DecodeTransactionRequest
	message, Error := Transaction.DecodeTransactionRequest(json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Equal(t, nil, Error)
	require.Equal(t, sdkTypes.AccAddress("addr"), message.GetSigners()[0])

	// RegisterCodec : No Panics
	Transaction.RegisterCodec(codec)

	// Command : No Panics
	Transaction.Command(codec)

	// HandleMessage
	_, Error = Transaction.HandleMessage(context, message)
	require.Nil(t, Error)

	// RESTRequestHandler : No Panics
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")
	Transaction.RESTRequestHandler(cliContext)
}
