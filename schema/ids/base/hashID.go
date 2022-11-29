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

//
//type hashID struct {
//	HashBytes []byte
//}

var _ ids.HashID = (*HashID)(nil)

func (hashID HashID) IsHashID() {}

func (hashID HashID) Bytes() []byte {
	return hashID.HashBytes
}
func (hashID HashID) Compare(listable traits.Listable) int {
	return bytes.Compare(hashID.Bytes(), hashIDFromInterface(listable).Bytes())
}
func hashIDFromInterface(i interface{}) HashID {
	switch value := i.(type) {
	case HashID:
		return value
	default:
		panic(constants.MetaDataError)
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
		return &HashID{HashBytes: nil}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return &HashID{HashBytes: hash.Sum(nil)}
}

func PrototypeHashID() ids.HashID {
	return GenerateHashID()
}

func ReadHashID(hashIDString string) (ids.HashID, error) {
	if hashBytes, err := base64.URLEncoding.DecodeString(hashIDString); err == nil {
		return hashID{HashBytes: hashBytes}, nil
	}

	if hashIDString == "" {
		return nil, nil
	}

	return hashID{}, constants.IncorrectFormat
}
