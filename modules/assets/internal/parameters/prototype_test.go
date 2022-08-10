package parameters

import (
	"github.com/AssetMantle/modules/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		//want helpers.Parameters
		want      string
		wantError error
	}{
		// TODO: Update test case.
		{"+ve", baseHelpers.NewParameters(dummy.Parameter).String(), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); tt.wantError != got.Validate() && !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
