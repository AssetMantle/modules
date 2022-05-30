// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package string

import (
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"

	"github.com/AssetMantle/modules/constants"
)

func Hash(meta ...string) string {
	var filteredMetaList []string

	for _, value := range meta {
		if value != "" {
			filteredMetaList = append(filteredMetaList, value)
		}
	}

	if len(filteredMetaList) == 0 {
		return ""
	}

	sort.Strings(filteredMetaList)
	toDigest := strings.Join(filteredMetaList, constants.ToHashSeparator)
	hash := sha256.New()

	if _, err := hash.Write([]byte(toDigest)); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
