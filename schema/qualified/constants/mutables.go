package constants

import (
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

var NubMutables = base.NewMutables(baseLists.NewPropertyList(constants.AuthenticationProperty))
var MaintainerMutables = base.NewMutables(baseLists.NewPropertyList(constants.MaintainedPropertiesProperty, constants.PermissionsProperty))
