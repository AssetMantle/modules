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

	listValue, Error := ReadAccAddressListData("address1")
	require.NotNil(t, Error)
	listValue, Error = ReadAccAddressListData("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")
	require.Nil(t, Error)
	testListData := NewListData(listValue)
	listValue2, Error := ReadAccAddressListData("")
	require.Nil(t, Error)
	testListData2 := NewListData(listValue2)


	require.Equal(t, "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae", testListData.String())
	require.Equal(t, NewID(meta.Hash("cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae")), testListData.GenerateHashID())
  
	require.Equal(t, NewID(""), testListData2.GenerateHashID())
	require.Equal(t, NewID("LD"), testListData.GetTypeID())

	dataAsString, Error := testListData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsHeight, Error := testListData.AsHeight()
	require.Equal(t, height{}, dataAsHeight)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsDec, Error := testListData.AsDec()
	require.Equal(t, sdkTypes.ZeroDec(), dataAsDec)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsID, Error := testListData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.IncorrectFormat, Error)

	dataAsList, Error := testListData.AsListData()
	require.Equal(t, testListData, dataAsList)

	require.Nil(t, Error)

	require.Equal(t, testListData.(listData).Value, testListData.Get())

	data, Error := ReadAccAddressListData("testString")
	require.Equal(t, listData{}, data)
	require.NotNil(t, Error)

	require.Panics(t, func() {
		require.Equal(t, false, testListData.Compare(NewStringData("")) == 0)
	})
	require.Panics(t, func() {
		require.Equal(t, false, testListData.Compare(NewHeightData(NewHeight(123))) == 0)
	})

	listValue3, Error := ReadAccAddressListData("cosmos1ce2gkxuug6g388qd535tk3p70ej2xkkvf5jm6r")
	require.Nil(t, Error)
	testListData3 := NewListData(listValue3)
	listValue4, Error := ReadAccAddressListData("cosmos1ce2gkxuug6g388qd535tk3p70ej2xkkvf5jm6r")
	require.Nil(t, Error)
	testListData4 := NewListData(listValue4)

	listValue6, Error := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry,cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	require.Nil(t, Error)
	listValue8, Error := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry")
	require.Nil(t, Error)
	dataAsList8, Error := listValue8.AsListData()
	require.Nil(t, Error)
	addr1, _ := sdkTypes.AccAddressFromBech32("cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	addr2, _ := sdkTypes.AccAddressFromBech32("cosmos1adf0nwjhg2anlfy5t7m4ztxvczhn342kvq806c")
	dataAsList8 = dataAsList8.Add(NewAccAddressData(addr1), NewAccAddressData(addr2))
	listValue9, Error := ReadAccAddressListData("cosmos1adf0nwjhg2anlfy5t7m4ztxvczhn342kvq806c")
	require.Nil(t, Error)
	listValue10, Error := ReadAccAddressListData("cosmos1lmx8c6dujhgt04a3f9wzx503pp763dgvuga8ry,cosmos1tqcxq4xxwjc3wtn6hqqc5f7nfyqz86ktv6hssp,cosmos1zz22dfpvw3zqpeyhvhmx944a588fgcalw744ts")
	require.Nil(t, Error)
	listValue11, Error := ReadAccAddressListData("cosmos1tqcxq4xxwjc3wtn6hqqc5f7nfyqz86ktv6hssp")
	require.Nil(t, Error)

	dataAsList6, Error := listValue6.AsListData()
	require.Nil(t, Error)
	dataAsList9, Error := listValue9.AsListData()
	require.Nil(t, Error)
	dataAsList10, Error := listValue10.AsListData()
	require.Nil(t, Error)


	a := dataAsList6.Add(dataAsList11)
	require.Equal(t, 0, testListData3.Compare(testListData4))

	require.Equal(t, false, dataAsList8.Search(dataAsList9.(listData).Value.GetList()[0]) == len(dataAsList8.(listData).Value.GetList()))
	require.Equal(t, dataAsList6, dataAsList8.Remove(listValue9.(listData).Value.GetList()[0]))
	require.Equal(t, dataAsList10, dataAsList6.Add(listValue11.(listData).Value.GetList()[0]))


}
