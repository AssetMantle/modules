// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	baseData "github.com/AssetMantle/schema/go/data/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"

	"github.com/AssetMantle/modules/helpers/constants"
)

const Name = "orders"
const ConsensusVersion = 1

var StoreKeyPrefix = constants.OrdersStoreKeyPrefix

// TODO move to common constants
// TODO move to proper package
var NubClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList(constantProperties.NubIDProperty)), baseQualified.NewMutables(baseLists.NewPropertyList(constantProperties.AuthenticationProperty)))
var ModuleIdentityID = baseIDs.NewIdentityID(NubClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(constantProperties.NubIDProperty.GetKey(), baseData.NewStringData(Name)))))
