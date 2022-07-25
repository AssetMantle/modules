package base

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"sort"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type hashID struct {
	HashBytes []byte
}

var _ ids.HashID = (*hashID)(nil)

// TODO test
func (hashID hashID) String() string {
	hash := sha256.New()

	if _, err := hash.Write(hashID.HashBytes); err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(hash.Sum(nil))
}
func (hashID hashID) Bytes() []byte {
	return hashID.HashBytes
}
func (hashID hashID) Compare(listable traits.Listable) int {
	if compareHashID, err := hashIDFromInterface(listable); err != nil {
		panic(err)
	} else {
		return bytes.Compare(hashID.Bytes(), compareHashID.Bytes())
	}
}
func hashIDFromInterface(i interface{}) (hashID, error) {
	switch value := i.(type) {
	case hashID:
		return value, nil
	default:
		return hashID{}, constants.MetaDataError
	}
}

// TODO test
func GenerateHashID(toHashList ...[]byte) ids.HashID {
	var nonEmptyByteList [][]byte

	for _, value := range toHashList {
		if len(value) != 0 {
			nonEmptyByteList = append(nonEmptyByteList, value)
		}
	}

	if len(nonEmptyByteList) == 0 {
		return hashID{HashBytes: []byte{}}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return hashID{HashBytes: hash.Sum(nil)}
}
