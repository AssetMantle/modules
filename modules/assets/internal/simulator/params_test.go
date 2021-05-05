package simulator

import (
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/persistenceOne/persistenceSDK/modules/assets/internal/common"
	_ "github.com/persistenceOne/persistenceSDK/schema/types/base"
	"testing"
)

func Test_simulator_ParamChangeList(t *testing.T) {
	//s2 := rand.NewSource(42)
	//r2 := rand.New(s2)
	//tbytes := func(r *rand.Rand) string {
	//	 r = r2
	//	bytes, Error := common.Codec.MarshalJSON(dummy.Parameter.Mutate(base.NewDecData(sdk.NewDecWithPrec(int64(r.Intn(99)), 2))).GetData())
	//	log.Println(bytes)
	//	if Error != nil {
	//		panic(Error)
	//	}
	//	return string(bytes)
	//}
	//require.Equal(t, simulator{}.ParamChangeList(r2), []simulation.ParamChange{
	//	simulation.NewSimParamChange(module.Name,
	//		dummy.ID.String(),tbytes),})

}
