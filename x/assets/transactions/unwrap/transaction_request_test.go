// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"encoding/json"
	"fmt"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/utilities/rest"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var (
	fromAddress       = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _ = types.AccAddressFromBech32(fromAddress)
	testBaseRequest   = rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	immutables        = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables          = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData())))
	classificationID  = baseIDs.NewClassificationID(immutables, mutables)
	fromID            = baseIDs.NewIdentityID(classificationID, immutables).(*baseIDs.IdentityID)
	coinAsset         = baseDocuments.NewCoinAsset("stake")
	testRate          = types.NewInt(100)
	assetID           = coinAsset.GetCoinAssetID().(*baseIDs.AssetID)
	coin              = types.NewCoin(coinAsset.GetDenom(), testRate)
	coins             = types.NewCoins(coin)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq rest.BaseReq
		fromID  string
		Coins   string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.AsString(), coins.String()}, transactionRequest{testBaseRequest, fromID.AsString(), coins.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.Coins); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID, constants.FromIdentityID, constants.Value})

	viper.Set(constants.FromIdentityID.GetName(), fromID.AsString())
	viper.Set(constants.AssetID.GetName(), assetID.AsString())
	viper.Set(constants.Value.GetName(), testRate.String())
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, args{cliCommand, client.Context{}.WithCodec(base.CodecPrototype())}, transactionRequest{testBaseRequest, fromID.AsString(), coins.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, fromID.AsString(), coins.String()))
	require.NoError(t, err)
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
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
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, args{jsonMessage}, newTransactionRequest(testBaseRequest, fromID.AsString(), coins.String()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
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
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, NewMessage(fromAccAddress, fromID, coins), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
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
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq rest.BaseReq
		FromID  string
		Coins   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.AsString(), coins.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq: tt.fields.BaseReq,
				FromID:  tt.fields.FromID,
				Coins:   tt.fields.Coins,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
