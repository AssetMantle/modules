// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/AssetMantle/modules/schema/documents"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/qualified"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

type document struct {
	ids.ClassificationID
	qualified.Immutables
	qualified.Mutables
}

var _ documents.Document = (*document)(nil)

func (document document) GenerateHashID() ids.HashID {
	return baseIDs.GenerateHashID(document.GetClassificationID().Bytes(), document.GetImmutables().GenerateHashID().Bytes())
}
func (document document) GetClassificationID() ids.ClassificationID {
	if document.ClassificationID == nil {
		return baseIDs.PrototypeClassificationID()
	}
	return document.ClassificationID
}
func (document document) GetProperty(propertyID ids.PropertyID) properties.Property {
	if sanitizedDocument, err := document.Sanitize(); err == nil {
		document.Immutables = sanitizedDocument.(documents.Document).GetImmutables()
		document.Mutables = sanitizedDocument.(documents.Document).GetMutables()
		document.ClassificationID = sanitizedDocument.(documents.Document).GetClassificationID()
	}
	if property := document.Immutables.GetImmutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else if property := document.Mutables.GetMutablePropertyList().GetProperty(propertyID); property != nil {
		return property
	} else {
		return nil
	}
}
func (document document) GetImmutables() qualified.Immutables {
	if document.Immutables == nil {
		return baseQualified.NewImmutables(baseLists.NewPropertyList())
	}
	return document.Immutables
}
func (document document) GetMutables() qualified.Mutables {
	if document.Mutables == nil {
		return baseQualified.NewMutables(baseLists.NewPropertyList())
	}
	return document.Mutables
}

// TODO write test case
func (document document) Mutate(propertyList ...properties.Property) documents.Document {
	if sanitizedDocument, err := document.Sanitize(); err == nil {
		document.Immutables = sanitizedDocument.(documents.Document).GetImmutables()
		document.Mutables = sanitizedDocument.(documents.Document).GetMutables()
		document.ClassificationID = sanitizedDocument.(documents.Document).GetClassificationID()
	}
	document.Mutables = document.Mutables.Mutate(propertyList...)
	return document
}

func (document document) Sanitize() (documents.Document, error) {
	if document.ClassificationID == nil {
		document.ClassificationID = baseIDs.PrototypeClassificationID()
	}
	if document.Immutables == nil {
		document.Immutables = baseQualified.NewImmutables(baseLists.NewPropertyList())
	}
	if document.Mutables == nil {
		document.Mutables = baseQualified.NewMutables(baseLists.NewPropertyList())
	}
	return document, nil
}

func NewDocument(classificationID ids.ClassificationID, immutables qualified.Immutables, mutables qualified.Mutables) documents.Document {
	return document{
		ClassificationID: classificationID,
		Immutables:       immutables,
		Mutables:         mutables,
	}
}
