// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data"
	idsConstants "github.com/AssetMantle/modules/schema/data/constants"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/traits"
)

func TestNewAccAddressData(t *testing.T) {
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	_fromAddress := sdkTypes.AccAddress(fromAddress)

	fromAddress1 := ""
	_fromAddress1 := sdkTypes.AccAddress(fromAddress1)
	type args struct {
		value sdkTypes.AccAddress
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		{"+ve", args{_fromAddress}, &AccAddressData{_fromAddress}},
		{"-ve with empty string", args{_fromAddress1}, &AccAddressData{_fromAddress1}},
		{"-ve", args{nil}, &AccAddressData{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAccAddressData(tt.args.value), "NewAccAddressData(%v)", tt.args.value)
		})
	}
}

//
// func Test_accAddressDataFromInterface(t *testing.T) {
//	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
//	type args struct {
//		listable traits.Listable
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *AccAddressData
//		wantErr assert.ErrorAssertionFunc
//	}{
//		{"-ve nil", args{nil}, &AccAddressData{}, assert.Error},
//		{"-ve empty String", args{&AccAddressData{}}, &AccAddressData{}, assert.NoError},
//		{"-ve wrong Address", args{&StringData{}}, &AccAddressData{}, assert.Error},
//		{"+ve", args{&AccAddressData{fromAccAddress}}, &AccAddressData{fromAccAddress}, assert.NoError},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := accAddressDataFromInterface(tt.args.listable)
//			if !tt.wantErr(t, err, fmt.Sprintf("accAddressDataFromInterface(%v)", tt.args.listable)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "accAddressDataFromInterface(%v)", tt.args.listable)
//		})
//	}
// }

func Test_accAddressData_Compare(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"-ve empty String", fields{fromAccAddress}, args{&AccAddressData{}}, 1},
		{"+ve", fields{fromAccAddress}, args{&AccAddressData{fromAccAddress}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_accAddressData_GenerateHashID(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"-ve empty String", fields{}, baseIDs.GenerateHashID()},
		{"+ve", fields{fromAccAddress}, baseIDs.GenerateHashID((&AccAddressData{fromAccAddress}).Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_accAddressData_Get(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   sdkTypes.AccAddress
	}{
		{"-ve empty String", fields{}, (&AccAddressData{}).Value},
		{"+ve", fields{fromAccAddress}, (&AccAddressData{fromAccAddress}).Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Get(), "Get()")
		})
	}
}

func Test_accAddressData_GetID(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		{"+ve with nil", fields{}, baseIDs.GenerateDataID(&AccAddressData{})},
		{"+ve", fields{fromAccAddress}, baseIDs.GenerateDataID(&AccAddressData{fromAccAddress})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GetID(), "GetID()")
		})
	}
}

func Test_accAddressData_GetType(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		{"+ve with nil", fields{}, idsConstants.AccAddressDataTypeID},
		{"+ve", fields{fromAccAddress}, idsConstants.AccAddressDataTypeID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GetTypeID(), "GetTypeID()")
		})
	}
}

func Test_accAddressData_String(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve with nil", fields{}, (&AccAddressData{}).String()},
		{"+ve", fields{fromAccAddress}, (&AccAddressData{fromAccAddress}).String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.String(), "String()")
		})
	}
}

func Test_accAddressData_ZeroValue(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")

	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		{"+ve with nil", fields{}, NewAccAddressData(sdkTypes.AccAddress{})},
		{"+ve", fields{fromAccAddress}, NewAccAddressData(sdkTypes.AccAddress{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.ZeroValue(), "ZeroValue()")
		})
	}
}

func Test_accAddressData_Bytes(t *testing.T) {
	_fromAddress, err := sdkTypes.AccAddressFromBech32("cosmos1u6xn6rv07p2yzzj2rm8st04x54xe5ur0t9nl5j")
	require.Nil(t, err)
	type fields struct {
		Value sdkTypes.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve with nil", fields{}, (&AccAddressData{}).Bytes()},
		{"+ve", fields{_fromAddress}, (&AccAddressData{_fromAddress}).Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := &AccAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Bytes(), "Bytes()")
		})
	}
}
