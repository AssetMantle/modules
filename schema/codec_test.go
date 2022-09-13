package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test for Register Codec", args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterCodec(tt.args.codec)
		})
	}
}
