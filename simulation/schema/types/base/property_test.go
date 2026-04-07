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
					t.Errorf("GenerateRandomProperty() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()
			got := GenerateRandomProperty(tt.args.r)
			typeName := reflect.TypeOf(got).String()
			if typeName != "*base.MesaProperty" && typeName != "*base.MetaProperty" {
				t.Errorf("GenerateRandomProperty() type = %v, want *base.MesaProperty or *base.MetaProperty", typeName)
			}
		})
	}
}
