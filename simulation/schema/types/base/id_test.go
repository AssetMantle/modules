package base

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenerateRandomID(t *testing.T) {
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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomID() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomID(tt.args.r); reflect.TypeOf(got).String() != "base.id" {
				t.Errorf("GenerateRandomID() = %v, want base.id", got)
			}
		})
	}
}

func TestGenerateRandomIDWithDec(t *testing.T) {
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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithDec() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomIDWithDec(tt.args.r); reflect.TypeOf(got).String() != "base.id" {
				t.Errorf("GenerateRandomIDWithDec() = %v, want base.id", got)
			}
		})
	}
}

func TestGenerateRandomIDWithInt64(t *testing.T) {

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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithInt64() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomIDWithInt64(tt.args.r); reflect.TypeOf(got).String() != "base.id" {
				t.Errorf("GenerateRandomIDWithInt64() = %v, want base.id", got)
			}
		})
	}
}
