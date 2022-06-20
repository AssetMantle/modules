package utlities

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	constantIDs "github.com/AssetMantle/modules/schema/data/constants"
	constants2 "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (data.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataTypeID, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var Data data.Data

		var err error

		switch baseIDs.NewID(dataTypeID) {
		case constantIDs.AccAddressDataID:
			Data, err = baseData.ReadAccAddressData(dataString)
		case constantIDs.BooleanDataID:
			Data, err = baseData.ReadBooleanData(dataString)
		case constantIDs.DecDataID:
			Data, err = baseData.ReadDecData(dataString)
		case constantIDs.HeightDataID:
			Data, err = baseData.ReadHeightData(dataString)
		case constantIDs.IDDataID:
			Data, err = baseData.ReadIDData(dataString)
		case constantIDs.ListDataID:
			Data, err = baseData.ReadListData(dataString)
		case constantIDs.StringDataID:
			Data, err = baseData.ReadStringData(dataString)
		default:
			Data, err = nil, constants2.UnsupportedParameter
		}

		if err != nil {
			return nil, err
		}

		return Data, nil
	}

	return nil, constants2.IncorrectFormat
}
