// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package split

import (
	"reflect"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/modules/splits/internal/common"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIds "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIds.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIds.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIds.NewIdentityID(classificationID, immutables)
	testOwnableID       = baseIds.NewOwnableID(baseIds.NewStringID("OwnerID"))
	splitID             = baseIds.NewSplitID(testOwnerIdentityID, testOwnableID)
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		splitID ids.SplitID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve", args{splitID}, newQueryRequest(splitID)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.splitID); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", args{newQueryRequest(splitID)}, newQueryRequest(splitID).(queryRequest)},
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
	encodedReq, err := common.Codec.MarshalJSON(newQueryRequest(splitID))
	require.NoError(t, err)
	encodedReq1, err1 := common.Codec.MarshalJSON(newQueryRequest(baseIds.PrototypeSplitID()))
	require.NoError(t, err1)
	type fields struct {
		SplitID ids.SplitID
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
		{"+ve", fields{splitID}, args{encodedReq}, newQueryRequest(splitID), false},
		{"+ve", fields{baseIds.PrototypeSplitID()}, args{encodedReq1}, newQueryRequest(baseIds.PrototypeSplitID()), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				SplitID: tt.fields.SplitID,
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
	encodedReq, err := common.Codec.MarshalJSON(newQueryRequest(splitID))
	require.NoError(t, err)
	encodedReq1, err1 := common.Codec.MarshalJSON(newQueryRequest(baseIds.PrototypeSplitID()))
	require.NoError(t, err1)
	type fields struct {
		SplitID ids.SplitID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{"+ve", fields{splitID}, encodedReq, false},
		{"+ve", fields{baseIds.PrototypeSplitID()}, encodedReq1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				SplitID: tt.fields.SplitID,
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
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.SplitID})
	viper.Set(constants.SplitID.GetName(), splitID.String())
	type fields struct {
		SplitID ids.SplitID
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
		{"+ve", fields{splitID}, args{cliCommand, context.NewCLIContext()}, newQueryRequest(splitID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				SplitID: tt.fields.SplitID,
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
	vars[Query.GetName()] = splitID.String()
	type fields struct {
		SplitID ids.SplitID
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
		{"+ve", fields{splitID}, args{vars}, newQueryRequest(splitID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := queryRequest{
				SplitID: tt.fields.SplitID,
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
		SplitID ids.SplitID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{splitID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				SplitID: tt.fields.SplitID,
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
