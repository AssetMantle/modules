// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

const Name = "orders"
const ConsensusVersion = 1

var StoreKeyPrefix = constants.OrdersStoreKeyPrefix

// TODO move to common constants
// TODO move to proper package
var NubClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.NubIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constantProperties.AuthenticationProperty)))
var ModuleIdentityID = baseIDs.NewIdentityID(NubClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(constantProperties.NubIDProperty.GetKey(), baseData.NewStringData(Name)))))
