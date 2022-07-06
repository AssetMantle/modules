package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/schema/ids/constansts"
)

// TODO write testcase for empty and singular input
func JoinIDStrings(idStrings ...string) string {
	return strings.Join(idStrings, constansts.IDSeparator)
}

func SplitCompositeIDString(idString string) []string {
	return strings.Split(idString, constansts.IDSeparator)
}
