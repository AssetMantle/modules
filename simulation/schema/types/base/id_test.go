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
		want      string
	}{
		// TODO: check for nil case
		{"test for panic case", args{nil}, true, ""},
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "base.stringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomID() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomID(tt.args.r); !reflect.DeepEqual(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("GenerateRandomID() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
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
		want      string
	}{
		// TODO: check for nil case
		{"test for panic case", args{nil}, true, ""},
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "base.stringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithDec() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomIDWithDec(tt.args.r); !reflect.DeepEqual(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("GenerateRandomIDWithDec() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
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
		want      string
	}{
		// TODO: check for nil case
		{"test for panic case", args{nil}, true, ""},
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "base.stringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithInt64() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			if got := GenerateRandomIDWithInt64(tt.args.r); !reflect.DeepEqual(reflect.TypeOf(got).String(), tt.want) {
				t.Errorf("GenerateRandomIDWithInt64() = %v, want %v", reflect.TypeOf(got).String(), got)
			}
		})
	}
}
