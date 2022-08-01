// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type document struct {
	ids.ClassificationID
	qualified.Immutables
	qualified.Mutables
}

var _ qualified.Document = (*document)(nil)

func (document document) GetClassificationID() ids.ClassificationID {
	return document.ClassificationID
}
func (document document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if property := document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document document) GetImmutables() qualified.Immutables {
	return document.Immutables
}
func (document document) GetMutables() qualified.Mutables {
	return document.Mutables
}

// TODO write test case
// TODO check is not metaProperty
func (document document) Mutate(propertyList ...properties.Property) qualified.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...).(mutables)
	return document
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) qualified.Document {
	// TODO check if the document conforms to the classification
	return document{
		ClassificationID: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
