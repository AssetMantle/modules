// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"github.com/AssetMantle/modules/x/metas/key"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

var (
	testDataID = baseData.NewStringData("testData").GetID()
	testKey    = key.NewKey(testDataID).(*key.Key)
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		dataID ids.DataID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve with nil", args{baseIDs.PrototypeDataID().(*baseIDs.DataID)}, &QueryRequest{key.NewKey(baseIDs.PrototypeDataID().(*baseIDs.DataID)).(*key.Key)}},
		{"+ve", args{testDataID}, &QueryRequest{testKey}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.dataID); !reflect.DeepEqual(got, tt.want) {
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
		want helpers.QueryRequest
	}{
		{"+ve", args{newQueryRequest(testDataID)}, &QueryRequest{key.NewKey(testDataID).(*key.Key)}},
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
	encodedQuery, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(newQueryRequest(testDataID))
	require.NoError(t, err)
	encodedQuery1, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(newQueryRequest(baseIDs.PrototypeDataID().(*baseIDs.DataID)))
	require.NoError(t, err)
	type fields struct {
		Key *key.Key
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
		{"+ve", fields{testKey}, args{encodedQuery}, newQueryRequest(testDataID), false},
		{"+ve with nil", fields{key.NewKey(baseIDs.PrototypeDataID()).(*key.Key)}, args{encodedQuery1}, newQueryRequest(baseIDs.PrototypeDataID().(*baseIDs.DataID)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				Key: tt.fields.Key,
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
	encodedQuery, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(newQueryRequest(testDataID))
	require.NoError(t, err)
	encodedQuery1, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(newQueryRequest(baseIDs.PrototypeDataID().(*baseIDs.DataID)))
	require.NoError(t, err)
	type fields struct {
		Key *key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve", fields{testKey}, encodedQuery, false},
		{"+ve with nil", fields{key.NewKey(baseIDs.PrototypeDataID()).(*key.Key)}, encodedQuery1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				Key: tt.fields.Key,
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.DataID})
	viper.Set(constants.DataID.GetName(), testKey)
	type fields struct {
		Key *key.Key
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testKey}, args{cliCommand, client.Context{}}, newQueryRequest(testDataID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QueryRequest{
				Key: tt.fields.Key,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.context)
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

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		Key *key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testKey}, false},
		{"+ve", fields{key.NewKey(baseIDs.PrototypeDataID()).(*key.Key)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				Key: tt.fields.Key,
			}
			if err := queryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"+ve", &QueryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
