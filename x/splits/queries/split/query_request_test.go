// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package split

import (
	"github.com/AssetMantle/modules/x/splits/key"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
)

var (
	immutables          = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables            = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID    = baseIDs.NewClassificationID(immutables, mutables)
	testOwnerIdentityID = baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID         = baseDocuments.NewCoinAsset("OwnerID").GetCoinAssetID().(*baseIDs.AssetID)
	splitID             = baseIDs.NewSplitID(testAssetID, testOwnerIdentityID).(*baseIDs.SplitID)
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

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.SplitID})

	viper.Set(constants.SplitID.GetName(), splitID.AsString())
	type fields struct {
		SplitID *baseIDs.SplitID
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
		{"+ve", fields{splitID}, args{cliCommand, client.Context{}.WithCodec(base.CodecPrototype())}, newQueryRequest(splitID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QueryRequest{
				Key: key.NewKey(tt.fields.SplitID).(*key.Key),
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
		SplitID *baseIDs.SplitID
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
			queryRequest := &QueryRequest{
				Key: key.NewKey(tt.fields.SplitID).(*key.Key),
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
