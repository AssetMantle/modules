package utlities

import (
	"strings"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
)

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (data.Data, error) {
	dataTypeString, dataValueString := SplitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data

		var err error

		switch baseIDs.NewID(dataTypeString) {
		case dataConstants.AccAddressDataID:
			Data, err = baseData.ReadAccAddressData(dataValueString)
		case dataConstants.BooleanDataID:
			Data, err = baseData.ReadBooleanData(dataValueString)
		case dataConstants.DecDataID:
			Data, err = baseData.ReadDecData(dataValueString)
		case dataConstants.HeightDataID:
			Data, err = baseData.ReadHeightData(dataValueString)
		case dataConstants.IDDataID:
			Data, err = baseData.ReadIDData(dataValueString)
		case dataConstants.ListDataID:
			Data, err = baseData.ReadListData(dataValueString)
		case dataConstants.StringDataID:
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

func JoinDataTypeAndValueStrings(dataType, dataValue string) string {
	return strings.Join([]string{dataType, dataValue}, dataConstants.DataTypeAndValueSeparator)
}

func SplitDataTypeAndValueStrings(dataTypeAndValueString string) (dataType, dataValue string) {
	if dataTypeAndValue := strings.SplitN(dataTypeAndValueString, dataConstants.DataTypeAndValueSeparator, 2); len(dataTypeAndValue) < 2 {
		return "", ""
	} else {
		return dataTypeAndValue[0], dataTypeAndValue[1]
	}
}
