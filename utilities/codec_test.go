// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package utilities

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/require"
)

func TestMakeModuleCodec(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, MakeModuleCode(nil, nil), nil)
	})
}

func TestRegisterModuleConcrete(t *testing.T) {
	type args struct {
		codec *codec.LegacyAmino
		o     interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "positive",
			args: args{
				codec: codec.NewLegacyAmino(),
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

func TestRegisterCodec(t *testing.T) {
	require.Panics(t, func() {
		require.Equal(t, MakeMessageCodec(nil), nil)
	})
}
