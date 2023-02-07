package codec

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
)

func TestRegisterModuleConcrete(t *testing.T) {
	type args struct {
		legacyAmino *codec.LegacyAmino
		o           interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive",
			args: args{
				legacyAmino: codec.NewLegacyAmino(),
				o:           args{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterModuleConcrete(tt.args.legacyAmino, tt.args.o)
		})
	}
}
