package utilities

import "regexp"

var IsValidStringID = regexp.MustCompile(`[A-Za-z0-9_]{1,250}$`).MatchString
