package base

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CliFlag(t *testing.T) {

	testCliFlag := NewCLIFlag("name", "value", ",usage")
	//GetName method test
	require.Equal(t, "name", testCliFlag.GetName())
	//GetValue method test
	require.Equal(t, "value", testCliFlag.GetValue())

	//ReadCLIValue method test
	require.Equal(t, "", testCliFlag.ReadCLIValue())
	require.Equal(t, int64(0), NewCLIFlag("name", int64(-1), ",usage").ReadCLIValue())
	require.Equal(t, 0, NewCLIFlag("name", 123, ",usage").ReadCLIValue())
	require.Equal(t, false, NewCLIFlag("name", false, ",usage").ReadCLIValue())
}
