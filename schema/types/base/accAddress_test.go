package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewAccAddressFromString(t *testing.T) {
	address := "persistence1xfg28czjxsf75x9th4kejrjv3n6t7wfcu6gjpe"
	accAddress := NewAccAddressFromString(address)
	require.Equal(t, address, sdkTypes.AccAddress(accAddress.GetValue()).String())
}

func TestAccAddress_String(t *testing.T) {
	address := "persistence1xfg28czjxsf75x9th4kejrjv3n6t7wfcu6gjpe"
	accAddress := NewAccAddressFromString(address)
	require.Equal(t, address, accAddress.String())
}

func TestAccAddress_GetBytes(t *testing.T) {
	address := "persistence1xfg28czjxsf75x9th4kejrjv3n6t7wfcu6gjpe"
	accAddress := NewAccAddressFromString(address)
	require.Equal(t, accAddress.GetBytes(), sdkTypes.AccAddress(accAddress.GetValue()).Bytes())
}

func TestAccAddress_AsSDKTypesAccAddress(t *testing.T) {
	address := "persistence1xfg28czjxsf75x9th4kejrjv3n6t7wfcu6gjpe"
	accAddress := NewAccAddressFromString(address)
	require.Equal(t, true, sdkTypes.AccAddress(accAddress.GetValue()).Equals(accAddress.AsSDKTypesAccAddress()))
}

func TestNewAccAddressFromSDKTypesAccAddress(t *testing.T) {
	address := "persistence1xfg28czjxsf75x9th4kejrjv3n6t7wfcu6gjpe"
	sdkAddress, _ := sdkTypes.AccAddressFromBech32(address)
	accAddress := NewAccAddressFromSDKTypesAccAddress(sdkAddress)
	require.Equal(t, true, sdkAddress.Equals(accAddress.AsSDKTypesAccAddress()))
}
