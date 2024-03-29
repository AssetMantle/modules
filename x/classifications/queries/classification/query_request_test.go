// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package classification

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
)

func createTestInput() *baseIDs.ClassificationID {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	return classificationID.(*baseIDs.ClassificationID)
}

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		classificationID ids.ClassificationID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve", args{createTestInput()}, &QueryRequest{createTestInput()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.classificationID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{newQueryRequest(createTestInput())}, &QueryRequest{createTestInput()}},
		{"+ve", args{newQueryRequest(baseIDs.PrototypeClassificationID())}, &QueryRequest{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}},
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
	encodedQueryResponse, err := (&QueryRequest{createTestInput()}).Encode()
	require.NoError(t, err)
	encodedQueryResponse1, err := (&QueryRequest{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}).Encode()
	require.NoError(t, err)
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
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
		{"+ve", fields{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, args{encodedQueryResponse}, &QueryRequest{createTestInput()}, false},
		{"+ve", fields{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, args{encodedQueryResponse1}, &QueryRequest{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				ClassificationID: tt.fields.ClassificationID,
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
	encodedQuery, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryRequest{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)})
	require.NoError(t, err)
	encodedQuery1, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryRequest{createTestInput()})
	require.NoError(t, err)
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve with nil", fields{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, encodedQuery, false},
		{"+ve with nil", fields{createTestInput()}, encodedQuery1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				ClassificationID: tt.fields.ClassificationID,
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.ClassificationID})

	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"+ve", fields{}, args{cliCommand, baseHelpers.TestClientContext}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QueryRequest{
				ClassificationID: tt.fields.ClassificationID,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		ClassificationID *baseIDs.ClassificationID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{createTestInput()}, false},
		{"+ve", fields{baseIDs.PrototypeClassificationID().(*baseIDs.ClassificationID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				ClassificationID: tt.fields.ClassificationID,
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
