package queries

import (
	"github.com/AssetMantle/modules/modules/assets/internal/queries/asset"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		getString string
	}{
		// TODO: Add test cases.
		{"+ve", asset.Query.GetName(), "assets"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got.Get(tt.getString).GetName(), tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
