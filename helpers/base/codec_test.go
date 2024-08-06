package base

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/cosmos/cosmos-sdk/client"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockTxConfig struct {
	client.TxConfig
}

func Test_codec_Initialize(t *testing.T) {
	type fields struct {
		interfaceRegistry types.InterfaceRegistry
		TxConfig          client.TxConfig
		legacyAmino       *sdkCodec.LegacyAmino
		ProtoCodec        *sdkCodec.ProtoCodec
	}
	type args struct {
		moduleManager helpers.ModuleManager
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Codec
	}{
		{
			name: "Test codec Initialize",
			fields: fields{
				interfaceRegistry: types.NewInterfaceRegistry(),
				TxConfig:          mockTxConfig{},
				legacyAmino:       sdkCodec.NewLegacyAmino(),
				ProtoCodec:        sdkCodec.NewProtoCodec(types.NewInterfaceRegistry()),
			},
			args: args{
				moduleManager: NewModuleManager(),
			},
			want: codec{
				interfaceRegistry: types.NewInterfaceRegistry(),
				TxConfig:          mockTxConfig{},
				legacyAmino:       sdkCodec.NewLegacyAmino(),
				ProtoCodec:        sdkCodec.NewProtoCodec(types.NewInterfaceRegistry()),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := codec{
				interfaceRegistry: tt.fields.interfaceRegistry,
				TxConfig:          tt.fields.TxConfig,
				legacyAmino:       tt.fields.legacyAmino,
				ProtoCodec:        tt.fields.ProtoCodec,
			}
			Codec := codec.Initialize(tt.args.moduleManager)
			assert.Equalf(t, tt.want.GetProtoCodec(), Codec.GetProtoCodec(), "Initialize(%v)", tt.args.moduleManager)
		})
	}
}

func TestCodecPrototype(t *testing.T) {
	tests := []struct {
		name              string
		expectedInterface helpers.Codec
	}{
		{
			name:              "codec prototype",
			expectedInterface: CodecPrototype(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := CodecPrototype()
			assert.IsType(t, test.expectedInterface, got)
			assert.NotNil(t, got.GetProtoCodec())
			assert.NotNil(t, got.GetLegacyAmino())
			assert.NotNil(t, got.InterfaceRegistry())
		})
	}
}
