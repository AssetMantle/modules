package base

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"testing"
)

func TestParameters(t *testing.T) {

	codec := baseTestUilities.MakeCodec()
	_, storeKeys := baseTestUilities.SetupTest(t)
	validatorFunction := func(interface{}) error { return nil }
	Parameters := NewParameters(base.NewParameter(base.NewID(""), base.NewStringData(""), validatorFunction))

	Parameters.Initialize(params.NewSubspace(codec, storeKeys, nil, "test"))
}
