package constants

import (
	"github.com/persistenceOne/persistenceSDK/types/utility/base"
)

var (
	AssetID          = base.NewCLIFlag("assetID", "", "OwnableID")
	Burn             = base.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = base.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID       = base.NewCLIFlag("identityID", "", "OwnerID")
	Lock             = base.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = base.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID        = base.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = base.NewCLIFlag("properties", "", "GetProperties")
	Split            = base.NewCLIFlag("split", "", "Split")
	SplitID          = base.NewCLIFlag("splitID", "", "SplitID")
	To               = base.NewCLIFlag("to", "", "To")
	ToID             = base.NewCLIFlag("toID", "", "ToID")
)
