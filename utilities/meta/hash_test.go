package meta

import (
	"crypto/sha256"
	"encoding/base64"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/constants"
)

func TestHash(t *testing.T) {
	var filteredMetaList []string
	filteredMetaList = append(filteredMetaList, "123")
	sort.Strings(filteredMetaList)
	toDigest := strings.Join(filteredMetaList, constants.ToHashSeparator)
	hash := sha256.New()

	if _, err := hash.Write([]byte(toDigest)); err != nil {
		panic(err)
	}

	require.Equal(t, Hash(""), "")
	require.Equal(t, Hash("123"), base64.URLEncoding.EncodeToString(hash.Sum(nil)))
}
