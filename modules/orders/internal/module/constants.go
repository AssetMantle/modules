// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package module

import (
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/ids/constansts"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	constantProperties "github.com/AssetMantle/modules/schema/properties/constants"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
)

const Name = "orders"

var StoreKeyPrefix = constants.OrdersStoreKeyPrefix
var ModuleIdentityID = baseIDs.NewIdentityID(constansts.NubClassificationID, baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewProperty(constantProperties.NubIDProperty.GetKey(), baseData.NewStringData(Name)))))
