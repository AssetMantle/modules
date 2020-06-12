package asset

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/constants"
	"github.com/persistenceOne/persistenceSDK/types"
	"github.com/spf13/cobra"
	"strings"
)

func QueryCommand(codec *codec.Codec) *cobra.Command {
	command := &cobra.Command{
		Use:   constants.AssetQuery,
		Short: "Query one or multiple assets.",
		Long:  "",
		RunE: func(command *cobra.Command, args []string) error {
			cliContext := context.NewCLIContext().WithCodec(codec)

			bytes := packageCodec.MustMarshalJSON(query{id: types.BaseID{IDString: constants.AssetID.ReadCLIValue().(string)}})

			response, _, queryWithDataError := cliContext.QueryWithData(strings.Join([]string{"", "custom", constants.QuerierRoute, constants.AssetQuery}, "/"), bytes)
			if queryWithDataError != nil {
				return queryWithDataError
			}

			var asset types.InterNFT
			unmarshalJSONError := codec.UnmarshalJSON(response, &asset)
			if unmarshalJSONError != nil {
				return unmarshalJSONError
			}
			return cliContext.PrintOutput(asset)
		},
	}

	constants.AssetQueryCommand.RegisterFlags(command)
	return flags.GetCommands(command)[0]
}
