package mapper

import (
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

func AssetIdentifiersFromString(assetID string) (chainID types.ID, maintainersID types.ID, classificationID types.ID, hashID types.ID) {
	base64IDs := strings.Split(assetID, constants.IDSeparator)
	baseAssetID := baseAssetID{base64.URLEncoding.DecodeString(base64IDs[0]),
		base64.URLEncoding.DecodeString(base64IDs[1]),
		base64.URLEncoding.DecodeString(base64IDs[2]),
		base64.URLEncoding.DecodeString(base64IDs[3]),
	}
	return baseAssetID.chainID, baseAssetID.maintainersID, baseAssetID.classificationID, baseAssetID.hashID
}
