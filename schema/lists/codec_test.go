// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package lists

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

		{"+ve", args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}
