package base

import (
	"strings"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type idGetter interface {
	get() ids.ID
}

var _ idGetter = (*AnyID_AssetID)(nil)
var _ idGetter = (*AnyID_ClassificationID)(nil)
var _ idGetter = (*AnyID_CoinID)(nil)
var _ idGetter = (*AnyID_DataID)(nil)
var _ idGetter = (*AnyID_HashID)(nil)
var _ idGetter = (*AnyID_IdentityID)(nil)
var _ idGetter = (*AnyID_MaintainerID)(nil)
var _ idGetter = (*AnyID_OrderID)(nil)
var _ idGetter = (*AnyID_OwnableID)(nil)
var _ idGetter = (*AnyID_PropertyID)(nil)
var _ idGetter = (*AnyID_SplitID)(nil)
var _ idGetter = (*AnyID_StringID)(nil)

func (m *AnyID_AssetID) get() ids.ID {
	return m.AssetID
}
func (m *AnyID_ClassificationID) get() ids.ID {
	return m.ClassificationID
}
func (m *AnyID_CoinID) get() ids.ID {
	return m.CoinID
}
func (m *AnyID_DataID) get() ids.ID {
	return m.DataID
}
func (m *AnyID_HashID) get() ids.ID {
	return m.HashID
}
func (m *AnyID_IdentityID) get() ids.ID {
	return m.IdentityID
}
func (m *AnyID_MaintainerID) get() ids.ID {
	return m.MaintainerID
}
func (m *AnyID_OrderID) get() ids.ID {
	return m.OrderID
}
func (m *AnyID_OwnableID) get() ids.ID {
	return m.OwnableID
}
func (m *AnyID_PropertyID) get() ids.ID {
	return m.PropertyID
}
func (m *AnyID_SplitID) get() ids.ID {
	return m.SplitID
}
func (m *AnyID_StringID) get() ids.ID {
	return m.StringID
}

var _ ids.AnyID = (*AnyID)(nil)

func (m *AnyID) GetTypeID() ids.StringID {
	return m.Impl.(idGetter).get().GetTypeID()
}
func (m *AnyID) FromString(idString string) (ids.ID, error) {
	idTypeString, idValueString := splitIDTypeAndValueStrings(idString)
	if idTypeString != "" {
		var id ids.ID
		var err error

		switch NewStringID(idTypeString).AsString() {
		case PrototypeAssetID().GetTypeID().AsString():
			id, err = PrototypeAssetID().FromString(idValueString)
		case PrototypeClassificationID().GetTypeID().AsString():
			id, err = PrototypeClassificationID().FromString(idValueString)
		case PrototypeCoinID().GetTypeID().AsString():
			id, err = PrototypeCoinID().FromString(idValueString)
		case PrototypeDataID().GetTypeID().AsString():
			id, err = PrototypeDataID().FromString(idValueString)
		case PrototypeHashID().GetTypeID().AsString():
			id, err = PrototypeHashID().FromString(idValueString)
		case PrototypeIdentityID().GetTypeID().AsString():
			id, err = PrototypeIdentityID().FromString(idValueString)
		case PrototypeMaintainerID().GetTypeID().AsString():
			id, err = PrototypeMaintainerID().FromString(idValueString)
		case PrototypeOrderID().GetTypeID().AsString():
			id, err = PrototypeOrderID().FromString(idValueString)
		case PrototypePropertyID().GetTypeID().AsString():
			id, err = PrototypePropertyID().FromString(idValueString)
		case PrototypeSplitID().GetTypeID().AsString():
			id, err = PrototypeSplitID().FromString(idValueString)
		case PrototypeStringID().GetTypeID().AsString():
			id, err = PrototypeStringID().FromString(idValueString)
		default:
			id, err = nil, errorConstants.IncorrectFormat.Wrapf("type identifier is not recognised")
		}

		if err != nil {
			return nil, err
		}

		return id, nil
	}

	return nil, errorConstants.IncorrectFormat.Wrapf("type identifier is missing")
}
func (m *AnyID) AsString() string {
	return joinIDTypeAndValueStrings(m.Impl.(idGetter).get().GetTypeID().AsString(), m.Impl.(idGetter).get().AsString())
}
func (m *AnyID) Get() ids.ID {
	return m.Impl.(idGetter).get()
}
func (m *AnyID) Compare(listable traits.Listable) int {
	return m.Impl.(idGetter).get().Compare(listable)
}
func (m *AnyID) Bytes() []byte {
	return m.Impl.(idGetter).get().Bytes()
}
func (m *AnyID) ToAnyID() ids.AnyID {
	return m.Impl.(idGetter).get().ToAnyID()
}
func (m *AnyID) ValidateBasic() error {
	return m.Impl.(idGetter).get().ValidateBasic()
}

func PrototypeAnyID() ids.AnyID {
	return &AnyID{
		Impl: &AnyID_StringID{
			StringID: PrototypeStringID().(*StringID),
		},
	}
}

func idFromListable(listable traits.Listable) (ids.ID, error) {
	switch listable.(type) {
	case ids.ID:
		return listable.(ids.ID), nil

	default:
		return nil, errorConstants.IncorrectFormat.Wrapf("unsupported type")
	}
}
func joinIDTypeAndValueStrings(idTypes, idValue string) string {
	return strings.TrimSpace(idTypes) + idTypeAndValueSeparator + strings.TrimSpace(idValue)
}
func splitIDTypeAndValueStrings(idTypeAndValueString string) (idType, idValue string) {
	if idTypeAndValue := strings.SplitN(idTypeAndValueString, idTypeAndValueSeparator, 2); len(idTypeAndValue) == 1 {
		return strings.TrimSpace(idTypeAndValue[0]), ""
	} else if len(idTypeAndValue) == 2 {
		return strings.TrimSpace(idTypeAndValue[0]), strings.TrimSpace(idTypeAndValue[1])
	} else {
		return "", ""
	}
}

const idTypeAndValueSeparator = "|"
