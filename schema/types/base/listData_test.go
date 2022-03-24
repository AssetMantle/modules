/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"reflect"
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants"

	"github.com/persistenceOne/persistenceSDK/utilities/random"

	"github.com/persistenceOne/persistenceSDK/schema/types"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
)

func Test_ListData(t *testing.T) {

	_, err := ReadAccAddressListData("address1")
	require.NotNil(t, err)

	_, err = ReadAccAddressListData("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")
	require.Nil(t, err)

	var listValue types.Data
	testListData := NewListData(listValue)
	listValue2, err := ReadAccAddressListData("")
	require.Nil(t, err)

	testListData2 := NewListData(listValue2)

	require.Equal(t, "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae", testListData.String())
	require.Equal(t, NewID(meta.Hash("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")), testListData.GenerateHashID())

	require.Equal(t, NewID(""), testListData2.GenerateHashID())
	require.Equal(t, listDataID, testListData.GetTypeID())

	dataAsString, err := testListData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsHeight, err := testListData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsDec, err := testListData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsID, err := testListData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, err)

	dataAsList, err := testListData.AsListData()
	require.Equal(t, testListData, dataAsList)

	require.Nil(t, err)

	require.Equal(t, testListData.(listData).Value, testListData.Get())

	data, err := ReadAccAddressListData("testString")
	require.Equal(t, listData{}, data)
	require.NotNil(t, err)

	require.Panics(t, func() {
		require.Equal(t, false, testListData.Compare(NewStringData("")) == 0)
	})

	require.Panics(t, func() {
		require.Equal(t, false, testListData.Compare(NewHeightData(NewHeight(123))) == 0)
	})

	listValue3, err := ReadAccAddressListData("cosmos1ce2gkxuug6g388qd535tk3p70ej2xkkvf5jm6r")
	require.Nil(t, err)
	testListData3 := NewListData(listValue3)

	listValue4, err := ReadAccAddressListData("cosmos1ce2gkxuug6g388qd535tk3p70ej2xkkvf5jm6r")
	require.Nil(t, err)
	testListData4 := NewListData(listValue4)

	listValue6, err := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry,cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	require.Nil(t, err)

	listValue8, err := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry")
	require.Nil(t, err)
	dataAsList8, err := listValue8.AsListData()
	require.Nil(t, err)

	addr1, _ := sdkTypes.AccAddressFromBech32("cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	addr2, _ := sdkTypes.AccAddressFromBech32("cosmos1adf0nwjhg2anlfy5t7m4ztxvczhn342kvq806c")
	dataAsList8 = dataAsList8.Add(NewAccAddressData(addr1), NewAccAddressData(addr2))

	listValue9, err := ReadAccAddressListData("cosmos1adf0nwjhg2anlfy5t7m4ztxvczhn342kvq806c")
	require.Nil(t, err)

	listValue10, err := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry,cosmos1tqcxq4xxwjc3wtn6hqqc5f7nfyqz86ktv6hssp,cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	require.Nil(t, err)

	listValue11, err := ReadAccAddressListData("cosmos1tqcxq4xxwjc3wtn6hqqc5f7nfyqz86ktv6hssp")
	require.Nil(t, err)

	dataAsList6, err := listValue6.AsListData()
	require.Nil(t, err)

	dataAsList9, err := listValue9.AsListData()
	require.Nil(t, err)

	dataAsList10, err := listValue10.AsListData()
	require.Nil(t, err)

	require.Equal(t, 0, testListData3.Compare(testListData4))

	require.Equal(t, false, dataAsList8.Search(dataAsList9.(listData).Value.GetList()[0]) == len(dataAsList8.(listData).Value.GetList()))
	require.Equal(t, dataAsList6, dataAsList8.Remove(listValue9.(listData).Value.GetList()[0]))
	require.Equal(t, dataAsList10, dataAsList6.Add(listValue11.(listData).Value.GetList()[0]))
}

func Test_listData_GenerateHashID(t *testing.T) {
	randomString := random.GenerateUniqueIdentifier()

	type fields struct {
		Value sortedDataList
	}
	tests := []struct {
		name   string
		fields fields
		want   types.ID
	}{
		{
			name:   "hash of nil list",
			fields: fields{Value: nil},
			want:   NewID(""),
		},
		{
			name:   "hash of empty list",
			fields: fields{Value: sortedDataList{}},
			want:   NewID(""),
		},
		{
			name:   "hash of single string data",
			fields: fields{Value: []types.Data{NewStringData(randomString)}},
			want:   NewStringData(randomString).GenerateHashID(),
		},
		{
			name:   "hash of single zero value string data",
			fields: fields{Value: []types.Data{NewStringData(randomString).ZeroValue()}},
			want:   NewStringData(randomString).ZeroValue().GenerateHashID(),
		},
		{
			name:   "hash of two string data",
			fields: fields{Value: []types.Data{NewStringData(randomString), NewStringData(randomString)}},
			want:   NewID(NewStringData(randomString).GenerateHashID().String() + constants.ListHashStringSeparator + NewStringData(randomString).GenerateHashID().String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listData := listData{
				Value: tt.fields.Value,
			}
			if got := listData.GenerateHashID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateHashID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListData(t *testing.T) {
	randomString := random.GenerateUniqueIdentifier()

	type args struct {
		value []types.Data
	}
	tests := []struct {
		name string
		args args
		want types.Data
	}{
		{
			name: "nil argument",
			args: args{nil},
			want: listData{},
		},
		{
			name: "empty arguments",
			args: args{},
			want: listData{},
		},
		{
			name: "zero arguments",
			args: args{[]types.Data{}},
			want: listData{},
		},
		{
			name: "one string data argument",
			args: args{[]types.Data{NewStringData(randomString)}},
			want: listData{[]types.Data{NewStringData(randomString)}},
		},
		{
			name: "two unsorted string data arguments",
			args: args{[]types.Data{NewStringData("a"), NewStringData("b")}},
			want: listData{[]types.Data{NewStringData("b"), NewStringData("a")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListData(tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewListData() = %v, want %v", got, tt.want)
			}
		})
	}
}
