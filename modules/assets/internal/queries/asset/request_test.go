// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/assets/internal/common"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	immutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID = baseIDs.NewClassificationID(immutables, mutables)
	testAssetID      = baseIDs.NewAssetID(classificationID, immutables)
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		assetID ids.AssetID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve", args{testAssetID}, newQueryRequest(testAssetID)},
		{"+ve", args{baseIDs.PrototypeAssetID()}, newQueryRequest(baseIDs.PrototypeAssetID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequestFromInterface(t *testing.T) {
	type args struct {
		request helpers.QueryRequest
	}
	tests := []struct {
		name string
		args args
		want queryRequest
	}{
		{"+ve", args{newQueryRequest(testAssetID)}, newQueryRequest(testAssetID).(queryRequest)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryRequestFromInterface(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryRequestFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Decode(t *testing.T) {
	encodedQuery, err := common.Codec.MarshalJSON(newQueryRequest(testAssetID))
	require.NoError(t, err)
	encodedQuery1, err := common.Codec.MarshalJSON(newQueryRequest(baseIDs.PrototypeAssetID()))
	require.NoError(t, err)
	type fields struct {
		AssetID ids.AssetID
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testAssetID}, args{encodedQuery}, newQueryRequest(testAssetID), false},
		{"+ve", fields{baseIDs.PrototypeAssetID()}, args{encodedQuery1}, newQueryRequest(baseIDs.PrototypeAssetID()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				AssetID: tt.fields.AssetID,
			}
			got, err := queryRequest.Decode(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Encode(t *testing.T) {
	encodedQuery, err := common.Codec.MarshalJSON(newQueryRequest(testAssetID))
	require.NoError(t, err)
	encodedQuery1, err := common.Codec.MarshalJSON(newQueryRequest(baseIDs.PrototypeAssetID()))
	require.NoError(t, err)
	type fields struct {
		AssetID ids.AssetID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve", fields{testAssetID}, encodedQuery, false},
		{"+ve with nil", fields{baseIDs.PrototypeAssetID()}, encodedQuery1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				AssetID: tt.fields.AssetID,
			}
			got, err := queryRequest.Encode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID})
	viper.Set(constants.AssetID.GetName(), testAssetID.String())
	type fields struct {
		AssetID ids.AssetID
	}
	type args struct {
		cliCommand helpers.CLICommand
		in1        context.CLIContext
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testAssetID}, args{cliCommand, context.NewCLIContext()}, newQueryRequest(testAssetID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				AssetID: tt.fields.AssetID,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.in1)
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

func Test_queryRequest_FromMap(t *testing.T) {
	vars := make(map[string]string)
	vars[Query.GetName()] = testAssetID.String()
	type fields struct {
		AssetID ids.AssetID
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testAssetID}, args{vars}, newQueryRequest(testAssetID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				AssetID: tt.fields.AssetID,
			}
			got, err := qu.FromMap(tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		AssetID ids.AssetID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testAssetID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				AssetID: tt.fields.AssetID,
			}
			if err := queryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"+ve", queryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
