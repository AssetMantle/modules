package mapper

import (
	"bytes"
	"encoding/base64"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"strings"
)

var _ types.ID = (*assetID)(nil)

type assetID struct {
	chainID          types.ID
	maintainersID    types.ID
	classificationID types.ID
	hashID           types.ID
}

func (assetID assetID) Bytes() []byte {
	return append(append(append(
		assetID.chainID.Bytes(),
		assetID.maintainersID.Bytes()...),
		assetID.classificationID.Bytes()...),
		assetID.hashID.Bytes()...)
}

func (assetID assetID) String() string {
	var values []string
	values = append(values, base64.URLEncoding.EncodeToString(assetID.chainID.Bytes()))
	values = append(values, base64.URLEncoding.EncodeToString(assetID.maintainersID.Bytes()))
	values = append(values, base64.URLEncoding.EncodeToString(assetID.classificationID.Bytes()))
	values = append(values, base64.URLEncoding.EncodeToString(assetID.hashID.Bytes()))
	return strings.Join(values, constants.IDSeparator)
}

func (assetID assetID) Compare(id types.ID) int {
	return bytes.Compare(assetID.Bytes(), id.Bytes())
}
