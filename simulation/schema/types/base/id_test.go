package base

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "*base.StringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomID() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := GenerateRandomID(tt.args.r)
			assert.Equal(t, tt.want, reflect.TypeOf(got).String())
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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "*base.StringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithDec() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := GenerateRandomIDWithDec(tt.args.r)
			assert.Equal(t, tt.want, reflect.TypeOf(got).String())
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
		{"test for id type", args{rand.New(rand.NewSource(7))}, false, "*base.StringID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("GenerateRandomIDWithInt64() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := GenerateRandomIDWithInt64(tt.args.r)
			assert.Equal(t, tt.want, reflect.TypeOf(got).String())
		})
	}
}
