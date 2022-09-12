// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/AssetMantle/modules/modules/identities/internal/common"
	"github.com/AssetMantle/modules/schema"
	"github.com/AssetMantle/modules/schema/helpers"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_newQueryRequest(t *testing.T) {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()
	vars := make(map[string]string)
	vars["identities"] = "randomString"
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})
	cliContext := context.NewCLIContext().WithCodec(Codec)

	type args struct {
		identityID ids.ID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{

		{"+ve", args{baseIDs.NewID("randomString")}, queryRequest{}.FromMap(vars)},
		{"+ve with empty String", args{baseIDs.NewID("")}, queryRequest{}.FromCLI(cliCommand, cliContext)},
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
	type args struct {
		request helpers.QueryRequest
	}
	tests := []struct {
		name string
		args args
		want queryRequest
	}{

		{"+ve", args{newQueryRequest(baseIDs.NewID("IdentityID"))}, queryRequest{baseIDs.NewID("IdentityID")}},
		{"+ve with empty string", args{newQueryRequest(baseIDs.NewID(""))}, queryRequest{baseIDs.NewID("")}},
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
	testQueryRequest := newQueryRequest(baseIDs.NewID("IdentityID"))
	encodedRequest, err := testQueryRequest.Encode()
	require.Nil(t, err)
	randomDecode, _ := queryRequest{baseIDs.NewID("")}.Encode()
	type fields struct {
		IdentityID ids.ID
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

		{"+ve", fields{baseIDs.NewID("IdentityID")}, args{encodedRequest}, testQueryRequest, false},
		{"+ve", fields{baseIDs.NewID("")}, args{randomDecode}, queryRequest{baseIDs.NewID("")}, false},
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
	byteArr, _ := common.Codec.MarshalJSON(newQueryRequest(baseIDs.NewID("IdentityID")))
	byteArr2, _ := common.Codec.MarshalJSON(newQueryRequest(baseIDs.NewID("")))

	type fields struct {
		IdentityID ids.ID
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{

		{"+ve", fields{baseIDs.NewID("IdentityID")}, byteArr, false},
		{"+ve with empty String ID", fields{baseIDs.NewID("")}, byteArr2, false},
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
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})
	cliContext := context.NewCLIContext().WithCodec(codec.New())
	type fields struct {
		IdentityID ids.ID
	}
	type args struct {
		cliCommand helpers.CLICommand
		in1        context.CLIContext
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.QueryRequest
	}{

		{"+ve", fields{baseIDs.NewID("IdentityID")}, args{cliCommand, cliContext}, queryRequest{}.FromCLI(cliCommand, cliContext)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				IdentityID: tt.fields.IdentityID,
			}
			if got := queryRequest.FromCLI(tt.args.cliCommand, tt.args.in1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_FromMap(t *testing.T) {
	vars := make(map[string]string)
	vars["identities"] = "randomString"
	type fields struct {
		IdentityID ids.ID
	}
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.QueryRequest
	}{

		{"+ve", fields{baseIDs.NewID("IdentityID")}, args{vars: vars}, newQueryRequest(baseIDs.NewID(vars[Query.GetName()]))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := queryRequest{
				IdentityID: tt.fields.IdentityID,
			}
			if got := queryRequest.FromMap(tt.args.vars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		IdentityID ids.ID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{"+ve", fields{baseIDs.NewID("IdentityID")}, false},
		{"-ve with empty String", fields{baseIDs.NewID("")}, true},
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
