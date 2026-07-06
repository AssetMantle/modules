// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package balances

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
)

var (
	testIdentityID = baseIDs.NewIdentityID(baseIDs.PrototypeClassificationID(), baseQualified.NewImmutables(baseLists.NewPropertyList()))
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		identityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"valid", args{testIdentityID}, newQueryRequest(testIdentityID)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newQueryRequest(tt.args.identityID)
			assert.Equal(t, tt.want, got, "newQueryRequest()")
		})
	}
}

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.IdentityID})

	viper.Set(constants.IdentityID.GetName(), testIdentityID.AsString())

	qu := &QueryRequest{
		IdentityID: testIdentityID.(*baseIDs.IdentityID),
	}
	got, err := qu.FromCLI(cliCommand, client.Context{}.WithCodec(base.CodecPrototype()))
	require.NoError(t, err, "FromCLI()")
	assert.Equal(t, newQueryRequest(testIdentityID), got, "FromCLI()")
}

func Test_queryRequest_Validate(t *testing.T) {
	// Prototype identity has nil internal key, causing Validate to panic
	require.Panics(t, func() {
		queryRequest := &QueryRequest{
			IdentityID: testIdentityID.(*baseIDs.IdentityID),
		}
		_ = queryRequest.Validate()
	})

	// Nil IdentityID also panics during validation
	require.Panics(t, func() {
		nilRequest := &QueryRequest{}
		_ = nilRequest.Validate()
	})
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"valid", &QueryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := requestPrototype()
			assert.Equal(t, tt.want, got, "requestPrototype()")
		})
	}
}
