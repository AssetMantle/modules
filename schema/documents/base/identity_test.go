package base

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	base2 "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestNewIdentity(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	type args struct {
		classificationID ids.ClassificationID
		immutables       qualified.Immutables
		mutables         qualified.Mutables
	}
	tests := []struct {
		name string
		args args
		want documents.Identity
	}{
		{"+ve", args{classificationID, immutables, mutables}, testIdentity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIdentity(tt.args.classificationID, tt.args.immutables, tt.args.mutables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_IsProvisioned(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	testIdentity2 := NewIdentity(classificationID, immutables, mutables)
	m := testIdentity2.(documents.Identity)
	m.ProvisionAddress(fromAccAddress) // failing

	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddress sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: panic: MetaDataError fix it after
		// https://github.com/AssetMantle/modules/issues/59
		{"+ve", fields{testIdentity2}, args{fromAccAddress}, true},
		{"-ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{fromAccAddress}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := tt.fields.Document
			if got := identity.IsProvisioned(tt.args.accAddress); got != tt.want {
				t.Errorf("IsProvisioned() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetExpiry(t *testing.T) {
	_, _, _, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	type fields struct {
		Document documents.Identity
	}
	tests := []struct {
		name   string
		fields fields
		want   types.Height
	}{
		{"+ve", fields{testIdentity}, constants.ExpiryHeightProperty.GetData().(data.HeightData).Get()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := tt.fields.Document

			if got := identity.GetExpiry(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExpiry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_ProvisionAddress(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	// testIdentity.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(testIdentity.(identity).GetAuthentication().Add(accAddressesToData(fromAccAddress)...))))
	// testIdentity.(identity).Identity.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), baseData.NewListData(testIdentity.(identity).GetAuthentication().Add(accAddressesToData(fromAccAddress)...))))

	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddresses []sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   documents.Identity
	}{
		// TODO: panic: MetaDataError fix it after
		// https://github.com/AssetMantle/modules/issues/59
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.ProvisionAddress(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProvisionAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_GetAuthentication(t *testing.T) {
	classificationID, immutables, mutables, _ := createTestInput()

	type fields struct {
		Document documents.Identity
	}

	tests := []struct {
		name   string
		fields fields
		want   lists.DataList
	}{
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, base2.NewDataList(constants.AuthenticationProperty.GetData().(data.ListData).Get()...)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				tt.fields.Document,
			}

			if got := identity.GetAuthentication(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthentication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_accAddressesToData(t *testing.T) {
	type args struct {
		accAddresses []sdkTypes.AccAddress
	}
	tests := []struct {
		name string
		args args
		want []data.Data
	}{
		// {"+ve", args{}, },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accAddressesToData(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accAddressesToData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_identity_UnprovisionAddress(t *testing.T) {
	classificationID, immutables, mutables, testDocument := createTestInput()
	testIdentity := identity{testDocument}
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		Document documents.Identity
	}
	type args struct {
		accAddresses []sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   documents.Identity
	}{
		// TODO: panic: MetaDataError fix it after
		// https://github.com/AssetMantle/modules/issues/59
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.UnprovisionAddress(tt.args.accAddresses...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnprovisionAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
