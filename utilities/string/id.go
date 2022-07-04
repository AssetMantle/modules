package string

import (
	"strings"
)

// TODO write testcase for empty and singular input
func JoinIDStrings(idStrings ...string) string {
	return strings.Join(idStrings, idSeparator)
}

func SplitCompositeIDString(idString string) []string {
	return strings.Split(idString, idSeparator)
}
