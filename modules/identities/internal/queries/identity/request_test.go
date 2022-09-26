// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/schema"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func createTestInput() (ids.IdentityID, ids.IdentityID) {
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	emptyMutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData(""))))

	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentity := baseIDs.NewIdentityID(testClassificationID, immutables)

	emptyTestClassificationID := baseIDs.NewClassificationID(immutables, emptyMutables)
	emptyTestIdentity := baseIDs.NewIdentityID(emptyTestClassificationID, immutables)

	return testIdentity, emptyTestIdentity
}

func Test_newQueryRequest(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	//vars := make(map[string]string)
	//vars["identities"] = "randomString"
	//cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})
	//cliContext := context.NewCLIContext().WithCodec(Codec)
	testIdentity, emptyTestIdentity := createTestInput()

	type args struct {
		identityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{

		{"+ve", args{testIdentity}, queryRequest{testIdentity}},
		{"+ve with empty String", args{emptyTestIdentity}, queryRequest{emptyTestIdentity}},
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
		want queryRequest
	}{

		{"+ve", args{newQueryRequest(testIdentity)}, queryRequest{testIdentity}},
		{"+ve with empty string", args{newQueryRequest(emptyTestIdentity)}, queryRequest{emptyTestIdentity}},
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
	randomDecode, _ := queryRequest{emptyTestIdentity}.Encode()
	type fields struct {
		IdentityID ids.IdentityID
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
		{"+ve", fields{emptyTestIdentity}, args{randomDecode}, queryRequest{emptyTestIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
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
	byteArr, _ := common.Codec.MarshalJSON(newQueryRequest(testIdentity))
	byteArr2, _ := common.Codec.MarshalJSON(newQueryRequest(emptyTestIdentity))

	type fields struct {
		IdentityID ids.IdentityID
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
			queryRequest := queryRequest{
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
	testIdentity, emptyTestIdentity := createTestInput()
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})
	cliContext := context.NewCLIContext().WithCodec(codec.New())
	type fields struct {
		IdentityID ids.IdentityID
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

		{"+ve", fields{testIdentity}, args{cliCommand, cliContext}, queryRequest{testIdentity}, false}, //Todo: Need help
		{"+ve with empty Identity", fields{emptyTestIdentity}, args{cliCommand, cliContext}, queryRequest{emptyTestIdentity}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				IdentityID: tt.fields.IdentityID,
			}
			if got, err := queryRequest.FromCLI(tt.args.cliCommand, tt.args.in1); !reflect.DeepEqual(got, tt.want) {
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
	testIdentity, emptyTestIdentity := createTestInput()
	vars := make(map[string]string)
	vars["identities"] = "9UNIA3_tulK2vRE0nSmsHKNzhDxoCBHI4z8XXfLO1FM=.pvamJCA8talIpNPu8fekxGhvFtTGtjSRhAaaKQOrHfg="
	vars2 := make(map[string]string)
	vars2["identities"] = "qlFr8g0R-Qe6CxKcU5Ncdj7kAnSEp8Wq6sckkmznGiI=.pvamJCA8talIpNPu8fekxGhvFtTGtjSRhAaaKQOrHfg="
	type fields struct {
		IdentityID ids.IdentityID
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

		{"+ve", fields{testIdentity}, args{vars: vars}, newQueryRequest(testIdentity), false},
		{"+ve with empty IdentityID", fields{emptyTestIdentity}, args{vars: vars2}, newQueryRequest(emptyTestIdentity), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
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
		IdentityID ids.IdentityID
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
			queryRequest := queryRequest{
				IdentityID: tt.fields.IdentityID,
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
