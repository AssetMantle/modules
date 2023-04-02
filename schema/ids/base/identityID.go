package base

import (
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/constants"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
	"strings"
)

var _ ids.IdentityID = (*IdentityID)(nil)

func (identityID *IdentityID) ValidateBasic() error {
	return identityID.HashID.ValidateBasic()
}
func (identityID *IdentityID) IsIdentityID() {}
func (identityID *IdentityID) GetTypeID() ids.StringID {
	return NewStringID(constants.IdentityIDType)
}
func (identityID *IdentityID) FromString(idString string) (ids.ID, error) {
	idString = strings.TrimSpace(idString)
	if idString == "" {
		return PrototypeIdentityID(), nil
	}

	if hashID, err := PrototypeHashID().FromString(idString); err != nil {
		return PrototypeIdentityID(), err
	} else {
		return &IdentityID{
			HashID: hashID.(*HashID),
		}, nil
	}
}
func (identityID *IdentityID) AsString() string {
	return identityID.HashID.AsString()
}
func (identityID *IdentityID) GetHashID() ids.HashID {
	return identityID.HashID
}
func (identityID *IdentityID) Bytes() []byte {
	return identityID.HashID.Bytes()
}
func (identityID *IdentityID) Compare(listable traits.Listable) int {
	return identityID.HashID.Compare(identityIDFromInterface(listable).HashID)
}
func (identityID *IdentityID) ToAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_IdentityID{
			IdentityID: identityID,
		},
	}
}

func identityIDFromInterface(i interface{}) *IdentityID {
	switch value := i.(type) {
	case *IdentityID:
		return value
	default:
		panic(errorConstants.IncorrectFormat.Wrapf("expected *IdentityID, got %T", i))
	}
}

func NewIdentityID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.IdentityID {
	return &IdentityID{
		HashID: GenerateHashID(classificationID.Bytes(), immutables.GenerateHashID().Bytes()).(*HashID),
	}
}

func PrototypeIdentityID() ids.IdentityID {
	return &IdentityID{
		HashID: PrototypeHashID().(*HashID),
	}
}

func ReadIdentityID(identityIDString string) (ids.IdentityID, error) {
	if hashID, err := ReadHashID(identityIDString); err == nil {
		return &IdentityID{
			HashID: hashID.(*HashID),
		}, nil
	} else {
		return PrototypeIdentityID(), err
	}
}
