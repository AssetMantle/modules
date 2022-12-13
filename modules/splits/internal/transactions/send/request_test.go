// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package send

import (
	"encoding/json"
	"fmt"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
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
	mutables          = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(baseLists.NewDataList()))))
	classificationID  = baseIDs.NewClassificationID(immutables, mutables)
	fromID            = baseIDs.NewIdentityID(classificationID, immutables)
	ownableID         = baseIDs.NewOwnableID(baseIDs.NewStringID("ownableid"))
	testRate          = types.NewDec(1)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		baseReq   rest.BaseReq
		fromID    string
		toID      string
		ownableID string
		value     string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, transactionRequest{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.toID, tt.args.ownableID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ToID, constants.FromID, constants.OwnableID, constants.Value})
	cliContext := context.NewCLIContext().WithCodec(codec.New()).WithFromAddress(fromAccAddress).WithChainID("test")
	viper.Set(constants.FromID.GetName(), fromID.String())
	viper.Set(constants.ToID.GetName(), fromID.String())
	viper.Set(constants.OwnableID.GetName(), ownableID.String())
	viper.Set(constants.Value.GetName(), testRate.String())
	type fields struct {
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
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
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, args{cliCommand, cliContext}, transactionRequest{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.cliContext)
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
	jsonMessage, err := json.Marshal(newTransactionRequest(testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()))
	require.NoError(t, err)
	type fields struct {
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
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
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, args{jsonMessage}, transactionRequest{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
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
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, testBaseRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, newMessage(fromAccAddress, fromID, fromID, ownableID, testRate), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
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
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
	}
	type args struct {
		codec *codec.Codec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, args{codec.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			tr.RegisterCodec(tt.args.codec)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		BaseReq   rest.BaseReq
		FromID    string
		ToID      string
		OwnableID string
		Value     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testBaseRequest, fromID.String(), fromID.String(), ownableID.String(), testRate.String()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:   tt.fields.BaseReq,
				FromID:    tt.fields.FromID,
				ToID:      tt.fields.ToID,
				OwnableID: tt.fields.OwnableID,
				Value:     tt.fields.Value,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
