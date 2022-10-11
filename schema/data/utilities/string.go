package utilities

import (
	"strconv"
	"strings"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/data/base"
	dataConstants "github.com/AssetMantle/modules/schema/data/constants"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func readAccAddressData(dataString string) (data.Data, error) {
	if dataString == "" {
		return GetZeroValueDataFromID(dataConstants.AccAddressDataID), nil
	}

	accAddress, err := sdkTypes.AccAddressFromBech32(dataString)
	if err != nil {
		return GetZeroValueDataFromID(dataConstants.AccAddressDataID), err
	}

	return base.NewAccAddressData(accAddress), nil
}
func readBooleanData(dataString string) (data.Data, error) {
	if dataString == "" {
		return GetZeroValueDataFromID(dataConstants.BooleanDataID), nil
	}

	Bool, err := strconv.ParseBool(dataString)
	if err != nil {
		return GetZeroValueDataFromID(dataConstants.BooleanDataID), err
	}

	return base.NewBooleanData(Bool), nil
}
func readDecData(dataString string) (data.Data, error) {
	if dataString == "" {
		return GetZeroValueDataFromID(dataConstants.DecDataID), nil
	}

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return GetZeroValueDataFromID(dataConstants.DecDataID), err
	}

	return base.NewDecData(dec), nil
}
func readHeightData(dataString string) (data.Data, error) {
	if dataString == "" {
		return GetZeroValueDataFromID(dataConstants.HeightDataID), nil
	}

	height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return GetZeroValueDataFromID(dataConstants.HeightDataID), err
	}

	return base.NewHeightData(baseTypes.NewHeight(height)), nil
}
func readIDData(idData string) (data.Data, error) {
	return base.NewIDData(baseIDs.NewStringID(idData)), nil
}
func readListData(dataString string) (data.Data, error) {
	if dataString == "" {
		return GetZeroValueDataFromID(dataConstants.ListDataID), nil
	}

	dataStringList := stringUtilities.SplitListString(dataString)
	dataList := make([]data.Data, len(dataStringList))

	for i, datumString := range dataStringList {
		Data, err := ReadData(datumString)
		if err != nil {
			return GetZeroValueDataFromID(dataConstants.ListDataID), err
		}

		dataList[i] = Data
	}

	return base.NewListData(baseLists.NewDataList(dataList...)), nil
}
func readStringData(stringData string) (data.Data, error) {
	return base.NewStringData(stringData), nil
}
func joinDataTypeAndValueStrings(dataType, dataValue string) string {
	return strings.Join([]string{dataType, dataValue}, dataConstants.DataTypeAndValueSeparator)
}
func splitDataTypeAndValueStrings(dataTypeAndValueString string) (dataType, dataValue string) {
	if dataTypeAndValue := strings.SplitN(dataTypeAndValueString, dataConstants.DataTypeAndValueSeparator, 2); len(dataTypeAndValue) < 2 {
		return "", ""
	} else {
		return dataTypeAndValue[0], dataTypeAndValue[1]
	}
}

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (data.Data, error) {
	dataTypeString, dataValueString := splitDataTypeAndValueStrings(dataString)
	if dataTypeString != "" {
		var Data data.Data

		var err error

		switch baseIDs.NewStringID(dataTypeString) {
		case dataConstants.AccAddressDataID:
			Data, err = readAccAddressData(dataValueString)
		case dataConstants.BooleanDataID:
			Data, err = readBooleanData(dataValueString)
		case dataConstants.DecDataID:
			Data, err = readDecData(dataValueString)
		case dataConstants.HeightDataID:
			Data, err = readHeightData(dataValueString)
		case dataConstants.IDDataID:
			Data, err = readIDData(dataValueString)
		case dataConstants.ListDataID:
			Data, err = readListData(dataValueString)
		case dataConstants.StringDataID:
			Data, err = readStringData(dataValueString)
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
