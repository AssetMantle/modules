// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	"github.com/AssetMantle/modules/schema/documents/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	qualifiedConstants "github.com/AssetMantle/modules/schema/qualified/constants"
)

const Name = "maintainers"

var StoreKeyPrefix = constants.MaintainersStoreKeyPrefix

var MaintainerClassificationID = baseIDs.NewClassificationID(qualifiedConstants.MaintainerImmutables, qualifiedConstants.MaintainerMutables)

var GenMaintainer = base.NewMaintainer(MaintainerClassificationID, qualifiedConstants.MaintainerImmutables, qualifiedConstants.MaintainerMutables)
