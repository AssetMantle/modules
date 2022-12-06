package constansts

import (
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	qualifiedConstants "github.com/AssetMantle/modules/schema/qualified/constants"
)

var NubClassificationID = baseIDs.GenerateClassificationID(qualifiedConstants.NubImmutables, qualifiedConstants.NubMutables)
var MaintainerClassificationID = baseIDs.GenerateClassificationID(qualifiedConstants.MaintainerImmutables, qualifiedConstants.MaintainerMutables)
