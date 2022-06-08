package codec

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
)

func TestRegisterModuleConcrete(t *testing.T) {
	type args struct {
		codec *codec.Codec
		o     interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive",
			args: args{
				codec: codec.New(),
				o:     args{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterModuleConcrete(tt.args.codec, tt.args.o)
		})
	}
}
