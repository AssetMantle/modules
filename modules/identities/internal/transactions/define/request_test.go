// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"encoding/json"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/utilities"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"reflect"
	"testing"
)

func Test_newTransactionRequest(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})
	//cliContext := context.NewCLIContext().WithCodec(Codec)

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	//immutableMetaProperties, err := utilities.ReadMetaProperties(immutableMetaPropertiesString)
	//require.Equal(t, nil, err)
	//immutableProperties, err := utilities.ReadProperties(immutablePropertiesString)
	//require.Equal(t, nil, err)
	//mutableMetaProperties, err := utilities.ReadMetaProperties(mutableMetaPropertiesString)
	//require.Equal(t, nil, err)
	//mutableProperties, err := utilities.ReadProperties(mutablePropertiesString)
	//require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	//fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	//require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type args struct {
		baseReq                 rest.BaseReq
		fromID                  string
		immutableMetaProperties string
		immutableProperties     string
		mutableMetaProperties   string
		mutableProperties       string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{

		{"+ve", args{testBaseReq, "fromID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString}, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.immutableMetaProperties, tt.args.immutableProperties, tt.args.mutableMetaProperties, tt.args.mutableProperties); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{

		{"+ve", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromCLI(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.ImmutableMetaProperties, constants.ImmutableProperties, constants.MutableMetaProperties, constants.MutableProperties})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	type args struct {
		cliCommand helpers.CLICommand
		cliContext context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{

		{"+ve", fields{BaseReq: testBaseReq, FromID: "", ImmutableMetaProperties: "", ImmutableProperties: "", MutableMetaProperties: "", MutableProperties: ""}, args{cliCommand, cliContext}, transactionRequest{cliCommand.ReadBaseReq(cliContext), cliCommand.ReadString(constants.FromID), cliCommand.ReadString(constants.ImmutableMetaProperties), cliCommand.ReadString(constants.ImmutableProperties), cliCommand.ReadString(constants.MutableMetaProperties), cliCommand.ReadString(constants.MutableProperties)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	jsonMessage, _ := json.Marshal(transactionRequest{testBaseReq, "fromID", immutableMetaPropertiesString, immutablePropertiesString, mutableMetaPropertiesString, mutablePropertiesString})

	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	type args struct {
		rawMessage json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, args{jsonMessage}, transactionRequest{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			got, err := transactionRequest.FromJSON(tt.args.rawMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	immutableMetaProperties, err := utilities.ReadMetaProperties(immutableMetaPropertiesString)
	require.Equal(t, nil, err)
	immutableProperties, err := utilities.ReadProperties(immutablePropertiesString)
	require.Equal(t, nil, err)
	mutableMetaProperties, err := utilities.ReadMetaProperties(mutableMetaPropertiesString)
	require.Equal(t, nil, err)
	mutableProperties, err := utilities.ReadProperties(mutablePropertiesString)
	require.Equal(t, nil, err)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name    string
		fields  fields
		want    sdkTypes.Msg
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, newMessage(fromAccAddress, baseIDs.NewID("fromID"), immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			got, err := transactionRequest.MakeMsg()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	immutableMetaPropertiesString := "defaultImmutableMeta1:S|defaultImmutableMeta1"
	immutablePropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutableMetaPropertiesString := "defaultMutableMeta1:S|defaultMutableMeta1"
	mutablePropertiesString := "defaultMutable1:S|defaultMutable1"

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"

	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: sdkTypes.NewCoins()}
	type fields struct {
		BaseReq                 rest.BaseReq
		FromID                  string
		ImmutableMetaProperties string
		ImmutableProperties     string
		MutableMetaProperties   string
		MutableProperties       string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{BaseReq: testBaseReq, FromID: "fromID", ImmutableMetaProperties: immutableMetaPropertiesString, ImmutableProperties: immutablePropertiesString, MutableMetaProperties: mutableMetaPropertiesString, MutableProperties: mutablePropertiesString}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:                 tt.fields.BaseReq,
				FromID:                  tt.fields.FromID,
				ImmutableMetaProperties: tt.fields.ImmutableMetaProperties,
				ImmutableProperties:     tt.fields.ImmutableProperties,
				MutableMetaProperties:   tt.fields.MutableMetaProperties,
				MutableProperties:       tt.fields.MutableProperties,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
