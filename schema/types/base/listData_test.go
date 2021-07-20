/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"

)

func Test_ListData(t *testing.T) {

	listValue, _ := ReadAccAddressListData("address1")
	testListData := NewListData(listValue)
	listValue2, _ := ReadAccAddressListData("")
	testListData2 := NewListData(listValue2)

	require.Equal(t, "address1", testListData.String())
	require.Equal(t, NewID(meta.Hash("address1")), testListData.GenerateHashID())
	require.Equal(t, NewID(""), testListData2.GenerateHashID())
	require.Equal(t, NewID("LD"), testListData.GetTypeID())

	dataAsString, Error := testListData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsHeight, Error := testListData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, nil, Error)

	dataAsDec, Error := testListData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsID, Error := testListData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsList, Error := testListData.AsListData()
	require.Equal(t, listValue, dataAsList)
	require.Equal(t, errors.IncorrectFormat, Error)

	require.Equal(t, listValue, testListData.Get())

	data, Error := ReadAccAddressListData("testString")
	require.Equal(t, nil, data)
	require.NotNil(t, Error)

	require.Equal(t, false, testListData.Equal(NewStringData("")))
	require.Equal(t, false, testListData.Equal(NewHeightData(NewHeight(123))))

	listValue3, _ := ReadAccAddressListData("address3")
	testListData3 := NewListData(listValue3)
	listValue4, _ := ReadAccAddressListData("address3")
	testListData4 := NewListData(listValue4)

	listValue6, _ := ReadAccAddressListData("address6, address7")
	testListData6 := NewListData(listValue6)
	listValue8, _ := ReadAccAddressListData("address6, address7,address8")
	testListData8 := NewListData(listValue8)
	listValue9, _ := ReadAccAddressListData("address8")
	testListData9 := NewListData(listValue9)
	listValue10, _ := ReadAccAddressListData("address6, address7,address9")
	testListData10 := NewListData(listValue10)
	listValue11, _ := ReadAccAddressListData("address9")
	testListData11 := NewListData(listValue11)

	dataAsList6, Error := testListData6.AsListData()
	dataAsList8, Error := testListData8.AsListData()
	dataAsList10, Error := testListData10.AsListData()

	require.Equal(t, 0, testListData3.Compare(testListData4))
	require.Equal(t, true, dataAsList8.IsPresent(testListData9))
	require.Equal(t, dataAsList6, dataAsList8.Remove(testListData9))
	require.Equal(t, dataAsList10, dataAsList6.Add(testListData11))

}
