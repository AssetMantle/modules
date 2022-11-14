// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/utilities/transaction"
)

func createTestInput(t *testing.T) (rest.BaseReq, string, ids.IdentityID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	testToAddress := "cosmos1vx8knpllrj7n963p9ttd80w47kpacrhuts497x"
	return testBaseReq, testToAddress, testFromID
}

func Test_newTransactionRequest(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type args struct {
		baseReq    rest.BaseReq
		fromID     string
		identityID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		// TODO: Add test cases.
		{"+ve with nil", args{}, transactionRequest{}},
		{"+ve", args{testBaseReq, testToAddress, testFromID.String()}, transactionRequest{testBaseReq, testToAddress, testFromID.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
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
	types.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.IdentityID})
	cliContext := context.NewCLIContext().WithCodec(Codec)
	testBaseReq, _, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
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
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, testFromID.String(), testFromID.String()}, args{cliCommand: cliCommand, cliContext: cliContext}, transactionRequest{cliCommand.ReadBaseReq(cliContext), cliCommand.ReadString(constants.FromID), cliCommand.ReadString(constants.IdentityID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
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
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
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
		{"+ve", fields{testBaseReq, testToAddress, testFromID.String()}, args{types.MustSortJSON(transaction.RegisterCodec(messagePrototype).MustMarshalJSON(message{types.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"), testFromID, testFromID}))}, transactionRequest{testBaseReq, testToAddress, testFromID.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
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
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, testToAddress, testFromID.String()}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	testBaseReq, _, testFromID := createTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve", fields{testBaseReq, testFromID.String(), testFromID.String()}, message{fromAccAddress, testFromID, testFromID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
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
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
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
		{"+ve with nil", fields{}, args{codec.New()}},
		{"+ve", fields{testBaseReq, testToAddress, testFromID.String()}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, true},
		{"+ve", fields{testBaseReq, testToAddress, testFromID.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
