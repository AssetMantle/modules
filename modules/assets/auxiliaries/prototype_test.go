package auxiliaries

import (
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Auxiliaries
	}{
		// TODO: Add test cases.
		{"+ve", baseHelpers.NewAuxiliaries()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
