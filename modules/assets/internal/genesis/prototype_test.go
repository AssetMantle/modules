package genesis

import (
	"github.com/AssetMantle/modules/modules/assets/internal/key"
	"github.com/AssetMantle/modules/modules/assets/internal/mappable"
	"github.com/AssetMantle/modules/modules/assets/internal/parameters"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Genesis
	}{
		// TODO: Add test cases.
		{"+ve", baseHelpers.NewGenesis(key.Prototype, mappable.Prototype, []helpers.Mappable{}, parameters.Prototype().GetList())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Panics(t, func() {
				require.Equal(t, Prototype(), tt.want)
			})
		})
	}
}
