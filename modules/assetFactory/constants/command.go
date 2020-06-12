package constants

import "github.com/persistenceOne/persistenceSDK/types"

var AssetQueryCommand = types.NewCLICommand(AssetQuery, "", "", []types.CLIFlag{AssetID})
