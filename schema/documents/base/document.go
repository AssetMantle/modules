// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
)

type document struct {
	ids.ClassificationID
	qualified.Immutables
	qualified.Mutables
}

var _ documents.Document = (*document)(nil)

func (document document) GenerateHashID() ids.HashID {
	return base.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
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
func (document document) Mutate(propertyList ...properties.Property) documents.Document {
	document.Mutables = document.Mutables.Mutate(propertyList...)
	return document
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return document{
		ClassificationID: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
