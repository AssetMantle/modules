package meta

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/stretchr/testify/require"
	"sort"
	"strings"
	"testing"
)

func TestHash(t *testing.T) {

	var filteredMetaList []string
	filteredMetaList = append(filteredMetaList, "123")
	sort.Strings(filteredMetaList)
	toDigest := strings.Join(filteredMetaList, constants.ToHashSeparator)
	hash := sha256.New()

	if _, Error := hash.Write([]byte(toDigest)); Error != nil {
		panic(Error)
	}

	require.Equal(t, Hash(""), "")
	require.Equal(t, Hash("123"), base64.URLEncoding.EncodeToString(hash.Sum(nil)))

}
