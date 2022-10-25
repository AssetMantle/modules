// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mappable

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types"
	"github.com/AssetMantle/modules/schema/types/base"
)

func createTestInput() (types.Identity, ids.ClassificationID, qualified.Immutables, qualified.Mutables) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testIdentity := base.NewIdentity(classificationID, immutables, mutables)

	return testIdentity, classificationID, immutables, mutables
}

func TestNewIdentity(t *testing.T) {
	_, classificationID, immutables, mutables := createTestInput()
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want types.Identity
	}{
		{"+ve", args{classificationID, immutables, mutables}, mappable{Identity: base.NewIdentity(classificationID, immutables, mutables)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.NewIdentity(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Mappable
	}{
		{"+ve", mappable{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_identity_GetAuthentication(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   lists.DataList
// 	}{
// 		{"+ve", fields{testIdentity.(mappable).Document}, baseLists.NewDataList(constants.AuthenticationProperty.GetData().(data.ListData).Get()...)},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.GetAuthentication(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetAuthentication() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_identity_GetExpiry(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   types.Height
// 	}{
// 		{"+ve", fields{testIdentity.(mappable).Document}, constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.GetExpiry(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetExpiry() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_identity_GetKey(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
//
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		want   helpers.Key
// 	}{
// 		{"+ve", fields{testIdentity.(mappable).Document}, key.NewKey(baseIDs.NewIdentityID(testIdentity.(mappable).Document.GetClassificationID(), testIdentity.(mappable).Document.GetImmutables()))},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.GetKey(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetKey() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_identity_IsProvisioned(t *testing.T) {
// 	testIdentity, classificationID, immutables, mutables := createTestInput()
// 	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
// 	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
// 	require.Nil(t, err)
// 	testIdentity2 := base.NewIdentity(classificationID, immutables, mutables)
// 	m := testIdentity2.(types.Identity)
// 	m.ProvisionAddress(fromAccAddress) // failing
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	type args struct {
// 		accAddress sdkTypes.AccAddress
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   bool
// 	}{
// 		// TODO: panic: MetaDataError fix it after
// 		// https://github.com/AssetMantle/modules/issues/59
// 		{"+ve", fields{testIdentity2.(mappable).Document}, args{fromAccAddress}, true},
// 		{"-ve", fields{testIdentity.(mappable).Document}, args{fromAccAddress}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.IsProvisioned(tt.args.accAddress); got != tt.want {
// 				t.Errorf("IsProvisioned() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_identity_ProvisionAddress(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
// 	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
// 	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
// 	require.Nil(t, err)
// 	// testIdentity.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(testIdentity.(identity).GetAuthentication().Add(accAddressesToData(fromAccAddress)...))))
// 	// testIdentity.(identity).Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(testIdentity.(identity).GetAuthentication().Add(accAddressesToData(fromAccAddress)...))))
//
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	type args struct {
// 		accAddresses []sdkTypes.AccAddress
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   types.Identity
// 	}{
// 		// TODO: panic: MetaDataError fix it after
// 		// https://github.com/AssetMantle/modules/issues/59
// 		{"+ve", fields{testIdentity.(mappable).Document}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.ProvisionAddress(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ProvisionAddress() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_identity_RegisterCodec(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
//
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	type args struct {
// 		codec *codec.Codec
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		{"+ve", fields{testIdentity.(mappable).Document}, args{codec.New()}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			id := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			id.RegisterCodec(tt.args.codec)
// 		})
// 	}
// }
//
// func Test_identity_UnprovisionAddress(t *testing.T) {
// 	testIdentity, _, _, _ := createTestInput()
// 	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
// 	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
// 	require.Nil(t, err)
// 	type fields struct {
// 		Document qualified.Document
// 	}
// 	type args struct {
// 		accAddresses []sdkTypes.AccAddress
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   types.Identity
// 	}{
// 		// TODO: panic: MetaDataError fix it after
// 		// https://github.com/AssetMantle/modules/issues/59
// 		{"+ve", fields{testIdentity.(mappable).Document}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			identity := mappable{
// 				Document: tt.fields.Document,
// 			}
// 			if got := identity.UnprovisionAddress(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("UnprovisionAddress() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
