package base

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParameters(t *testing.T) {

	codec := baseTestUilities.MakeCodec()
	_, storeKeys := baseTestUilities.SetupTest(t)
	validatorFunction := func(interface{}) error { return nil }
	Parameters := NewParameters(base.NewParameter(base.NewID("testParameter1"), base.NewStringData("testData1"), validatorFunction))

	Parameters = Parameters.Initialize(params.NewSubspace(codec, storeKeys, nil, "test"))
	Parameters.ParamSetPairs()
	Parameters.GetKeyTable()
	Parameters.Equal(Parameters)
	Parameters.GetList()
	Parameters.String()
	Error := Parameters.Validate()
	require.Nil(t, Error)

	fmt.Println(Parameters.ParamSetPairs())
	fmt.Println(Parameters.GetKeyTable())
	fmt.Println(Parameters.Equal(Parameters))
	fmt.Println(Parameters.GetList())
	fmt.Println(Parameters.String())
	fmt.Println(Parameters.Validate())
	//Parameters.Fetch(context, Parameters.Get(base.NewID("testParameter1")).GetID())
	//Parameters.Mutate(context, Parameters.Get(base.NewID("")))
}
