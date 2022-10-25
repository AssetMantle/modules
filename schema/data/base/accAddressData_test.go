// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types"
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
		value types.AccAddress
	}
	tests := []struct {
		name string
		args args
		want data.Data
	}{
		// TODO: Add test cases.
		{"+ve", args{_fromAddress}, accAddressData{_fromAddress}},
		{"-ve with empty string", args{_fromAddress1}, accAddressData{_fromAddress1}},
		{"-ve", args{nil}, accAddressData{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewAccAddressData(tt.args.value), "NewAccAddressData(%v)", tt.args.value)
		})
	}
}

func Test_accAddressDataFromInterface(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type args struct {
		listable traits.Listable
	}
	tests := []struct {
		name    string
		args    args
		want    accAddressData
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{"-ve empty String", args{accAddressData{}}, accAddressData{}, assert.NoError},
		{"-ve wrong Address", args{stringData{}}, accAddressData{}, assert.Error},
		{"+ve", args{accAddressData{fromAccAddress}}, accAddressData{fromAccAddress}, assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accAddressDataFromInterface(tt.args.listable)
			if !tt.wantErr(t, err, fmt.Sprintf("accAddressDataFromInterface(%v)", tt.args.listable)) {
				return
			}
			assert.Equalf(t, tt.want, got, "accAddressDataFromInterface(%v)", tt.args.listable)
		})
	}
}

func Test_accAddressData_Compare(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
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
		// TODO: Add test cases.
		{"-ve empty String", fields{fromAccAddress}, args{accAddressData{}}, 1},
		{"+ve", fields{fromAccAddress}, args{accAddressData{fromAccAddress}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Compare(tt.args.listable), "Compare(%v)", tt.args.listable)
		})
	}
}

func Test_accAddressData_GenerateHashID(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"-ve empty String", fields{}, baseIDs.GenerateHashID()},
		{"+ve", fields{fromAccAddress}, baseIDs.GenerateHashID(accAddressData{fromAccAddress}.Bytes())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GenerateHashID(), "GenerateHashID()")
		})
	}
}

func Test_accAddressData_Get(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   types.AccAddress
	}{
		// TODO: Add test cases.
		{"-ve empty String", fields{}, accAddressData{}.Value},
		{"+ve", fields{fromAccAddress}, accAddressData{fromAccAddress}.Value},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Get(), "Get()")
		})
	}
}

func Test_accAddressData_GetID(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.DataID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, baseIDs.NewDataID(accAddressData{})},
		{"+ve", fields{fromAccAddress}, baseIDs.NewDataID(accAddressData{fromAccAddress})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GetID(), "GetID()")
		})
	}
}

func Test_accAddressData_GetType(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.ID
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, idsConstants.AccAddressDataID},
		{"+ve", fields{fromAccAddress}, idsConstants.AccAddressDataID},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.GetType(), "GetType()")
		})
	}
}

func Test_accAddressData_String(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, accAddressData{}.Value.String()},
		{"+ve", fields{fromAccAddress}, accAddressData{fromAccAddress}.Value.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.String(), "String()")
		})
	}
}

func Test_accAddressData_ZeroValue(t *testing.T) {
	fromAccAddress := sdkTypes.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")

	type fields struct {
		Value types.AccAddress
	}
	tests := []struct {
		name   string
		fields fields
		want   data.Data
	}{
		// TODO: Add test cases.
		{"+ve with nil", fields{}, NewAccAddressData(sdkTypes.AccAddress{})},
		{"+ve", fields{fromAccAddress}, NewAccAddressData(sdkTypes.AccAddress{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
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
		// TODO: Add test cases.
		{"+ve with nil", fields{}, accAddressData{}.Value.Bytes()},
		{"+ve", fields{_fromAddress}, accAddressData{_fromAddress}.Value.Bytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accAddressData := accAddressData{
				Value: tt.fields.Value,
			}
			assert.Equalf(t, tt.want, accAddressData.Bytes(), "Bytes()")
		})
	}
}
