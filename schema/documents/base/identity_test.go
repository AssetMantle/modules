package base

import (
	"fmt"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
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
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	testIdentity := NewIdentity(classificationID, immutables, mutables)
	testIdentity.ProvisionAddress(fromAccAddress)

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
		{"+ve", fields{testIdentity}, args{fromAccAddress}, true},
		{"-ve", fields{testIdentity}, args{fromAccAddress2}, false},
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
	testIdentity2 := testIdentity
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	fromAccAddress2, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	testIdentity.Document = testIdentity.Document.Mutate(baseProperties.NewMetaProperty(constants.AuthenticationProperty.GetKey(), base.NewListData(testIdentity.GetAuthentication().Add(accAddressesToData([]sdkTypes.AccAddress{fromAccAddress}...)...))))
	fmt.Println("TEST:	", testIdentity.IsProvisioned(fromAccAddress))
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
		want   bool
	}{
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress}}, testIdentity.IsProvisioned(fromAccAddress)},
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, args{[]sdkTypes.AccAddress{fromAccAddress2}}, testIdentity2.IsProvisioned(fromAccAddress)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			identity := identity{
				Document: tt.fields.Document,
			}
			if got := identity.ProvisionAddress(tt.args.accAddresses...).IsProvisioned(tt.args.accAddresses[0]); !reflect.DeepEqual(got, tt.want) {
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
		{"+ve", fields{identity{NewDocument(classificationID, immutables, mutables)}}, baseLists.NewDataList(constants.AuthenticationProperty.GetData().(data.ListData).Get()...)},
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
