package queries

import (
	"github.com/AssetMantle/modules/modules/identities/internal/queries/identity"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"reflect"
	"testing"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Queries
	}{
		// TODO: Add test cases.
		{"+ve", baseHelpers.NewQueries(identity.Query)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got.GetList(), tt.want.GetList()) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
