// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/ids"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

var (
	mutableProperties      = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	immutableProperties    = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData()))
	toMutateMetaProperties = baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(baseData.NewStringData("Test"))))
	immutables             = baseQualified.NewImmutables(immutableProperties)
	mutables               = baseQualified.NewMutables(mutableProperties)
	testClassificationID   = baseIDs.NewClassificationID(immutables, mutables)
	testFromID             = baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress            = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _      = types.AccAddressFromBech32(fromAddress)
)

func TestNewAuxiliaryRequest(t *testing.T) {
	type args struct {
		address    types.AccAddress
		identityID ids.IdentityID
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"+ve", args{fromAccAddress, testFromID}, auxiliaryRequest{fromAccAddress, testFromID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.address, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequestFromInterface(t *testing.T) {
	type args struct {
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name string
		args args
		want auxiliaryRequest
	}{
		{"+ve with nil", args{}, auxiliaryRequest{}},
		{"+ve", args{auxiliaryRequest{fromAccAddress, testFromID}}, auxiliaryRequest{fromAccAddress, testFromID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auxiliaryRequestFromInterface(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auxiliaryRequestFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	type fields struct {
		Address    types.AccAddress
		IdentityID ids.IdentityID
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{fromAccAddress, testFromID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				Address:    tt.fields.Address,
				IdentityID: tt.fields.IdentityID,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
