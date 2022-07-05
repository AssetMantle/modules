package utlities

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/utilities/string"
)

func GetZeroValueDataFromID(typeID ids.ID) data.Data {
	if zeroDataValue, err := ReadData(stringUtilities.JoinDataTypeAndValueStrings(typeID.String(), "")); err != nil {
		panic(err)
	} else {
		return zeroDataValue
	}
}
