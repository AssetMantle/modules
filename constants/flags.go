package constants

import (
	"github.com/persistenceOne/persistenceSDK/types/utility"
)

var (
	AssetID          = utility.NewCLIFlag("assetID", "", "OwnableID")
	Burn             = utility.NewCLIFlag("burn", int64(-1), "Burn")
	ClassificationID = utility.NewCLIFlag("classificationID", "", "GetClassificationID")
	IdentityID       = utility.NewCLIFlag("identityID", "", "OwnerID")
	Lock             = utility.NewCLIFlag("lock", int64(-1), "Lock")
	MaintainersID    = utility.NewCLIFlag("maintainersID", "", "GetMaintainersID")
	OwnableID        = utility.NewCLIFlag("ownableID", "", "OwnableID")
	Properties       = utility.NewCLIFlag("properties", "", "GetProperties")
	Split            = utility.NewCLIFlag("split", "", "Split")
	SplitID          = utility.NewCLIFlag("splitID", "", "SplitID")
	To               = utility.NewCLIFlag("to", "", "To")
	ToID             = utility.NewCLIFlag("toID", "", "ToID")
)
