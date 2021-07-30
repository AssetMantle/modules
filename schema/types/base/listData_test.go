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

	listValue, _ := ReadAccAddressListData("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	testListData := NewListData(listValue)
	listValue2, _ := ReadAccAddressListData("")
	testListData2 := NewListData(listValue2)

	require.Equal(t, "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c", testListData.String())
	require.Equal(t, NewID(meta.Hash("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")), testListData.GenerateHashID())
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
	require.Equal(t, nil, Error)


	listValue3, _ := ReadAccAddressListData("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	testListData3 := NewListData(listValue3)
	listValue4, _ := ReadAccAddressListData("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c")
	testListData4 := NewListData(listValue4)

	listValue6, _ := ReadAccAddressListData("cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u,cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6")
	testListData6 := NewListData(listValue6)
	//listValue8, _ := ReadAccAddressListData("cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u,cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6,cosmos1nm0rrq86ucezaf8uj35pq9fpwr5r82cl8sc7p5")
	//testListData8 := NewListData(listValue8)
	//listValue9, _ := ReadAccAddressListData("cosmos1nm0rrq86ucezaf8uj35pq9fpwr5r82cl8sc7p5")
	//testListData9 := NewListData(listValue9)
	listValue10, _ := ReadAccAddressListData("cosmos1sjllsnramtg3ewxqwwrwjxfgc4n4ef9u0tvx7u,cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6,cosmos1a3yjj7d3qnx4spgvjcwjq9cw9snrrrhu3rw8nv")
	testListData10 := NewListData(listValue10)
	listValue11, _ := ReadAccAddressListData("cosmos1a3yjj7d3qnx4spgvjcwjq9cw9snrrrhu3rw8nv")
	testListData11 := NewListData(listValue11)

	dataAsList6, Error := testListData6.AsListData()
	//dataAsList8, Error := testListData8.AsListData()
	dataAsList10, Error := testListData10.AsListData()
	dataAsList11, Error := testListData11.AsListData()

	a := dataAsList6.Add(dataAsList11)
	require.Equal(t, 0, testListData3.Compare(testListData4))
	//require.Equal(t, 0, dataAsList8.Search(testListData9))
	//require.Equal(t, 0, dataAsList8.Remove(testListData9).Compare(dataAsList6))
	require.Equal(t, 0, a.Compare(dataAsList10))

}
