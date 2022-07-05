package simulation

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomAddresses(t *testing.T) {

	type args struct {
		r *rand.Rand
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		//TODO: check for nil case
		{"type test", args{rand.New(rand.NewSource(7))}, false},
		{"panic case", args{nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomAddresses() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomAddresses(tt.args.r); reflect.TypeOf(got).String() != "[]types.AccAddress" {
				t.Errorf("GenerateRandomAddresses() = %v, want []types.AccAddress", got)
			}
		})
	}
}

func TestRandomBool(t *testing.T) {

	type args struct {
		r *rand.Rand
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		//TODO: check for nil case
		{"test panic case", args{nil}, true},
		{"+ve case", args{rand.New(rand.NewSource(7))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("RandomBool() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := RandomBool(tt.args.r); reflect.TypeOf(got).String() != "bool" {
				t.Errorf("RandomBool() = %v, want bool", got)
			}
		})
	}
}
