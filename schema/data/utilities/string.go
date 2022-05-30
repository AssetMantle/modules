package utilities

import (
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	idsConstants "github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/types"
)

// ReadData
// CHECK-TODO if data type added see if added here
func ReadData(dataString string) (types.Data, error) {
	dataTypeAndString := strings.SplitN(dataString, constants.DataTypeAndValueSeparator, 2)
	if len(dataTypeAndString) == 2 {
		dataTypeID, dataString := dataTypeAndString[0], dataTypeAndString[1]

		var data types.Data

		var Error error

		switch baseIDs.NewID(dataTypeID) {
		case idsConstants.AccAddressDataID:
			data, Error = base.ReadAccAddressData(dataString)
		case idsConstants.BooleanDataID:
			data, Error = base.ReadBooleanData(dataString)
		case idsConstants.DecDataID:
			data, Error = base.ReadDecData(dataString)
		case idsConstants.HeightDataID:
			data, Error = base.ReadHeightData(dataString)
		case idsConstants.IDDataID:
			data, Error = base.ReadIDData(dataString)
		case idsConstants.ListDataID:
			data, Error = base.ReadListData(dataString)
		case idsConstants.StringDataID:
			data, Error = base.ReadStringData(dataString)
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
