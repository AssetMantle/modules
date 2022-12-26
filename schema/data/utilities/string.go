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
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func readAccAddressData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.AccAddressDataPrototype(), nil
	}

	accAddress, err := sdkTypes.AccAddressFromBech32(dataString)
	if err != nil {
		return base.AccAddressDataPrototype(), err
	}

	return base.NewAccAddressData(accAddress), nil
}
func readBooleanData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.BooleanDataPrototype(), nil
	}

	Bool, err := strconv.ParseBool(dataString)
	if err != nil {
		return base.BooleanDataPrototype(), err
	}

	return base.NewBooleanData(Bool), nil
}
func readDecData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.DecDataPrototype(), nil
	}

	dec, err := sdkTypes.NewDecFromStr(dataString)
	if err != nil {
		return base.DecDataPrototype(), err
	}

	return base.NewDecData(dec), nil
}
func readHeightData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.HeightDataPrototype(), nil
	}

	height, err := strconv.ParseInt(dataString, 10, 64)
	if err != nil {
		return base.HeightDataPrototype(), err
	}

	return base.NewHeightData(baseTypes.NewHeight(height)), nil
}

// TODO read complex IDs than string ID
func readIDData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.IDDataPrototype(), nil
	}

	return base.NewIDData(baseIDs.NewStringID(dataString)), nil
}
func readListData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.ListDataPrototype(), nil
	}

	dataStringList := stringUtilities.SplitListString(dataString)
	dataList := make([]data.Data, len(dataStringList))

	for i, datumString := range dataStringList {
		Data, err := ReadData(datumString)
		if err != nil {
			return base.ListDataPrototype(), err
		}

		dataList[i] = Data
	}

	return base.NewListData(base.NewDataList(dataList...)), nil
}
func readStringData(dataString string) (data.Data, error) {
	if dataString == "" {
		return base.StringDataPrototype(), nil
	}
	return base.NewStringData(dataString), nil
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

		switch baseIDs.NewStringID(dataTypeString).String() {
		case dataConstants.AccAddressDataID.String():
			Data, err = readAccAddressData(dataValueString)
		case dataConstants.BooleanDataID.String():
			Data, err = readBooleanData(dataValueString)
		case dataConstants.DecDataID.String():
			Data, err = readDecData(dataValueString)
		case dataConstants.HeightDataID.String():
			Data, err = readHeightData(dataValueString)
		case dataConstants.IDDataID.String():
			Data, err = readIDData(dataValueString)
		case dataConstants.ListDataID.String():
			Data, err = readListData(dataValueString)
		case dataConstants.StringDataID.String():
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
