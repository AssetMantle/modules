package constants

import (
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	constantsQualified "github.com/AssetMantle/modules/schema/qualified/constants"
)

var NubClassification = baseDocuments.NewClassification(constantsQualified.NubImmutables, constantsQualified.NubMutables)

var MaintainerClassification = baseDocuments.NewClassification(constantsQualified.MaintainerImmutables, constantsQualified.MaintainerMutables)
