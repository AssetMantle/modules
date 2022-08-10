package mappable

import (
	"github.com/AssetMantle/modules/schema/helpers"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		// TODO: Add test cases.
		{"+ve", identity{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
