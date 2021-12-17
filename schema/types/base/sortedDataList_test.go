package base

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/schema/types"
)

func Test_sortedDataList(t *testing.T) {

	addr1 := NewAccAddressData(sdkTypes.AccAddress("addr1"))
	addr2 := NewAccAddressData(sdkTypes.AccAddress("addr2"))
	addr3 := NewAccAddressData(sdkTypes.AccAddress("addr3"))
	addr4 := NewAccAddressData(sdkTypes.AccAddress("addr4"))

	testSortedDataList := types.SortedDataList(sortedDataList{})
	require.Equal(t, []types.Data{}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr1)
	require.Equal(t, []types.Data{addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr1)
	require.Equal(t, []types.Data{addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr2)
	require.Equal(t, []types.Data{addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr1)
	require.Equal(t, []types.Data{addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr2)
	require.Equal(t, []types.Data{addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr3)
	require.Equal(t, []types.Data{addr3, addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr3)
	require.Equal(t, []types.Data{addr3, addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr2)
	require.Equal(t, []types.Data{addr3, addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Add(addr4)
	require.Equal(t, []types.Data{addr4, addr3, addr2, addr1}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Remove(addr1)
	require.Equal(t, []types.Data{addr4, addr3, addr2}, testSortedDataList.GetList())
	testSortedDataList = testSortedDataList.Remove(addr3)
	require.Equal(t, []types.Data{addr4, addr2}, testSortedDataList.GetList())
	require.Equal(t, 1, testSortedDataList.Search(addr2))

}
