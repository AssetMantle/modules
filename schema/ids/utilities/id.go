package utilities

import (
	"strings"
)

const IDSeparator = "."

// TODO write testcase for empty and singular input
func JoinIDStrings(idStrings ...string) string {
	return strings.Join(idStrings, IDSeparator)
}

func SplitCompositeIDString(idString string) []string {
	return strings.Split(idString, IDSeparator)
}
