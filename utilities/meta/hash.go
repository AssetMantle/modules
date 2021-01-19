/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package meta

import (
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"

	"github.com/persistenceOne/persistenceSDK/constants"
)

func Hash(meta ...string) string {
	sort.Strings(meta)
	toDigest := strings.Join(meta, constants.ToHashSeparator)
	hash := sha256.New()

	if _, Error := hash.Write([]byte(toDigest)); Error != nil {
		panic(Error)
	}

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
