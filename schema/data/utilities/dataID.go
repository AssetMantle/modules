package utilities

import (
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
)

func GetZeroValueDataFromID(typeID ids.ID) data.DataI {
	if zeroDataValue, err := ReadData(joinDataTypeAndValueStrings(typeID.String(), "")); err != nil {
		panic(err)
	} else {
		return zeroDataValue
	}
}
