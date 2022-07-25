package base

import (
	"github.com/AssetMantle/modules/schema/ids"
)

type classificationID struct {
	ids.HashID
}

var _ ids.ClassificationID = (*classificationID)(nil)
