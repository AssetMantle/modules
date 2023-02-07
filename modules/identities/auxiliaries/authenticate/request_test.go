// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	base3 "github.com/AssetMantle/modules/schema/properties/base"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

var (
	mutableProperties      = base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData(base.NewDataList())))
	immutableProperties    = base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(base.NewDataList())))
	toMutateMetaProperties = base.NewPropertyList(base3.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewListData(base.NewDataList(base.NewStringData("Test")))))
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
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
