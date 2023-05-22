// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/identities/common"
)

func createTestInput() (*baseIDs.IdentityID, *baseIDs.IdentityID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	emptyMutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData(""))))

	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentity := baseIDs.NewIdentityID(testClassificationID, immutables)

	emptyTestClassificationID := baseIDs.NewClassificationID(immutables, emptyMutables)
	emptyTestIdentity := baseIDs.NewIdentityID(emptyTestClassificationID, immutables)

	return testIdentity.(*baseIDs.IdentityID), emptyTestIdentity.(*baseIDs.IdentityID)
}

func Test_newQueryRequest(t *testing.T) {
	var legacyAmino = codec.NewLegacyAmino()
	schemaCodec.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	testIdentity, emptyTestIdentity := createTestInput()

	type args struct {
		identityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{

		{"+ve", args{testIdentity}, &QueryRequest{testIdentity}},
		{"+ve with empty String", args{emptyTestIdentity}, &QueryRequest{emptyTestIdentity}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequestFromInterface(t *testing.T) {
	testIdentity, emptyTestIdentity := createTestInput()
	type args struct {
		request helpers.QueryRequest
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{

		{"+ve", args{newQueryRequest(testIdentity)}, &QueryRequest{testIdentity}},
		{"+ve with empty string", args{newQueryRequest(emptyTestIdentity)}, &QueryRequest{emptyTestIdentity}},
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
	testIdentity, emptyTestIdentity := createTestInput()
	testQueryRequest := newQueryRequest(testIdentity)
	encodedRequest, err := testQueryRequest.Encode()
	require.Nil(t, err)
	randomDecode, _ := (&QueryRequest{emptyTestIdentity}).Encode()
	type fields struct {
		IdentityID *baseIDs.IdentityID
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

		{"+ve", fields{testIdentity}, args{encodedRequest}, testQueryRequest, false},
		{"+ve", fields{emptyTestIdentity}, args{randomDecode}, &QueryRequest{emptyTestIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				IdentityID: tt.fields.IdentityID,
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
	testIdentity, emptyTestIdentity := createTestInput()
	byteArr, _ := common.LegacyAmino.MarshalJSON(newQueryRequest(testIdentity))
	byteArr2, _ := common.LegacyAmino.MarshalJSON(newQueryRequest(emptyTestIdentity))

	type fields struct {
		IdentityID *baseIDs.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{

		{"+ve", fields{testIdentity}, byteArr, false},
		{"+ve with empty String ID", fields{emptyTestIdentity}, byteArr2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				IdentityID: tt.fields.IdentityID,
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
	testIdentity, _ := createTestInput()
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})

	viper.Set(constants.IdentityID.GetName(), testIdentity.String())
	type fields struct {
		IdentityID *baseIDs.IdentityID
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
		{"+ve", fields{testIdentity}, args{cliCommand, base.TestClientContext}, newQueryRequest(testIdentity), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				IdentityID: tt.fields.IdentityID,
			}
			if got, err := queryRequest.FromCLI(tt.args.cliCommand, tt.args.context); !reflect.DeepEqual(got, tt.want) {
				if (err != nil) != tt.wantErr {
					t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func Test_queryRequest_FromMap(t *testing.T) {
	vars := make(map[string]string)
	vars["identities"] = "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM="
	vars2 := make(map[string]string)
	vars2["identities"] = "qlFr8g0R-Qe6CxKcU5Ncdj7kAnSEp8Wq6sckkmznGiI="
	testIdentity, err := baseIDs.ReadIdentityID(vars["identities"])
	require.NoError(t, err)
	emptyTestIdentity, err := baseIDs.ReadIdentityID(vars2["identities"])
	require.NoError(t, err)
	type fields struct {
		IdentityID *baseIDs.IdentityID
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

		{"+ve", fields{testIdentity.(*baseIDs.IdentityID)}, args{vars: vars}, newQueryRequest(testIdentity), false},
		{"+ve with empty IdentityID", fields{emptyTestIdentity.(*baseIDs.IdentityID)}, args{vars: vars2}, newQueryRequest(emptyTestIdentity), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				IdentityID: tt.fields.IdentityID,
			}
			if got, err := queryRequest.FromMap(tt.args.vars); !reflect.DeepEqual(got, tt.want) {
				if (err != nil) != tt.wantErr {
					t.Errorf("FromMap() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromMap() got = %v, want %v", got, tt.want)
				}

			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	testIdentity, emptyTestIdentity := createTestInput()
	type fields struct {
		IdentityID *baseIDs.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{"+ve", fields{testIdentity}, false},
		{"+ve with empty IdentityID", fields{emptyTestIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				IdentityID: tt.fields.IdentityID,
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
