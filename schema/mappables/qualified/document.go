// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package qualified

import (
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/qualified/base"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Document struct {
	ID               types.ID `json:"id" valid:"required~required field id missing"`
	ClassificationID types.ID `json:"classificationID" valid:"required~required field classificationID missing"`
	base.HasImmutables
	base.HasMutables //nolint:govet
}

var _ mappables.Document = (*Document)(nil)

func (document Document) GetID() types.ID               { return document.ID }
func (document Document) GetClassificationID() types.ID { return document.ClassificationID }
func (document Document) GetProperty(id types.ID) types.Property {
	if property := document.HasImmutables.GetImmutableProperties().Get(id); property != nil {
		return property
	} else if property := document.HasMutables.GetMutableProperties().Get(id); property != nil {
		return property
	} else {
		return nil
	}
}
