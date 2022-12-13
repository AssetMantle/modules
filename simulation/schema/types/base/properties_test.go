package base

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomProperties(t *testing.T) {

	type args struct {
		r *rand.Rand
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		// TODO: check for nil case
		{"test for panic case", args{nil}, true},
		{"test for type", args{rand.New(rand.NewSource(7))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomPropertyList() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			if got := GenerateRandomPropertyList(tt.args.r); reflect.TypeOf(got).String() != "base.propertyList" {
				t.Errorf("GenerateRandomPropertyList() = %v, want base.propertyList", got)
			}
		})
	}
}
