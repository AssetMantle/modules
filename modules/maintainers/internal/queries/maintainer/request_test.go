// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainer

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/maintainers/internal/common"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
)

var (
	testMaintainerID, _ = createTestData()
	testMaintainerID1   = baseIds.PrototypeMaintainerID()
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		maintainerID ids.MaintainerID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve", args{testMaintainerID}, queryRequest{testMaintainerID}},
		{"+ve with nil", args{testMaintainerID1}, queryRequest{testMaintainerID1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.maintainerID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{newQueryRequest(testMaintainerID)}, queryRequest{testMaintainerID}},
		{"+ve with nil", args{newQueryRequest(testMaintainerID1)}, queryRequest{testMaintainerID1}},
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
	encodedQuery, err := queryRequest{testMaintainerID}.Encode()
	require.NoError(t, err)
	encodedQuery1, err := queryRequest{testMaintainerID1}.Encode()
	require.NoError(t, err)
	type fields struct {
		MaintainerID ids.MaintainerID
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
		{"+ve", fields{testMaintainerID}, args{encodedQuery}, queryRequest{testMaintainerID}, false},
		{"+ve", fields{testMaintainerID1}, args{encodedQuery1}, queryRequest{testMaintainerID1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				MaintainerID: tt.fields.MaintainerID,
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
	encodedQuery, err := common.Codec.MarshalJSON(queryRequest{testMaintainerID})
	require.NoError(t, err)
	encodedQuery1, err := common.Codec.MarshalJSON(queryRequest{testMaintainerID1})
	require.NoError(t, err)
	type fields struct {
		MaintainerID ids.MaintainerID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve", fields{testMaintainerID}, encodedQuery, false},
		{"+ve with nil", fields{testMaintainerID1}, encodedQuery1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				MaintainerID: tt.fields.MaintainerID,
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
	type fields struct {
		MaintainerID ids.MaintainerID
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				MaintainerID: tt.fields.MaintainerID,
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
	vars[Query.GetName()] = testMaintainerID.String()
	vars1 := make(map[string]string)
	vars1[Query.GetName()] = testMaintainerID.String()
	type fields struct {
		MaintainerID ids.MaintainerID
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
		{"+ve", fields{testMaintainerID}, args{vars}, newQueryRequest(testMaintainerID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				MaintainerID: tt.fields.MaintainerID,
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
		MaintainerID ids.MaintainerID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testMaintainerID}, false},
		{"+ve with nil", fields{testMaintainerID1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				MaintainerID: tt.fields.MaintainerID,
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
