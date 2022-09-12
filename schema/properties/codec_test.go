// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

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

		{"+ve", args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterCodec(tt.args.codec)
		})
	}
}
