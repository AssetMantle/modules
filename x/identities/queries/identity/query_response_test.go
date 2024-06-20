// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/record"
	base6 "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_newQueryResponse(t *testing.T) {
	type args struct {
		record helpers.Record
		error  error
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryResponse
	}{

		{"+ve", args{record: record.NewRecord(nil), error: nil}, &QueryResponse{}},
		{"-ve with error", args{record: record.NewRecord(nil), error: errorConstants.IncorrectFormat}, &QueryResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryResponse(tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResponse_Decode(t *testing.T) {
	testQueryResponse := newQueryResponse(record.NewRecord(nil))
	encodedResponse, _ := testQueryResponse.Encode()
	type fields struct {
		Success bool
		Error   string
		Record  *record.Record
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryResponse
		wantErr bool
	}{

		{"+ve", fields{Success: true, Error: ""}, args{bytes: encodedResponse}, testQueryResponse, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Record: tt.fields.Record,
			}
			got, err := queryResponse.Decode(tt.args.bytes)
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

func Test_queryResponse_Encode(t *testing.T) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("test"), base6.NewStringData("test"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("test"), base6.NewStringData("test"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	Record := record.NewRecord(base.NewIdentity(classificationID, immutables, mutables))
	encodedByte, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryResponse{Record: Record.(*record.Record)})
	encodedByteWithError, _err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(&QueryResponse{Record: Record.(*record.Record)})
	require.Nil(t, err)
	type fields struct {
		Success bool
		Error   error
		Record  *record.Record
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{

		{"+ve", fields{Success: true, Error: nil, Record: Record.(*record.Record)}, encodedByte, false},
		{"-ve with error", fields{Error: _err, Record: Record.(*record.Record)}, encodedByteWithError, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryResponse := &QueryResponse{
				Record: tt.fields.Record,
			}
			got, err := queryResponse.Encode()
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

func Test_responsePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryResponse
	}{

		{"+ve", &QueryResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := responsePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("responsePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
