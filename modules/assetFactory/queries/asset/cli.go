package asset

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
)

func QueryCommand(Codec *codec.Codec) *cobra.Command {

	makeQueryBytes := func(cliCommand types.CLICommand) []byte {
		return packageCodec.MustMarshalJSON(request{ID: types.NewID(cliCommand.ReadString(constants.AssetID))})
	}

	marshallResponse := func(bytes []byte) interface{} {
		var assets types.InterNFTs
		if err := Codec.UnmarshalJSON(bytes, &assets); err != nil {
			return nil
		}
		return assets
	}

	return types.NewCLICommand(
		constants.AssetQuery,
		constants.AssetQueryShort,
		constants.AssetQueryLong,
		[]types.CLIFlag{
			constants.AssetID,
		},
	).CreateQueryCommand(Codec, constants.QuerierRoute, makeQueryBytes, marshallResponse)
}
