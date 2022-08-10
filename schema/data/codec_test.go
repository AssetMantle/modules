package data

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"testing"
)

func TestRegisterCodec(t *testing.T) {
	//var Codec = codec.New()
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"+ve Codec", args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterCodec(tt.args.codec)
		})
	}
}
