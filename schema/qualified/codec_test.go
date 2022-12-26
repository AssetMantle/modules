package qualified

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
)

func TestRegisterCodec(t *testing.T) {
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test for Register Codec", args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
