package utlities

import (
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	constantIDs "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (data.Data, error) {
	dataTypeString, dataValueString := stringUtilities.SplitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data

		var err error

		switch baseIDs.NewID(dataTypeString) {
		case constantIDs.AccAddressDataID:
			Data, err = baseData.ReadAccAddressData(dataValueString)
		case constantIDs.BooleanDataID:
			Data, err = baseData.ReadBooleanData(dataValueString)
		case constantIDs.DecDataID:
			Data, err = baseData.ReadDecData(dataValueString)
		case constantIDs.HeightDataID:
			Data, err = baseData.ReadHeightData(dataValueString)
		case constantIDs.IDDataID:
			Data, err = baseData.ReadIDData(dataValueString)
		case constantIDs.ListDataID:
			Data, err = baseData.ReadListData(dataValueString)
		case constantIDs.StringDataID:
			Data, err = baseData.ReadStringData(dataValueString)
		default:
			Data, err = nil, errorConstants.UnsupportedParameter
		}

		if err != nil {
			return nil, err
		}

		return Data, nil
	}

	return nil, errorConstants.IncorrectFormat
}
