package base

import (
	"bytes"

	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/traits"
)

type maintainerID struct {
	ids.ClassificationID
	ids.IdentityID
}

var _ ids.MaintainerID = (*maintainerID)(nil)

func (maintainerID maintainerID) Bytes() []byte {
	return append(
		maintainerID.ClassificationID.Bytes(),
		maintainerID.IdentityID.Bytes()...)
}
func (maintainerID maintainerID) String() string {
	return stringUtilities.JoinIDStrings(maintainerID.ClassificationID.String(), maintainerID.IdentityID.String())
}
func (maintainerID maintainerID) Compare(listable traits.Listable) int {
	return bytes.Compare(maintainerID.Bytes(), maintainerIDFromInterface(listable).Bytes())
}
func (maintainerID maintainerID) GetClassificationID() ids.ClassificationID {
	return maintainerID.ClassificationID
}
func maintainerIDFromInterface(i interface{}) maintainerID {
	switch value := i.(type) {
	case maintainerID:
		return value
	default:
		panic(constants.MetaDataError)
	}
}
func NewMaintainerID(classificationID ids.ClassificationID, identityID ids.IdentityID) ids.MaintainerID {
	return maintainerID{
		ClassificationID: classificationID,
		IdentityID:       identityID,
	}
}

func ReadMaintainerID(maintainerIDString string) ids.MaintainerID {

}
