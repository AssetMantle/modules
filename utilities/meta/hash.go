/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package meta

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/constants"
	"sort"
	"strings"
)

func Hash(meta ...string) string {
	sort.Strings(meta)
	toDigest := strings.Join(meta, constants.ToHashSeparator)
	hash := sha1.New()
	hash.Write([]byte(toDigest))
	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
