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

//type hashID struct {
//	HashBytes []byte
//}

var _ ids.HashID = (*HashID)(nil)

func (hashID *HashID) IsHashID() {}

// TODO test if nil and empty result in ""
func (hashID *HashID) EncodedString() string {
	return base64.URLEncoding.EncodeToString(hashID.IdBytes)
}
func (hashID *HashID) Bytes() []byte {
	return hashID.IdBytes
}
func (hashID *HashID) Compare(listable traits.Listable) int {
	return bytes.Compare(hashID.Bytes(), hashIDFromInterface(listable).Bytes())
}
func (hashID *HashID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_HashID{
			HashID: hashID,
		},
	}
}

func hashIDFromInterface(i interface{}) *HashID {
	switch value := i.(type) {
	case *HashID:
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
		return &HashID{IdBytes: nil}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return &HashID{IdBytes: hash.Sum(nil)}
}

func PrototypeHashID() ids.HashID {
	return GenerateHashID()
}

func ReadHashID(hashIDString string) (ids.HashID, error) {
	if hashBytes, err := base64.URLEncoding.DecodeString(hashIDString); err == nil {
		return &HashID{IdBytes: hashBytes}, nil
	}

	if hashIDString == "" {
		return nil, nil
	}

	return &HashID{}, constants.IncorrectFormat
}
