// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/qualified/base"
)

//type document struct {
//	ids.ClassificationID
//	qualified.Immutables
//	qualified.Mutables
//}

var _ documents.Document = (*Document)(nil)

func (document *Document) GenerateHashID() ids.ID {
	return baseIDs.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document *Document) GetProperty(propertyID ids.ID) properties.Property {
	if property := document.Immutables.GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}

// TODO write test case
func (document *Document) Mutate(propertyList ...properties.Property) documents.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...).(*base.Mutables)
	return document
}

func NewDocument(classificationID ids.ID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return &Document{
		ClassificationID: classificationID.(*baseIDs.ID),
		Immutables:       immutables.(*base.Immutables),
		Mutables:         mutables.(*base.Mutables),
	}
}
