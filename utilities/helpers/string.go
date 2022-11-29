package helpers

import (
	"strings"

	"github.com/AssetMantle/modules/schema/lists/constants"
)

// TODO write testcase for empty and singular input
func JoinListStrings(listStrings ...string) string {
	return strings.Join(listStrings, constants.ListStringSeparator)
}

func SplitListString(listString string) []string {
	return strings.Split(listString, constants.ListStringSeparator)
}
