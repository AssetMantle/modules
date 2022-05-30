package utlities

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	constantIDs "github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (data.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataTypeID, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data data.Data

		var Error error

		switch baseIDs.NewID(dataTypeID) {
		case constantIDs.AccAddressDataID:
			data, Error = baseData.ReadAccAddressData(dataString)
		case constantIDs.BooleanDataID:
			data, Error = baseData.ReadBooleanData(dataString)
		case constantIDs.DecDataID:
			data, Error = baseData.ReadDecData(dataString)
		case constantIDs.HeightDataID:
			data, Error = baseData.ReadHeightData(dataString)
		case constantIDs.IDDataID:
			data, Error = baseData.ReadIDData(dataString)
		case constantIDs.ListDataID:
			data, Error = baseData.ReadListData(dataString)
		case constantIDs.StringDataID:
			data, Error = baseData.ReadStringData(dataString)
		default:
			data, Error = nil, errors.UnsupportedParameter
		}

		if Error != nil {
			return nil, Error
		}

		return data, nil
	}

	return nil, errors.IncorrectFormat
}
