package base

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"sort"

	"buf.build/gen/go/assetmantle/schema/protocolbuffers/go/schema/ids/base"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

//
// type hashID struct {
//	IdBytes []byte
// }
type hashID base.HashID

func (hashID *hashID) String() string {
	return (hashID).String()
}

var _ ids.HashID = (*hashID)(nil)

func (hashID *hashID) IsHashID() {}

func (hashID *hashID) Bytes() []byte {
	return hashID.IdBytes
}
func (hashID *hashID) Compare(listable traits.Listable) int {
	return bytes.Compare(hashID.Bytes(), hashIDFromInterface(listable).Bytes())
}
func hashIDFromInterface(i interface{}) *hashID {
	switch value := i.(type) {
	case *hashID:
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
		return &hashID{IdBytes: nil}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return &hashID{IdBytes: hash.Sum(nil)}
}

func PrototypeHashID() ids.HashID {
	return GenerateHashID()
}

func ReadHashID(hashIDString string) (ids.HashID, error) {
	if hashBytes, err := base64.URLEncoding.DecodeString(hashIDString); err == nil {
		return &hashID{IdBytes: hashBytes}, nil
	}

	if hashIDString == "" {
		return nil, nil
	}

	return &hashID{}, constants.IncorrectFormat
}
