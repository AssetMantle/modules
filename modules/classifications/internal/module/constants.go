// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	baseDocuments "github.com/AssetMantle/modules/schema/documents/base"
	constantsHelpers "github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	constantsQualified "github.com/AssetMantle/modules/schema/qualified/constants"
)

const Name = "classifications"

var StoreKeyPrefix = constantsHelpers.ClassificationsStoreKeyPrefix

// MaxPropertyCount TODO convert it to module param
const MaxPropertyCount = 22

var NubClassification = baseDocuments.NewClassification(constantsQualified.NubImmutables, constantsQualified.NubMutables)

var MaintainerClassificationID = baseIDs.NewClassificationID(constantsQualified.MaintainerImmutables, constantsQualified.MaintainerMutables)

var GenesisMaintainer = baseDocuments.NewMaintainer(MaintainerClassificationID, constantsQualified.MaintainerImmutables, constantsQualified.MaintainerMutables)
