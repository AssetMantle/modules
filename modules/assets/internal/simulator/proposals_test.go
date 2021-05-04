package simulator

import (
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"reflect"
	"testing"
)

//func Test_simulateTextProposalContent(t *testing.T) {
//	r := rand.New(nil)
//	ct := context.Context(nil)
//
//	want :=types.NewTextProposal(
//		simulation.RandStringOfLength(r, 140),
//		simulation.RandStringOfLength(r, 5000),
//	)
//
//	require.Equal(t,want, simulateTextProposalContent(r *rand.Rand, _sdk.Context, _[]simulation.Account))
//}

func Test_simulator_WeightedProposalContentList(t *testing.T) {
	tests := []struct {
		name string
		want []simulation.WeightedProposalContent
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := simulator{}
			if got := si.WeightedProposalContentList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeightedProposalContentList() = %v, want %v", got, tt.want)
			}
		})
	}
}
