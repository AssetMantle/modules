package base

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"github.com/AssetMantle/modules/schema/ids/constants"
	"sort"
	"strings"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// type hashID struct {
//	HashBytes []byte
// }

var _ ids.HashID = (*HashID)(nil)

func (hashID *HashID) IsHashID() {}
func (hashID *HashID) ValidateBasic() error {
	if len(hashID.IDBytes) != 32 && hashID.IDBytes != nil {
		return errorConstants.IncorrectFormat
	}
	return nil
}
func (hashID *HashID) GetTypeID() ids.StringID {
	return NewStringID(constants.HashIDType)
}
func (hashID *HashID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeHashID(), nil
	}

	if hashBytes, err := base64.URLEncoding.DecodeString(idString); err != nil {
		return PrototypeHashID(), err
	} else {
		return &HashID{
			IDBytes: hashBytes,
		}, nil
	}
}

// TODO test if nil and empty result in ""
func (hashID *HashID) AsString() string {
	return base64.URLEncoding.EncodeToString(hashID.IDBytes)
}
func (hashID *HashID) Bytes() []byte {
	return hashID.IDBytes
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
		panic(errorConstants.IncorrectFormat.Wrapf("expected *HashID, got %T", i))
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
		return &HashID{IDBytes: nil}
	}

	sort.Slice(nonEmptyByteList, func(i, j int) bool { return bytes.Compare(nonEmptyByteList[i], nonEmptyByteList[j]) == -1 })

	hash := sha256.New()

	// TODO check if nil elements in slice
	if _, err := hash.Write(bytes.Join(nonEmptyByteList, nil)); err != nil {
		panic(err)
	}

	return &HashID{IDBytes: hash.Sum(nil)}
}

func PrototypeHashID() ids.HashID {
	return GenerateHashID()
}

func ReadHashID(hashIDString string) (ids.HashID, error) {
	if hashBytes, err := base64.URLEncoding.DecodeString(hashIDString); err == nil {
		return &HashID{IDBytes: hashBytes}, nil
	}

	if hashIDString == "" {
		return PrototypeHashID(), nil
	}

	return PrototypeHashID(), errorConstants.IncorrectFormat.Wrapf("incorrect format for HashID, expected base64 encoded string, got %s", hashIDString)
}
