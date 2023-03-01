package utilities

import "regexp"

var IsValidStringData = regexp.MustCompile(`[A-Za-z0-9 ]{0,280}$`).MatchString
