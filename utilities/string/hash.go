// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package string

import (
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"
)

func Hash(toHashStrings ...string) string {
	var nonEmptyStrings []string

	for _, value := range toHashStrings {
		if value != "" {
			nonEmptyStrings = append(nonEmptyStrings, value)
		}
	}

	if len(nonEmptyStrings) == 0 {
		return ""
	}

	sort.Strings(nonEmptyStrings)
	toDigest := strings.Join(nonEmptyStrings, "")
	hash := sha256.New()

	if _, err := hash.Write([]byte(toDigest)); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
