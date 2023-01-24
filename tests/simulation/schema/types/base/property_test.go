package base

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomProperty(t *testing.T) {
	type args struct {
		r *rand.Rand
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		// TODO: check for nil case
		{"test for panic case", args{nil}, "", true},
		{"test for type", args{rand.New(rand.NewSource(7))}, "base.mesaProperty", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomProperty() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomProperty(tt.args.r); !reflect.DeepEqual(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("GenerateRandomProperty() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
			}
		})
	}
}
