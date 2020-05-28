package mapper

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"sort"
	"strings"
)

func AssetIDFromString(assetIDString string) assetID {
	base64IDs := strings.Split(assetIDString, constants.IDSeparator)
	chainID, _ := base64.URLEncoding.DecodeString(base64IDs[0])
	classificationID, _ := base64.URLEncoding.DecodeString(base64IDs[1])
	maintainersID, _ := base64.URLEncoding.DecodeString(base64IDs[2])
	hashID, _ := base64.URLEncoding.DecodeString(base64IDs[3])

	return assetID{
		chainID:          types.BaseID{BaseBytes: chainID},
		maintainersID:    types.BaseID{BaseBytes: classificationID},
		classificationID: types.BaseID{BaseBytes: maintainersID},
		hashID:           types.BaseID{BaseBytes: hashID},
	}
}

func MakeAsset(assetID assetID, properties [2][]string, lock int, burn int) types.Asset {
	types.BaseProperty{assetID}
}

func GenerateHashID(values []string) types.ID {
	sort.Strings(values)
	toDigest := strings.Join(values, constants.PropertySeparator)
	h := sha1.New()
	h.Write([]byte(toDigest))
	return types.BaseID{BaseBytes: h.Sum(nil)}
}
